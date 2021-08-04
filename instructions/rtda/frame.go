package rtda

//Frame 栈帧结构体定义
type Frame struct {
	next         *Frame        //用来实现链表数据结构
	localVars    LocalVars     //保存局部变量表指针
	operandStack *OperandStack //保存操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
