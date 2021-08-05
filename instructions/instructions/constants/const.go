package constants

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

/*################### const系列指令 ###################*/
type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

// 把null引用压入操作数栈
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushRef(nil)
}

// 把Double类型的0.0引用压入操作数栈
func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushDouble(0.0)
}

// 把Double类型的1.0引用压入操作数栈
func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushDouble(1.0)
}

// 把Float类型的0.0引用压入操作数栈
func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushFloat(0.0)
}

// 把Float类型的1.0引用压入操作数栈
func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushFloat(1.0)
}

// 把Float类型的2.0引用压入操作数栈
func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushFloat(2.0)
}

// 把Int类型的-1引用压入操作数栈
func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(-1)
}

// 把Int类型的0引用压入操作数栈
func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(0)
}

// 把Int类型的1引用压入操作数栈
func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(1)
}

// 把Int类型的2引用压入操作数栈
func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(2)
}

// 把Int类型的3引用压入操作数栈
func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(3)
}

// 把Int类型的4引用压入操作数栈
func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(4)
}

// 把Int类型的5引用压入操作数栈
func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(5)
}

// 把Long类型的0引用压入操作数栈
func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushLong(0)
}

// 把Long类型的1引用压入操作数栈
func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushLong(1)
}

