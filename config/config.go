package config

import "strings"
import "io/ioutil"

var sectionDelimeter string = ">>>"
var typeDelimeter    string = "<<<"
var matchDelimeter   string = "--"

func Read(file string) []*Section {
	configContent, err := ioutil.ReadFile(file)

	if (err != nil) {
		panic("Could not read file!")
	}

	return readSections(string(configContent))
}

func readSections(fileContent string) []*Section {
	stringifySections := strings.Split(fileContent, sectionDelimeter)

	var sections []*Section
	for _, stringifySection := range stringifySections {
		stringifySection = trimSection(stringifySection)
		
		if (isEmptyString(stringifySection)) {
			continue
		}

		section := new(Section)
		parsedSource, parsedPreset, parsedMatch := section.parseFromString(stringifySection)

		if (isNotValidPreset(parsedPreset)) {
			panic(parsedPreset + " is not valid preset name")
		}

		section.Source = parsedSource
		section.Preset = parsedPreset
		section.Match  = parsedMatch

		sections = append(sections, section)
	}

	return sections;
}

type Section struct {
	Source string
	Preset string
	Match  string
}
func (b *Section) parseFromString(stringifySection string) (string, string, string) {
	sorceAndPresetAndMatch := strings.SplitN(stringifySection, "<<<", 2)

	source := strings.Trim(sorceAndPresetAndMatch[0], " ")

	presetAndMatch := strings.SplitN(sorceAndPresetAndMatch[1], "--", 2)
	preset := strings.Trim(presetAndMatch[0], " ")
	match  := strings.Trim(presetAndMatch[1], " ")

	return source, preset, match
}

func isNotValidPreset(name string) bool {
	validPresetsList := [2]string{"contains", "equals"}

	for _, presetName := range validPresetsList {
		if (presetName == name) {
			return false
		}
	}

	return true
}

func trimSection(line string) string {
	return strings.Trim(strings.Trim(line, "\n"), " ")
}

func isEmptyString(String string) bool {
	return String == ""
}