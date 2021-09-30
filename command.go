// Define a type Executable and one method to run it (handle stdout, err)
// Executable.Run(cmd string, arg[]string)
// Error handle:
// 			pathexec not found
//			Error starting Cmd
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Executable struct {
	Name   string
	Params []string
}

func printCommand(c *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(c.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
		os.Exit(1)
	}
}

func printOutput(o []byte) {
	fmt.Printf("==> Output: %s\n\n", string(o))
}

func runCommand(c *exec.Cmd) string {
	// Buffer output
	cOutput := &bytes.Buffer{}
	c.Stdout = cOutput

	// Run command
	err := c.Run()
	printError(err)

	// Display output and return it
	printCommand(c)

	o := cOutput.Bytes()
	printOutput(o)

	return string(o)
}

func (e Executable) Run() string {
	//Find executable path
	pathexec, err := exec.LookPath(e.Name)

	//Exit program if pathexec not found
	printError(err)

	//Execute cmd with all given arguments
	if n := len(e.Params); n > 0 {
		r := exec.Command(pathexec, e.Params...)
		o := runCommand(r)
		return string(o)
	}

	//Execute cmd without argument
	r := exec.Command(pathexec)
	o := runCommand(r)
	return string(o)
}
