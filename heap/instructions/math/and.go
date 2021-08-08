package math

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//IAND int类型按位与指令
type IAND struct {
	base.NoOperandsInstruction
}

//LAND long类型按位与指令
type LAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
