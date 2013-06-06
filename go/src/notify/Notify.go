package notify

import (
	"bytes"
	"checks"
)

func BuildOutOfDateMessage(outOfDatePrograms []checks.OutOfDateProgram) string {
	var message bytes.Buffer
	for _, program := range outOfDatePrograms {
		message.WriteString(program.Name + " appears to have a new release! Version " + program.LatestVersion + " upstream is newer than current AUR version " + program.AurVersion + "  ")
		message.WriteString("https://aur.archlinux.org/packages/" + program.Name + "\n")
	}

	return message.String()

}
