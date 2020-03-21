package mac

import (
	. "fmt"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func Expanduser(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if path[:2] == "~/" {
		path = filepath.Join(dir, path[2:])
	}
	return path
}

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

func GetFrontMostApplication() (string, error) {
	script := `path to frontmost application as Unicode text`
	path, err := TellSystemEvents(script)
	if err != nil {
		return path, err
	}
	nameSplit := strings.Split(path, ":")
	name := nameSplit[len(nameSplit)-2]
	appSplit := strings.Split(name, ".")
	app := appSplit[len(appSplit)-2]

	return app, err
}

func TellSystemEvents(script string) (string, error) {
	out, err := Tell("System Events", script)
	return out, err
}

func Tell(application string, commands ...string) (string, error) {
	return RunApplescript(buildTell(application, commands...))
}

func wrapInQuotes(text string) string {
	return "\"" + text + "\""
}

// Parse the Tell options and build the command
func buildTell(application string, commands ...string) string {
	application = wrapInQuotes(application)
	args := []string{"tell application", application, "\n"}
	for _, command := range commands {
		args = append(args, command, "\n")
	}
	args = append(args, "end", "tell")
	return build(args...)
}

func build(params ...string) string {
	var validParams []string

	for _, param := range params {
		if param != "" {
			validParams = append(validParams, param)
		}
	}

	return strings.Join(validParams, " ")
}
