package adapters

import "os/exec"

type ExecRunner struct{}

func (ExecRunner) CombinedOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}