package presets

import "../config"
import "io/ioutil"
import "path/filepath"

func Equals(section *config.Section) bool {
	if (section.Preset != "equals") {
		panic("Invalid preset for this section")
	}

	absSource, _ := filepath.Abs(section.Source)
	sourceFile, err := ioutil.ReadFile(absSource)

	if (err != nil) {
		panic("Cannot read " + section.Source + " file")
	}

	matchFile, err := ioutil.ReadFile(section.Match)

	if (err != nil) {
		panic("Cannot read " + section.Source + " file")
	}

	return string(sourceFile) == string(matchFile)
}