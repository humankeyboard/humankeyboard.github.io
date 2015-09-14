package main

import (
	"flag"
	"fmt"
	"github.com/humankeyboard/awsmagic-tools/functions"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of awsplus:\n")
		fmt.Printf("    awsplus help: displays help contents\n")
		fmt.Printf("    awsplus configure: configures awsplus tool\n")
		fmt.Printf("    awsplus create-users -file=myfile.txt: creates users from source file\n")
		fmt.Println()
		flag.PrintDefaults()
	}

	var srcFile string
	var profile string
	configure := false
	var createUsers bool
	var unCreateUsers bool

	flag.StringVar(&srcFile, "file", "users.txt", "file to read from")
	flag.StringVar(&profile, "profile", "default", "profile to use for credentials")
	flag.BoolVar(&configure, "configure", false, "configures awsplus tool")
	flag.BoolVar(&createUsers, "create-users", false, "creates users defined in a file")
	flag.BoolVar(&unCreateUsers, "create-users-undo", false, "deletes users created with create-users as defined in a file")
	flag.Parse()

	if configure {
		functions.Configure()
		return
	}

	if createUsers {
		functions.CreateUsers(profile, srcFile)
	}

	if unCreateUsers {
		functions.UnCreateUsers(profile, srcFile)
	}

}
