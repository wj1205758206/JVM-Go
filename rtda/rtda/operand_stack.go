package rtda

import "math"

//OperandStack 操作数栈结构体定义
type OperandStack struct {
	size  uint32 //用于记录栈顶位置
	slots []Slot //基本单位是Slot变量槽
}

//newOperandStack 创建一个指定大小的操作数栈实例
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

//PushInt 向操作数栈中压入int类型数据
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

//PopInt 弹出操作数栈顶int类型数据
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

//PushFloat float变量还是先 转成int类型，然后按int变量处理
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

//PushLong 把long变量推入栈顶时，要拆成两个int变量
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

//PopLong 弹出时，先弹出 两个int变量，然后组装成一个long变量
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high<<32) | int64(low)
}

//PushDouble double变量先转成long类型，然后按long变量处理
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

//PushRef 压入引用类型数据
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

//PopRef 弹出引用类型数据
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	/*弹出引用后，把Slot结构体的ref字段设置成nil，这样做是为了帮助Go的垃圾收集器回收Object结构体实例*/
	self.slots[self.size].ref = nil
	return ref
}

func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
