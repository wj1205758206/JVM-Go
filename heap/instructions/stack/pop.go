package stack

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}

//用于弹出int、float等占用一个操作数栈位置的变量
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PopSlot()
}

//用于弹出double和long等类型占用两个操作数栈位置的变量
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
