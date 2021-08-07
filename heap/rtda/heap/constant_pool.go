package heap

import (
	"JVM-Go/heap/classfile"
	"fmt"
)

//Constant 定义常量接口
type Constant interface{}

//ConstantPool 常量池存放两类信息，字面量(整数、浮点数和字符串) 和 符号引用(类符号引用、字段符号引用、方法符号引用和接口方法符号引用)
type ConstantPool struct {
	class  *Class
	consts []Constant
}

//newConstantPool 把classFile文件中的常量池信息转换成运行时常量池
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) { //核心逻辑就是把[]classfile.ConstantInfo转换成[]heap.Constant
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.GetValue()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.GetValue()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.GetValue()
			i++ //long类型占据两个位置
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.GetValue()
			i++ //double类型占据两个位置
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String() //如果是字符串常量，直接调用ConstantStringInfo的String()方法，从常量池中根据索引取出字符串常量
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldRerInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldRerInfo)
		case *classfile.ConstantMethodrefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfaceMethodRefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, interfaceMethodRefInfo)
		}
	}
	return rtCp
}

//GetConstant 根据索引返回常量池中的常量
func (self *ConstantPool) GetConstant(index uint) Constant {
	c := self.consts[index]
	if c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
