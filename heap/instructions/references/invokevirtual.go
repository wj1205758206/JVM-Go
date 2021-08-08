package references

import (
	"JVM-Go/heap/instructions/base"
	"JVM-Go/heap/rtda"
	"JVM-Go/heap/rtda/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	cp := frame.GetMethod().GetClass().GetConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	if methodRef.GetName() == "println" {
		stack := frame.GetOperandStack()
		switch methodRef.GetDescriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V":
			fmt.Printf("%c\n", stack.PopInt())
		case "(B)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(S)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(I)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(J)V":
			fmt.Printf("%v\n", stack.PopLong())
		case "(F)V":
			fmt.Printf("%v\n", stack.PopFloat())
		case "(D)V":
			fmt.Printf("%v\n", stack.PopDouble())
		default:
			panic("println: " + methodRef.GetDescriptor())
		}
		stack.PopRef()
	}
}
