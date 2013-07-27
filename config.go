package boxcars

import (
	"io/ioutil"
	"encoding/json"
)

type Config map[string]map[string]string
type ConfigRaw map[string]interface{}

var filename string

func ReadConfig () {
	debug("Reading %s", filename)

	configRaw := make(ConfigRaw)
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		debug("Failed to read %s", filename)
		return
	}

	err = json.Unmarshal(content, &configRaw)

	if err != nil {
		debug("Failed to parse %s", filename)
		return
	}

	reloadConfig(configRaw)
}

func SetFilename (input string) {
	filename = input
}

func reloadConfig (raw ConfigRaw) {
	debug("Setting up the config document")

	config := make(Config)

	for hostname, options := range raw {
		switch t:= options.(type) {
		case string:
			config[hostname] = make(map[string]string)
			config[hostname]["*"] = t
		case map[string]interface{}:
			config[hostname] = make(map[string]string)

			for path, uri := range t {
				if str, ok := uri.(string); ok {
					config[hostname][path] = str
				}
			}
		}
	}

	SetupSites(config)
}
