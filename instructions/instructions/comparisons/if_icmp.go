package comparisons

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//IF_ICMPEQ ==成立则跳转
type IF_ICMPEQ struct {
	base.BranchInstruction
}

//IF_ICMPNE !=成立则跳转
type IF_ICMPNE struct {
	base.BranchInstruction
}

//IF_ICMPLT <成立则跳转
type IF_ICMPLT struct {
	base.BranchInstruction
}

//IF_ICMPLE <=成立则跳转
type IF_ICMPLE struct {
	base.BranchInstruction
}

//IF_ICMPGT >成立则跳转
type IF_ICMPGT struct {
	base.BranchInstruction
}

//IF_ICMPGE >=成立则跳转
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}
