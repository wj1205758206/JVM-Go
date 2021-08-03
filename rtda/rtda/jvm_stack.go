package rtda

//Stack 栈结构体定义
type Stack struct {
	maxSize uint   //栈的最大容量
	size    uint   //当前栈的大小
	_top    *Frame //保存栈顶指针
}

//newStack 创建一个栈实例
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

//push 压入栈帧
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.next = self._top //链表头插法
	}
	self._top = frame //新插入的frame成为栈顶栈帧
	self.size++
}

//pop 弹出栈帧
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.next
	top.next = nil
	self.size--
	return top
}

//top 获取栈顶栈帧
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
