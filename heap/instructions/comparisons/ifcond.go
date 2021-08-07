package comparisons

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//IFEQ ==
type IFEQ struct {
	base.BranchInstruction
}

//IFNE !=
type IFNE struct {
	base.BranchInstruction
}

//IFLT <
type IFLT struct {
	base.BranchInstruction
}

//IFLE <=
type IFLE struct {
	base.BranchInstruction
}

//IFGT >
type IFGT struct {
	base.BranchInstruction
}

//IFGE >=
type IFGE struct {
	base.BranchInstruction
}

//IFGE null
type IFNULL struct {
	base.BranchInstruction
}

//IFGE !null
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}

