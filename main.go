package main

import "./config"
import "./presets"
import "./printer"

var configFile string = "./config"

func main() {
	sections := config.Read(configFile)

	for _, section := range sections {
		if (section.Preset == "contains") {
			printer.Message(section, presets.Contains(section))
		} else if (section.Preset == "equals") {
			printer.Message(section, presets.Equals(section))
		}
	}
}