# Automated testing of software routers

Software routers allow for low-cost, high-flexibility routing in networks. Bugs and vulnerabilities in software routers disrupt connectivity and may have significant impact on end-users and services.

In this project we will extend and improve a framework for testing BGP announcements sent to [BIRD](https://github.com/CZ-NIC/bird) and [FRR](https://github.com/FRRouting/frr) software routers. This project will involve the following tasks:

- Extend and update the existing testing framework to
- Run software routers in Docker containers
- Build containers images for multiple major versions of each software router
- Streamline creation of prefix announcements using [ExaBGP](https://github.com/Exa-Networks/exabgp)
- Automate checking for the correct propagation of the updates
- Run emulations of possible failure scenarios or including uncommon features (e.g., long AS-paths, unassigned BGP attributes, large AS-sets, announcements with many large, extended, and normal BGP communities).

Project contents:

- `/routers`: BIRD and FRR router configuration files. The directories `/routers/bird/confs` and `/routers/frr/confs` are mounted
as volumes into each router specific Docker container when running the application.  
- `/bgp-announcer`: Configuration file for ExaBGP and a Python script that receives requests to annouce new BGP routes.
This instance of ExaBGP is responsible for announcing new routes in the routers.
- `/bgp-listener`: Configuration file for ExaBGP and a Python script that receives requests to gather the execution logs.
This instace of ExaBGP is responsible for receiving information regarding the annouced routes.
- `/disco`: CLI application used for managing test cases (creation, execution, etc.).

## DISCO CLI

DISCO is a CLI application that serves as a helper to create, store, manage and run software router tests.  
In addition to storing defined tests in a config file, it is able to build and run all the Docker containers and networks required for each test case,
as well as sending requests to specific containers, such as ExaBGP ones, so new routes can be announced.

### Install

DISCO uses [Docker SDK](https://docs.docker.com/engine/api/sdk/) to manage containers, images and networks.
So before installing DISCO be sure to have the [Docker Engine](https://docs.docker.com/engine/install/) installed.  

The CLI was build using Golang, so the final application is just a binary file. To download it, check the most recent released version in this repo's [release page](https://github.com/hfscheid/ai-project/releases).  
After downloading it, add the directory in which the binary is stored to your `$PATH`, so it will be available in your terminal.  
To test it, just run the `disco` commmand. The output should be similar to this:

```sh
disco - tool for creating, configuring and testing software routers. Use 'disco help' to list all available commands

Usage:
  disco [flags]
  disco [command]

Available Commands:
  annouce     Command for annoucing BGP routes
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  test        Command for managing tests

Flags:
  -h, --help      help for disco
  -v, --version   version for disco

Use "disco [command] --help" for more information about a command.
```

The first time DISCO is run, it creates the folder and file `~/.disco/config.yaml`, which is used to store information that allows the application to work properly.  

### Usage

The application has two main commands, `disco test [command]` for managing tests and `disco annouce [container_name] [announcement]` that sends a POST request to annouce new routes.  

#### `disco test`

| Command                       | Description                                                                                                                                                                                                                                                | Example                               |
| ----------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------- |
| `create </path/to/test.yaml>` | Creates a new test case defined by the config file (see config file example below)                                                                                                                                                                         | `disco test create test-example.yaml` |
| `list`                        | Lists all available tests                                                                                                                                                                                                                                  | `disco test list`                     |
| `select <test_name>`          | Selects an existing test. The commands `run`, `stop`, `describe` and `delete` are applied only in the currently selected test                                                                                                                              | `disco test select ExampleTest`       |
| `describe`                    | Prints a detailed description of the currently selected test                                                                                                                                                                                               | `disco test describe`                 |
| `run [--flag]`                | Executes the selected test (builds and runs containers, networks, images, etc.). Optionally the flag `--watch` can be used to attach to the containers' logs, similar to `docker compose up` (the logs will be printed to stdout once the test is stopped) | `disco test run`                      |
| `stop`                        | Stops the currently selected test, if it's running                                                                                                                                                                                                         | `disco test stop`                     |
| `delete`                      | Deletes the currently selected test. The test is removed from DISCO's config file, and all the Docker structures related to it are also deleted (containers, networks)                                                                                     | `disco test delete`                   |

#### `disco annouce`

For this command to work, the target ExaBGP container must have an `exposedport` (see config file example below) and a script that defines and runs a server which receives
POST requests. An example of this server and its usage can be found [here](https://github.com/hfscheid/ai-project/tree/master/bgp-announcer).  

| Command                                     | Description                                                      | Example                                                                    |
| ------------------------------------------- | ---------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `<container_name> <announcement_command>` | Sends request to server which will create the given announcement | `disco announce exa-announcer "annouce route 100.10.0.0/24 next-hop self"` |

### Test Case config file

To create a new test, the refence to a YAML file containing the test definition must be passed to `disco test create` command.  
Below is the base-structure of a test definition, a working example can be found [here](https://github.com/hfscheid/ai-project/tree/master/test-example.yaml).

```yaml
name: <test name>
network: # Network which all containers will be connected to
  name: <network name>
  subnet: <IPAM subnet>
  gateway: <IPAM gateway>
containers: # List of containers to be created during the test
  - name: <container name>
    image:
      name: <docker image name>
      version: <docker image version>
    exposedport: <container port that will be exposed and mapped to localhost>
    ip: <container IP in the network>
    configpaths: # Path to hosts' directories that will be mounted in the container
      - <path to host directory>:<path to container directory>
```
