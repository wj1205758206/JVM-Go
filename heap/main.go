package main

import (
	"JVM-Go/rtda/classfile"
	"JVM-Go/rtda/classpath"
	"JVM-Go/rtda/rtda"
	"fmt"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 1.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpreter(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

	/*frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.GetLocalVars())
	testOperandStack(frame.GetOperandStack())*/

	/*cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath: %s\nclass: %v\nargs: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)*/

}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.GetMethods() {
		if m.GetName() == "main" && m.GetDescriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.GetMajorVersion(), cf.GetMinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.GetConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.GetAccessFlags())
	fmt.Printf("this class: %v\n", cf.GetClassName())
	fmt.Printf("super class: %v\n", cf.GetSuperClassName())
	fmt.Printf("interfaces: %v\n", cf.GetInterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.GetFields()))
	for _, f := range cf.GetFields() {
		fmt.Printf(" %s\n", f.GetName())
	}
	fmt.Printf("methods count: %v\n", len(cf.GetMethods()))
	for _, m := range cf.GetMethods() {
		fmt.Printf(" %s\n", m.GetName())
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
