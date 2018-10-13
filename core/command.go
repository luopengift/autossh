package core

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/luopengift/ssh"
)

// Bash bash, local shell
func Bash(ctx context.Context, command string, writers []io.Writer) error {
	cmdList := strings.Split(command, " ")
	var cmd *exec.Cmd
	if len(cmdList) == 1 {
		cmd = exec.CommandContext(ctx, cmdList[0])
	} else {
		cmd = exec.CommandContext(ctx, cmdList[0], cmdList[1:]...)
	}

	cmd.Stdin = os.Stdin

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
