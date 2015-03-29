package Config

import (
	"code.google.com/p/gcfg"
	"fmt"
	"os"
)

var Global Config
var CWD string

type Config struct {
	Service struct {
		Address string
		Port string
	}
	Database struct {
		Address string
		Port string
		User string
		Pass string
		Name string
	}
}

func Load() {
	// Get the working directory
	dir, err := os.Getwd()
	// If there was an error...
	if err != nil {
		fmt.Println("ERROR: Couldn't get the working directory")
		os.Exit(1)
		return
	}

	// Set the working directory globally
	CWD = dir

	// Get the OS separator
	sep := string(os.PathSeparator)

	// Config path
	configPath = CWD + sep + "config.cfg"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("ERROR: Couldn't find the configuration file (config.cfg) in the current working directory ("+ configPath +")")
		os.Exit(1)
		return
	}

	// Create the Config element to be filled with config file data
	var cfg Config
	err = gcfg.ReadFileInto(&cfg, CWD + sep + "config.cfg")
	if err != nil {
		fmt.Println("ERROR: The configuration file has a wrong formatting: " + err.Error())
		os.Exit(1)
		return
	}

	// Set the global configuration
	Global = cfg
}