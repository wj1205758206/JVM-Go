package extended

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//IFNULL 根据引用是否是null进行跳转,如果是null则跳转
type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref := stack.PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
