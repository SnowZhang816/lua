package main

import "fmt"
import "strings"
import "main/cmd"
import "main/classpath"

func main() {
	lCmd := cmd.ParseCmd()
	fmt.Println(lCmd)
	if lCmd.GetVersionFlag() {
		fmt.Println("version 0.0.1")
	} else if lCmd.GetHelpFlag() || lCmd.GetClass() == "" {
		cmd.PrintUsage()
	} else {
		startJVM(lCmd)
	}
}

func startJVM(lCmd *cmd.Cmd) {
	cp := classpath.Parse(lCmd.GetXjreOption(), lCmd.GetCpOption())

	fmt.Printf("classpath:%v class:%v args:%v\n", cp, lCmd.GetClass(), lCmd.GetArgs())

	className := strings.Replace(lCmd.GetClass(), ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", lCmd.GetClass())
		return
	}

	fmt.Printf("class data:%v\n", classData)
}
