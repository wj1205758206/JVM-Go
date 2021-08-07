package heap

type Object struct {
	class  *Class //存放对象的Class指针
	fields Slots  //存放实例变量
}
