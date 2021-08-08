package references

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
)

//PUT_STATIC 给类的某个静态变量赋值
type PUT_STATIC struct {
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.GetMethod()                      //从当前栈帧获取当前方法
	currentClass := currentMethod.GetClass()                //从当前方法获取所属的当前类
	cp := currentClass.GetConstantPool()                    //从当前类获取该类的运行时常量池
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef) //根据指令的索引，从当前运行时常量池中获取字段符号引用
	field := fieldRef.ResolvedField()                       //解析字段符号引用
	class := field.GetClass()                               //如果声明字段的类还没有被初始化，则需要先初始化该类
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError") //如果解析后的字段是实例字段而非静态字段，则抛出异常
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.GetName() != "<clinit>" {
			panic("java.lang.IllegalAccessError") //如果是final字段，则实际操作的是静态常量，只能在类初始化方法中给它赋值
		}
	}

	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()
	slots := class.GetStaticVars()
	stack := frame.GetOperandStack()
	switch descriptor[0] { //根据字段类型从操作数栈中弹出相应的值，然后赋给静态变量
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
