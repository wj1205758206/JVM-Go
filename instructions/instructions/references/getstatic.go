package references

import (
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"JVM-Go/instructions/instructions/base"
)

//GET_STATIC 获取静态变量的值到操作数栈
type GET_STATIC struct {
	base.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.GetMethod().GetClass().GetConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.GetClass() //如果声明字段的类还没有初始化好，也需要先初始化
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError") //如果解析后的字段不是静态字段，也要抛出异常
	}

	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()
	slots := class.GetStaticVars()
	stack := frame.GetOperandStack()
	switch descriptor[0] { //根据字段类型，从静态变量中取出相应的值，然后推入操作数栈顶
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}

}
