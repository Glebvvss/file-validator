package presets

import "../config"
import "io/ioutil"

func Equals(section *config.Section) bool {
	if (section.Preset != "equals") {
		panic("Invalid preset for this section")
	}

	sourceFile, err := ioutil.ReadFile(section.Source)

	if (err != nil) {
		panic("Cannot read " + section.Source + " file")
	}

	matchFile, err := ioutil.ReadFile(section.Match)

	if (err != nil) {
		panic("Cannot read " + section.Source + " file")
	}

	return string(sourceFile) == string(matchFile)
}