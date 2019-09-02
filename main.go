package main

import "./config"
import "./presets"
import "./print"

var configFile string = "config"

func main() {
	sections := config.Read(configFile)

	for _, section := range sections {
		switch section.Preset {
		case "contains":
			print.Message(section, presets.Contains(section))
		case "equals":
			print.Message(section, presets.Equals(section))
		}
	}
}