package classfile

//ConstantIntegerInfo 结构体定义,使用4字节存储整数常量
type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

func (self *ConstantIntegerInfo) GetValue() int32 {
	return self.val
}

//ConstantFloatInfo 结构体定义,使用4字节存储单精度浮点数常量
type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = float32(bytes)
}

func (self *ConstantFloatInfo) GetValue() float32 {
	return self.val
}

//ConstantLongInfo 结构体定义,使用8字节存储Long类型整数常量
type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

func (self *ConstantLongInfo) GetValue() int64 {
	return self.val
}

//ConstantDoubleInfo 结构体定义，使用8字节存储双精度浮点数
type ConstantDoubleInfo struct {
	val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = float64(bytes)
}

func (self *ConstantDoubleInfo) GetValue() float64 {
	return self.val
}
