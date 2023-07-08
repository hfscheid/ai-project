package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ReadConfigFile tries to read ~/.disco/config.yaml file, if it exists.
// If it doesn`t, the folder and file are created.
func ReadConfigFile() (*Tests, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return nil, fmt.Errorf("Failed to get home dir: %v", err)
    }
    cfgFolder := filepath.Join(home, ".disco") 
    cfgFile := filepath.Join(cfgFolder, "config.yaml")
    tests := Tests{}
    if _, err := os.Stat(cfgFile); err == nil { // File exists
        data, err := os.ReadFile(cfgFile)
        if err != nil {
            return nil, fmt.Errorf("Failed to read disco config file: %v", err)
        }
        err = yaml.Unmarshal(data, &tests)
        if err != nil {
            return nil, fmt.Errorf("Error unmarshaling config.yaml: %v", err)
        }
        return &tests, nil
    } else if errors.Is(err, os.ErrNotExist) { // File doesn`t exist
        err := os.Mkdir(cfgFolder, os.ModePerm)
        if err != nil {
            return nil, fmt.Errorf("Failed to create disco config folder: %v", err)
        }
        data, err := yaml.Marshal(tests)
        if err != nil {
            return nil, fmt.Errorf("Error marshaling Test structure to yaml: %v", err)
        }
        err = os.WriteFile(cfgFile, data, os.ModePerm)
        if err != nil {
            return nil, fmt.Errorf("Failed to create disco config file: %v", err)
        }
        return &tests, nil
    } else { // Unknown error
        return nil, fmt.Errorf("Unable to find or create disco config file: %v", err)
    }
}
