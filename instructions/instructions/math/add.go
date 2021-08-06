package math

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//IADD int类型加法指令
type IADD struct {
	base.NoOperandsInstruction
}

//LADD long类型加法指令
type LADD struct {
	base.NoOperandsInstruction
}

//FADD float类型加法指令
type FADD struct {
	base.NoOperandsInstruction
}

//DADD double类型加法指令
type DADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}
