package comparisons

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//IF_ACMPEQ 比较两个引用类型数值是否相等，相等则跳转
type IF_ACMPEQ struct {
	base.BranchInstruction
}

//IF_ACMPNE 比较两个引用类型数值是否不相等，不相等则跳转
type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, self.Offset)
	}
}
