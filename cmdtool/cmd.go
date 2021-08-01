package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd /*命令行选项和参数的结构体定义*/
type Cmd struct {
	helpFlag    bool     //帮助选项，输出帮助信息
	versionFlag bool     //版本选项，输出版本信息
	cpOption    string   //目录和 zip/jar 文件的类搜索路径，用于搜索类文件
	class       string   //主类名，用于指定主类
	args        []string //参数信息
}

/*Go语言内置了`flag`包可以用来处理命令行选项，`os`包中定义了`Args`变量，其中存放传递给命令行的全部参数*/
func parseCmd() *Cmd {
	cmd := &Cmd{}
	//编写printUsage()函数，用来输出到控制台
	flag.Usage = printUsage
	//Var函数用来设置需要解析的选项
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//调用Parse函数解析选项
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0] //第一个参数是主类名
		cmd.args = args[1:] //从第二个参数开始往后，都是main()方法的参数
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
