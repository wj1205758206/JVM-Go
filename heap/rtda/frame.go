package rtda

import "JVM-Go/heap/rtda/heap"

//Frame 栈帧结构体定义
type Frame struct {
	next         *Frame        //用来实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //保存操作数栈指针
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.GetMaxLocals()),
		operandStack: newOperandStack(method.GetMaxStack()),
	}
}
func (self *Frame) GetLocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) GetOperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) GetMethod() *heap.Method {
	return self.method
}

func (self *Frame) GetThread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) GetNextPC() int {
	return self.nextPC
}
