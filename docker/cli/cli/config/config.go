package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"docker.io/go-docker/api/types"
	"github.com/appcelerator/amp/docker/cli/cli/config/configfile"
	"github.com/appcelerator/amp/docker/cli/cli/config/credentials"
	"github.com/appcelerator/amp/docker/docker/pkg/homedir"
	"github.com/pkg/errors"
)

const (
	// ConfigFileName is the name of config file
	ConfigFileName = "config.json"
	configFileDir  = ".docker"
	oldConfigfile  = ".dockercfg"
)

var (
	configDir = os.Getenv("DOCKER_CONFIG")
)

func init() {
	if configDir == "" {
		configDir = filepath.Join(homedir.Get(), configFileDir)
	}
}

// Dir returns the directory the configuration file is stored in
func Dir() string {
	return configDir
}

// SetDir sets the directory the configuration file is stored in
func SetDir(dir string) {
	configDir = dir
}

// LegacyLoadFromReader is a convenience function that creates a ConfigFile object from
// a non-nested reader
func LegacyLoadFromReader(configData io.Reader) (*configfile.ConfigFile, error) {
	configFile := configfile.ConfigFile{
		AuthConfigs: make(map[string]types.AuthConfig),
	}
	err := configFile.LegacyLoadFromReader(configData)
	return &configFile, err
}

// LoadFromReader is a convenience function that creates a ConfigFile object from
// a reader
func LoadFromReader(configData io.Reader) (*configfile.ConfigFile, error) {
	configFile := configfile.ConfigFile{
		AuthConfigs: make(map[string]types.AuthConfig),
	}
	err := configFile.LoadFromReader(configData)
	return &configFile, err
}

// Load reads the configuration files in the given directory, and sets up
// the auth config information and returns values.
// FIXME: use the internal golang config parser
func Load(configDir string) (*configfile.ConfigFile, error) {
	if configDir == "" {
		configDir = Dir()
	}

	filename := filepath.Join(configDir, ConfigFileName)
	configFile := configfile.New(filename)

	// Try happy path first - latest config file
	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return configFile, errors.Errorf("%s - %v", filename, err)
		}
		defer file.Close()
		err = configFile.LoadFromReader(file)
		if err != nil {
			err = errors.Errorf("%s - %v", filename, err)
		}
		return configFile, err
	} else if !os.IsNotExist(err) {
		// if file is there but we can't stat it for any reason other
		// than it doesn't exist then stop
		return configFile, errors.Errorf("%s - %v", filename, err)
	}

	// Can't find latest config file so check for the old one
	confFile := filepath.Join(homedir.Get(), oldConfigfile)
	if _, err := os.Stat(confFile); err != nil {
		return configFile, nil //missing file is not an error
	}
	file, err := os.Open(confFile)
	if err != nil {
		return configFile, errors.Errorf("%s - %v", confFile, err)
	}
	defer file.Close()
	err = configFile.LegacyLoadFromReader(file)
	if err != nil {
		return configFile, errors.Errorf("%s - %v", confFile, err)
	}
	return configFile, nil
}

// LoadDefaultConfigFile attempts to load the default config file and returns
// an initialized ConfigFile struct if none is found.
func LoadDefaultConfigFile(stderr io.Writer) *configfile.ConfigFile {
	configFile, err := Load(Dir())
	if err != nil {
		fmt.Fprintf(stderr, "WARNING: Error loading config file: %v\n", err)
	}
	if !configFile.ContainsAuth() {
		configFile.CredentialsStore = credentials.DetectDefaultStore(configFile.CredentialsStore)
	}
	return configFile
}
