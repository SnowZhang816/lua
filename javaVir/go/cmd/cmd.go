package cmd

import "flag"
import "fmt"
import "os"

// java [-options] class [args...]
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	fmt.Println(cmd.helpFlag, cmd.versionFlag, cmd.cpOption, cmd.class, cmd.args)
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}

func (self *Cmd) GetHelpFlag() bool {
	return self.helpFlag
}

func (self *Cmd) GetVersionFlag() bool {
	return self.versionFlag
}

func (self *Cmd) GetClass() string {
	return self.class
}

func (self *Cmd) GetCpOption() string {
	return self.cpOption
}

func (self *Cmd) GetArgs() []string {
	return self.args
}



