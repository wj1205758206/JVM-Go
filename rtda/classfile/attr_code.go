package classfile

//CodeAttribute 变长属性，只存在于method_info结构中,Code属性中存放字节码等方法相关信息
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}
type ExceptionTableEntry struct {
	startPC   uint16
	endPC     uint16
	handlerPC uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *CodeAttribute) GetMaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *CodeAttribute) GetMaxStack() uint {
	return uint(self.maxStack)
}

func (self *CodeAttribute) GetCode() []byte {
	return self.code
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPC:   reader.readUint16(),
			endPC:     reader.readUint16(),
			handlerPC: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
