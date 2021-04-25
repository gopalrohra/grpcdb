package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Config to hold environment based configuration
var Config = map[string]string{}

// LoadEnvironment loads .env for production and .env.<env_name> for specific environment
func LoadEnvironment() {
	fmt.Println("Initializing environment.")
	//set mode to production, development, staging etc
	mode := os.Getenv("GO_ENV")
	if mode == "" {
		mode = "production"
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)
	env_file := dir + "/.env"
	if mode != "production" {
		env_file = env_file + "." + mode
	}
	data, err := ioutil.ReadFile(env_file)
	if err != nil {
		fmt.Println(err.Error())
	}
	lines := strings.Split(fmt.Sprintf("%s", data), "\n")
	fmt.Println(len(lines))
	for _, line := range lines {
		fmt.Println("Inside loop")
		fmt.Println(line)
		kvPair := strings.Split(line, "=")
		if len(kvPair) == 2 {
			if strings.Trim(kvPair[0], " ") != "" && strings.Trim(kvPair[1], " ") != "" {
				Config[kvPair[0]] = kvPair[1]
			}
		}
	}
	fmt.Println(Config)
}
