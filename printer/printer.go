package printer

import "fmt"
import "../config"

func Message(section *config.Section, success bool) {
	if (success) {
		fmt.Println("\x1b[32m+ Valid\x1b[0m: file \"" + section.Source + "\" " + section.Preset + " \"" + section.Match + "\" ")
		return
	}

	fmt.Println("\x1b[31m+ Invalid\x1b[0m: file \"" + section.Source + "\" not " + section.Preset + " \"" + section.Match + "\" ")
}