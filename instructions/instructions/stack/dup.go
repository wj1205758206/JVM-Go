package stack

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//DUP 复制栈顶的单个变量
type DUP struct {
	base.NoOperandsInstruction
}

//DUP_X1 复制栈顶变量，插入到前两个值的下面
type DUP_X1 struct {
	base.NoOperandsInstruction
}

//DUP_X2 复制栈顶变量，插入到前两个或者三个值的下面
type DUP_X2 struct {
	base.NoOperandsInstruction
}

//DUP2 复制栈顶的一个或者两个变量
type DUP2 struct {
	base.NoOperandsInstruction
}

//DUP2_X1 复制栈顶的一个或者两个变量，插入到前两个或者三个值下面
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

//DUP2_X2 复制栈顶的一个或者两个变量，插入到前两个或者三个或者四个值下面
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}