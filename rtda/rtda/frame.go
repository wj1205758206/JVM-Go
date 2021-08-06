package rtda

//Frame 栈帧结构体定义
type Frame struct {
	next         *Frame        //用来实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //保存操作数栈指针
	thread       *Thread       //保存当前线程
	nextPC       int           //保存下一条指令
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
func (self *Frame) GetLocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) GetOperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) GetThread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

func (self *Frame) GetNextPC() int {
	return self.nextPC
}
