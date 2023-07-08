package config

import "fmt"

type Tests struct {
    TestCases map[string]TestCase `yaml:"tests"`
}

type TestCase struct {
    Network Network `yaml:"network"`
    Containers []Container `yaml:"containers"`
}

type Network struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Subnet string `yaml:"subnet"`
    Gateway string `yaml:"gateway"` 
}

type Container struct {
    ID string `yaml:"id"`
    Name string `yaml:"name"`
    Type Structures `yaml:"type"` 
    ConfigPath string `yaml:"configpath"`
    IP string `yaml:"ip"`
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

func (s Structures) String() string {
    if value, ok := structToString[s]; ok {
        return value
    }
    return fmt.Sprintf("UNKNOWN[%d]", s)
}
