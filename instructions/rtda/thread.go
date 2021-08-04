package rtda

//Thread 结构体定义
type Thread struct {
	pc    int    //PC计数器
	stack *Stack //Java虚拟机栈
}

//newThread 创建线程实例
func newThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

//PC 获取PC计数器的值
func (self *Thread) PC() int {
	return self.pc
}

//SetPC 设置PC计数器的值
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

//PushFrame 压入栈帧
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

//PopFrame 弹出栈帧
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

//CurrentFrame 获取当前栈帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
