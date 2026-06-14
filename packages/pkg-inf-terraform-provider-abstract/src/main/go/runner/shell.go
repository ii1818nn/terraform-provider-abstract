package runner

import (
	"os"
	"os/exec"
	"strings"
)

func Shell(command string, env map[string]string) (Result, error) {
	cmd := exec.Command("sh", "-c", command)

	cmdEnv := os.Environ()
	for k, v := range env {
		cmdEnv = append(cmdEnv, k+"="+v)
	}
	cmd.Env = cmdEnv

	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return Result{Output: strings.TrimSpace(out.String()), ExitCode: exitErr.ExitCode()}, nil
		}
		return Result{}, err
	}

	return Result{Output: strings.TrimSpace(out.String()), ExitCode: 0}, nil
}
