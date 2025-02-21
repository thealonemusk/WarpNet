package service

import (
	"os/exec"
	"path/filepath"

	process "github.com/thealonemusk/go-processmanager"
)

// NewProcessController returns a new process controller associated with the state directory
func NewProcessController(statedir string) *ProcessController {
	return &ProcessController{stateDir: statedir}
}

// ProcessController syntax sugar around go-processmanager
type ProcessController struct {
	stateDir string
}

// Process returns a process associated within binaries inside the state dir
func (a *ProcessController) Process(state, p string, opts ...process.Option) *process.Process {
	return process.New(
		append(opts,
			process.WithName(a.BinaryPath(p)),
			process.WithStateDir(filepath.Join(a.stateDir, "proc", state)),
		)...,
	)
}

// BinaryPath returns the binary path of the program requested as argument.
// The binary path is relative to the process state directory
func (a *ProcessController) BinaryPath(b string) string {
	return filepath.Join(a.stateDir, "bin", b)
}

// Run simply runs a command from a binary in the state directory
func (a *ProcessController) Run(command string, args ...string) (string, error) {
	cmd := exec.Command(a.BinaryPath(command), args...)
	out, err := cmd.CombinedOutput()

	return string(out), err
}
