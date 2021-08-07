package rtda

import (
	"JVM-Go/heap/rtda/heap"
	"math"
)

//LocalVars 数组实现，数组元素是Slot
type LocalVars []Slot

//newLocalVars 创建指定大小的局部变量表实例
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//SetInt 存放int类型整数
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

//GetInt 获取int类型整数
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

//SetFloat float变量可以先转成int类型，然后按int变量来处理
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

//SetLong Long变量则需要拆成两个int变量,需要2个Slot来存放
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)         //低32bit
	self[index+1].num = int32(val >> 32) //高32bit
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low) //按位或，拼接高32bit和低32bit
}

//SetDouble double变量可以先转成long类型，然后按照long变量来处理
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

//SetRef 存放引用值
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}

//GetRef 获取引用值
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}
