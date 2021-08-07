package heap

const (
	ACC_PUBLIC       = 0x0001 //类、字段、方法
	ACC_PRIVATE      = 0x0002 //字段、方法
	ACC_PROTECTED    = 0x0004 //字段、方法
	ACC_STATIC       = 0x0008 //字段、方法
	ACC_FINAL        = 0x0010 //类、字段、方法
	ACC_SUPER        = 0x0020 //类
	ACC_SYNCHRONIZED = 0x0020 //方法
	ACC_VOLATILE     = 0x0040 //字段
	ACC_BRIDGE       = 0x0040 //方法
	ACC_TRANSIENT    = 0x0080 //字段
	ACC_VARARGS      = 0x0080 //方法
	ACC_NATIVE       = 0x0100 //方法
	ACC_INTERFACE    = 0x0200 //类
	ACC_ABSTRACT     = 0x0400 //类、方法
	ACC_STRICT       = 0x0800 //方法
	ACC_SYNTHETIC    = 0x1000 //类、字段、方法
	ACC_ANNOTATION   = 0x2000 //类
	ACC_ENUM         = 0x4000 //类、字段
)
