package main

import "fmt"

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
	fmt.Printf("searchclassfile: %s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
