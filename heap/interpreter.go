package main

import (
	"JVM-Go/instructions/instructions"
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/classfile"
	"JVM-Go/rtda/rtda"
	"fmt"
)

//interpreter 解释器
func interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.GetCodeAttribute() //获取它的Code属性
	maxLocals := codeAttr.GetMaxLocals()      //获得执行方法所需的局部变量表
	maxStack := codeAttr.GetMaxStack()        //获得执行方法所需的操作数栈
	byteCode := codeAttr.GetCode()            //获得执行方法所需的字节码
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, byteCode)
}

//loop 循环执行“计算pc、解码指令、执行指令”这三个步骤
func loop(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//计算pc
		pc := frame.GetNextPC()
		thread.SetPC(pc)
		//解码指令
		reader.Reset(byteCode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.GetPC())
		//执行指令
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.GetLocalVars())
		fmt.Printf("OperandStack:%v\n", frame.GetOperandStack())
		panic(r)
	}
}
