package core

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/luopengift/ssh"
)

// Bash bash, local shell
func Bash(ctx context.Context, command string, writers []io.Writer) error {
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", command)
	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if writers == nil {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		outs := []io.Writer{os.Stdout}
		errs := []io.Writer{os.Stderr}

		outs = append(outs, writers...)
		errs = append(errs, writers...)

		go ssh.Copy(outs, stdout)
		go ssh.Copy(errs, stderr)
	}
	return cmd.Run()
}
