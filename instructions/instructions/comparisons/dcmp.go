package comparisons

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//DCMPG 比较两个double类型变量的大小，如果遇到NaN值，dcmpg会压入1
type DCMPG struct {
	base.NoOperandsInstruction
}

//DCMPL 比较两个double类型变量的大小，如果遇到NaN值，dcmpl会压入-1
type DCMPL struct {
	base.NoOperandsInstruction
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}
