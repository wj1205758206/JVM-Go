package main

import (
	"JVM-Go/rtda/classfile"
	"JVM-Go/rtda/classpath"
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
	//fmt.Printf("classpath: %s\nclass: %v\nargs: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)

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
