package main

import (
	"JVM-Go/heap/instructions"
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"fmt"
)

//interpreter 解释器
func interpreter(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.GetCode())
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
