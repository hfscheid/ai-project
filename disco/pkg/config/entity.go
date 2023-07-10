package config

type Tests struct {
    TestCases map[string]*TestCase `yaml:"tests"`
    SelectedTest *TestCase `yaml:"selectedtest"`
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
    Image DockerImage `yaml:"image"` 
    ConfigPaths []string `yaml:"configpaths"`
    IP string `yaml:"ip"`
    ExposedPort uint `yaml:"exposedport"`
}

type DockerImage struct {
    Name string `yaml:"name"`
    Version string `yaml:"version"`
}
