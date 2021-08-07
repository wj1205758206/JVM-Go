package math

import (
	"JVM-Go/instructions/instructions/base"
	"JVM-Go/rtda/rtda"
)

//ISHL int类型左移指令
type ISHL struct {
	base.NoOperandsInstruction
}

//ISHR int类型算数右移指令
type ISHR struct {
	base.NoOperandsInstruction
}

//IUSHR int类型逻辑右移指令
type IUSHR struct {
	base.NoOperandsInstruction
}

//LSHL long类型左移指令
type LSHL struct {
	base.NoOperandsInstruction
}

//LSHR long类型算数右移指令
type LSHR struct {
	base.NoOperandsInstruction
}

//LUSHR long类型逻辑右移指令
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f //int变量只有32位，所以只取v2的前5个比特就足够表示位移位数了
	result := v1 << s
	stack.PushInt(result)
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s) //先把v1转成无符号整数，位移操作之后，再转回有符号整数
	stack.PushInt(result)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f //long变量有64位，所以取v2的前6个比特
	result := v1 >> s
	stack.PushLong(result)
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f //long变量有64位，所以取v2的前6个比特
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
