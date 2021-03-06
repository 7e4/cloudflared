package shell

import (
	"io"
	"os"
	"os/exec"
)

// OpenBrowser opens the specified URL in the default browser of the user
func OpenBrowser(url string) error {
	return getBrowserCmd(url).Start()
}

// Run will kick off a shell task and pipe the results to the respective std pipes
func Run(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	go func() {
		io.Copy(os.Stdout, stdout)
	}()
	return c.Run()
}
