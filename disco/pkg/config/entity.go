package config

import (
	"fmt"
	"strings"
)

type Tests struct {
    TestCases map[string]*TestCase `yaml:"tests"`
}

type TestCase struct {
    Name string `yaml:"name"`
    Network *Network `yaml:"network"`
    Containers []*Container `yaml:"containers"`
}

type Network struct {
    ID string
    Name string `yaml:"name"`
    Subnet string `yaml:"subnet"`
    Gateway string `yaml:"gateway"` 
}

type Container struct {
    ID string
    Name string `yaml:"name"`
    Type Structures `yaml:"type"` 
    ConfigPaths []string `yaml:"configpaths"`
    IP string `yaml:"ip"`
    ExposedPort uint `yaml:"exposedport"`
}

type Structures int 
const (
    EXABGP Structures = iota
    BIRD
    FRR
)
var structToString = map[Structures]string{
    EXABGP: "ExaBGP",
    BIRD: "BIRD",
    FRR: "FRR",
}
var stringToStruct = map[string]Structures{
    "exabgp": EXABGP,
    "bird": BIRD,
    "frr": FRR,
}

func (s Structures) String() string {
    if value, ok := structToString[s]; ok {
        return value
    }
    return fmt.Sprintf("UNKNOWN[%d]", s)
}

func (s Structures) StringToStructure(str string) Structures {
    str = strings.ToLower(str)
    if value, ok := stringToStruct[str]; ok {
        return value
    }
    return -1
}
