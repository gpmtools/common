package exc

import (
	"encoding/json"
	"strings"

	"github.com/cli/go-gh"
)

func NewCmd(s string) Command {
	clean := strings.ReplaceAll(s, "gh ", "")
	ptrs := strings.Split(clean, " ")
	return Command(ptrs)
}

func NewCmdArgs(args ...string) Command {
	return Command(args)
}

type Command []string

func (c Command) Exec() (string, error) {
	out, _, err := gh.Exec(c.StringArray()...)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// ExecUnmarshal unmarshals the output of the command into the provided interface with JSON
func (c Command) ExecUnmarshal(i any) error {
	out, err := c.Exec()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(out), i)
}

// StringArray returns the command as an array of strings
func (c Command) StringArray() []string {
	return []string(c)
}
