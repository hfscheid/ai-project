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

// WriteToConfigFile tries to add a test to ~/.disco/config.yaml file, if it exists.
// If it doesn`t, return an error.
func WriteToConfigFile(tests *Tests) error {
    home, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf("Failed to get home dir: %v", err)
    }
    cfgFolder := filepath.Join(home, ".disco") 
    cfgFile := filepath.Join(cfgFolder, "config.yaml")
    if _, err := os.Stat(cfgFile); err == nil { // File exists
        f, err := os.OpenFile(cfgFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
        if err != nil {
            return fmt.Errorf("Failed to open disco config file: %v", err)
        }
        data, err := yaml.Marshal(tests)
        if err != nil {
            return fmt.Errorf("Error marshaling new test: %v", err)
        }
        if _, err := f.WriteString(string(data)); err != nil {
            return fmt.Errorf("Error updating disco config file: %v", err)
        }
        return f.Close()
    } else {
        return fmt.Errorf("Unable to add test to disco config file: %v", err)
    }
}

func ReadTestConfig(file string) (*TestCase, error) {
    if _, err := os.Stat(file); err == nil { // File exists
        testCase := TestCase{}
        data, err := os.ReadFile(file)
        if err != nil {
            return nil, fmt.Errorf("Failed to read test config file: %v", err)
        }
        err = yaml.Unmarshal(data, &testCase)
        if err != nil {
            return nil, fmt.Errorf("Error unmarshaling %s: %v", file, err)
        }
        return &testCase, nil
    } else {
        return nil, fmt.Errorf("Unable to find test config: %v", err)
    }
}
