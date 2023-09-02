package songbirds

import (
	"bytes"
	"os/exec"
)

func RunCommand(command ...string) (stdout, stderr string, err error) {
	execCommand := exec.Command(command[0], command[1:]...)

	var outBuf, errBuf bytes.Buffer

	execCommand.Stdout = &outBuf
	execCommand.Stderr = &errBuf

	err = execCommand.Start()

	if err != nil {
		return
	}

	cmdChan := make(chan error, 1)
	go func() {
		cmdChan <- execCommand.Wait()
	}()

	select {
	case err = <-cmdChan:
		stdout = outBuf.String()
		stderr = errBuf.String()

	}
	return
}
