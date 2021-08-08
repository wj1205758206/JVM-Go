package comparisons

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
)

//FCMPG 比较两个float类型变量的大小，如果遇到NaN值，fcmpg会压入1
type FCMPG struct {
	base.NoOperandsInstruction
}

//FCMPL 比较两个float类型变量的大小，如果遇到NaN值，fcmpl会压入-1
type FCMPL struct {
	base.NoOperandsInstruction
}

//_fcmp 统一进行比较，由于浮点数计算有可能产生NaN值，所以比较两个浮点数时，除了大于、等于、小于之外， 还有无法比较
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag { //如果遇到NaN值，fcmpg会压入1 fcmpl会压入-1
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
