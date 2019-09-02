package presets

import "../config"
import "io/ioutil"
import "strings"

func Contains(section *config.Section) bool {
	if (section.Preset != "contains") {
		panic("Invalid preset for this section")
	}

	sourceContent, err := ioutil.ReadFile(section.Source)

	if (err != nil) {
		panic("Cannot read " + section.Source + " file")
	}

	return strings.Contains(string(sourceContent), section.Match)
}