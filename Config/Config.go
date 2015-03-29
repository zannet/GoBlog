package Config

import (
	"code.google.com/p/gcfg"
	"fmt"
	"os"
	"path/filepath"
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
	// Get the actual path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// If there was an error...
	if err != nil {
		fmt.Println("ERROR: Couldn't get the actual path")
		os.Exit(1)
		return
	}

	// Set the current working directory
	CWD = dir

	sep := string(os.PathSeparator)

	if _, err := os.Stat(CWD + sep + "config.cfg"); os.IsNotExist(err) {
		fmt.Println("ERROR: Couldn't find the configuration file (config.cfg) in the current working directory ("+ (CWD + sep + "config.cfg") +")")
		os.Exit(1)
		return
	}

	var cfg Config
	err = gcfg.ReadFileInto(&cfg, CWD + sep + "config.cfg")
	if err != nil {
		fmt.Println("ERROR: The configuration file has a wrong formatting: " + err.Error())
		os.Exit(1)
		return
	}

	Global = cfg
}