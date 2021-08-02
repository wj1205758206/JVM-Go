package main

import (
	"JVM-Go/searchclassfile/classpath"
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
	fmt.Printf("classpath: %s\nclass: %v\nargs: %v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data: %v\n", classData)
}
