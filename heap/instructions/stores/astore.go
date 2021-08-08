package stores

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//ASTORE 把引用类型变量从操作数栈顶弹出，然后存入局部变量表

type ASTORE struct {
	base.Index8Instruction
}

type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.GetOperandStack().PopRef()
	frame.GetLocalVars().SetRef(index, ref)
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
