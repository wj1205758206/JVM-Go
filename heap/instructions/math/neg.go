package math

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

type INEG struct {
	base.NoOperandsInstruction
}

type LNEG struct {
	base.NoOperandsInstruction
}

type FNEG struct {
	base.NoOperandsInstruction
}

type DNEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}
