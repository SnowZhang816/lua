package main

import "fmt"
import "main/cmd"

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
	fmt.Printf("classpath:%s class:%s args:%v\n",
	lCmd.GetCpOption(), lCmd.GetClass(), lCmd.GetArgs())
}
