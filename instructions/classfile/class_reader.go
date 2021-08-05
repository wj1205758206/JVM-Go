package classfile

import "encoding/binary"

//ClassReader 封装[]byte类型
type ClassReader struct {
	data []byte
}

//readUint8 读取u1类型数据
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//readUint16 读取u2类型数据
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//readUint32 读取u4类型数据
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

//readUint64 读取uint64(8字节)的数据
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

//readUint16Table 读取表结构数据
func (self *ClassReader) readUint16Table() []uint16 {
	count := self.readUint16() //读取计数器数据，表示有多少个表项
	table := make([]uint16, count)
	for i := range table {
		table[i] = self.readUint16() //读取所有的表项
	}
	return table
}

//readBytes 读取指定数量的字节数
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
