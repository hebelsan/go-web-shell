package handlers

import (
	"bytes"
	"os/exec"
)

type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type CmdOut struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

func runCmd(command Command) (out CmdOut, err error) {
	cmd := exec.Command(command.Name, command.Args...)

	// Set output to Byte Buffers
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()
	if err != nil {
		return
	}
	out.Stdout = outb.String()
	out.Stderr = errb.String()
	return
}
