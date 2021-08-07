package heap

import "math"

//Slot 存放类变量和实例变量的基本单位，既可以存放一个int值也可以存放一个引用值
type Slot struct {
	num int32   //存放整数
	ref *Object //存放引用
}

type Slots []Slot

//newSlots 创建指定大小的Slots实例
func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

//SetInt 存放int类型整数
func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}

//GetInt 获取int类型整数
func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

//SetFloat float变量可以先转成int类型，然后按int变量来处理
func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

//SetLong Long变量则需要拆成两个int变量,需要2个Slot来存放
func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)         //低32bit
	self[index+1].num = int32(val >> 32) //高32bit
}

func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low) //按位或，拼接高32bit和低32bit
}

//SetDouble double变量可以先转成long类型，然后按照long变量来处理
func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

//SetRef 存放引用值
func (self Slots) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

//GetRef 获取引用值
func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}
