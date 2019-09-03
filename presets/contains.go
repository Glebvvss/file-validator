package presets

import "../config"
import "io/ioutil"
import "strings"
import "path/filepath"

func Contains(section *config.Section) bool {
	if (section.Preset != "contains") {
		panic("Invalid preset for this section")
	}

	absSource, _ := filepath.Abs(section.Source)
	sourceContent, err := ioutil.ReadFile(absSource)

	if (err != nil) {
		panic("Cannot read " + section.Source + " file")
	}

	return strings.Contains(string(sourceContent), section.Match)
}