package references

import (
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"JVM-Go/heap/instructions/base"
)

//PUT_FIELD 给实例变量赋值
type PUT_FIELD struct {
	base.Index16Instruction //它需要三个操作数。前两个操作数是常量池索引和变量值,第三个操作数是 对象引用，从操作数栈中弹出
}

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetClass()
	cp := currentClass.GetConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.GetClass()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError") //解析后的字段必须是实例字段，否则抛出异常
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.GetName() != "<init>" {
			panic("java.lang.IllegalAccessError") //如果是final字段，则只能在构造函数中初始化，否则抛出异常
		}
	}

	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()
	stack := frame.GetOperandStack()
	switch descriptor[0] { //根据字段类型从操作数栈中弹出相应的变量值
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException") //如果引用是null，需要抛出著名的空指针异常
		}
		ref.GetFields().SetInt(slotId, val) //通过引用给实例变量赋值
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetLong(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetFloat(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetRef(slotId, val)
	}
}
