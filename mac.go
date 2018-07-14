package mac

import (
	"os/exec"
	"os/user"
	"path/filepath"
)

//////////////////////////////////////////////
/// System
//////////////////////////////////////////////

func Expanduser(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if path[:2] == "~/" {
		path = filepath.Join(dir, path[2:])
	}
	return path
}

//////////////////////////////////////////////
/// RunApplescript
//////////////////////////////////////////////

// RunApplescript runs the ApplesScript contained in script, returning the output and any errors.
func RunApplescript(script string) (string, error) {
	// This implementation is adapted from the unmaintained "github.com/everdev/mack" project.
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.CombinedOutput()
	prettyOutput := strings.Replace(string(output), "\n", "", -1)

	if err != nil {
		return "", Errorf("%v: %v (%v)", err.Error(), prettyOutput, script)
	}

	return prettyOutput, nil
}
