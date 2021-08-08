package math

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//IMUL int类型乘法指令
type IMUL struct {
	base.NoOperandsInstruction
}

//LMUL long类型乘法指令
type LMUL struct {
	base.NoOperandsInstruction
}

//FMUL float类型乘法指令
type FMUL struct {
	base.NoOperandsInstruction
}

//DMUL double类型乘法指令
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}
