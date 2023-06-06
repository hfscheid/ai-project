# FRR useful references

## Links

- [GitHub repo](https://github.com/FRRouting/frr)
- [Docker Images](https://quay.io/repository/frrouting/frr?tab=tags)
- [FRR Docker Image usage](http://docs.frrouting.org/projects/dev-guide/en/latest/building-frr-for-alpine.html#usage)
- [Basic Setup](http://docs.frrouting.org/en/stable-8.5/setup.html#daemons-configuration-file)
    - This doc contains a config file example for setting up all of FRR supported daemons. We'll probably just use the `bgpd` daemon (meaning that this file probably won't be changed once set up)
- [BGP daemon config](http://docs.frrouting.org/en/stable-8.5/bgp.html#miscellaneous-configuration-examples)
    - This lists examples of configs for the BGP daemon. We'll most likely have a bunch of different configs for testing
- [VTYSH reference](https://docs.frrouting.org/projects/dev-guide/en/latest/vtysh.html)

## Ideas

## Setting up the configs

Based on the "FRR Docker Image usage", the way we can set this up is by having a config file for FRR in the repo.
We then can create a `docker-compose.yaml` file in which we'll mount the local config file into the container.
The same goes to the BGP daemon config file.  
[Docker Compose volume mount reference](https://docs.docker.com/compose/compose-file/compose-file-v3/#short-syntax-3).

- [FRR Config Files](https://github.com/FRRouting/frr/tree/master/tools/etc/frr)
- [FRR docker-compose example](https://github.com/ksator/frrouting_demo/tree/master)
