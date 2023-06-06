# BIRD useful references

## Links

- [GitLab repo](https://gitlab.nic.cz/labs/bird)
- [Docker images](https://hub.docker.com/r/pierky/bird/tags)
- BIRD Docker Image usage: it's in the "Overview" part of the DockerHub page
- [BIRD Setup](https://bird.network.cz/doc/bird-3.html)
    - [Example 1](https://github.com/pierky/bgp-large-communities-playground/blob/master/bird/bird.conf) - easiest to understand
    - [Example 2](https://gitlab.nic.cz/labs/bird/-/blob/master/doc/bird.conf.example2) - "Official" example

## Ideas

## Setting up the configs

Based on the "BIRD Docker Image usage", the way we can set this up is by having a config file for BIRD in the repo.
We then can create a `docker-compose.yaml` file in which we'll mount the local config file into the container.
[Docker Compose volume mount reference](https://docs.docker.com/compose/compose-file/compose-file-v3/#short-syntax-3).

- [BIRD Config File](https://github.com/pierky/bgp-large-communities-playground/blob/master/bird/bird.conf)

