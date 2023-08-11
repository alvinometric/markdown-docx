package main

import (
	"fmt"
	"os"
	"os/exec"
)


func runCommand(){
	cmd := exec.Command("ls", "-a")

	out, err := cmd.Output()

	if err != nil{
		fmt.Println("Error running Pandoc");
		return
	}

	fmt.Println(string(out))
}


func main() {

	args := os.Args[1:]

	runCommand()

	fmt.Println(args)
}
