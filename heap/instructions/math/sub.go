package math

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//ISUB int类型减法指令
type ISUB struct {
	base.NoOperandsInstruction
}

//LSUB long类型减法指令
type LSUB struct {
	base.NoOperandsInstruction
}

//FSUB float类型减法指令
type FSUB struct {
	base.NoOperandsInstruction
}

//DSUB double类型减法指令
type DSUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}
