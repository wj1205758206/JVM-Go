package extended

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//IFNONNULL 根据引用是否是null进行跳转,如果不是null则跳转
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref := stack.PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
