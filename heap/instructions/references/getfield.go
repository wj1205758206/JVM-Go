package references

import (
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"JVM-Go/heap/instructions/base"
)

//GET_FIELD 获取对象的实例变量值，然后推入操作数栈
type GET_FIELD struct {
	base.Index16Instruction //第一个操作数是uint16索引,第二个操作数是对象引用
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.GetMethod().GetClass().GetConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField() //字段符号引用解析
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	stack := frame.GetOperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException") //弹出对象引用，如果是null则抛出NullPointerException
	}

	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()
	slots := ref.GetFields()
	switch descriptor[0] { //根据字段类型，获取相应的实例变量值，然后推入操作数栈
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
