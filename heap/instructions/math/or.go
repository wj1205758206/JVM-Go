package math

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//IOR int类型按位或指令
type IOR struct {
	base.NoOperandsInstruction
}

//LOR long类型按位或指令
type LOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
