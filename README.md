# ai-project

# Automated testing of software routers

Software routers allow for low-cost, high-flexibility routing in networks. Bugs and vulnerabilities in software routers disrupt connectivity and may have significant impact on end-users and services.

In this project we will extend and improve a framework for testing BGP announcements sent to BIRD and FRR software routers. This project will involve the following tasks:

Extend and update the existing testing framework to:

- Run software routers in Docker containers
- Build containers images for multiple major versions of each software router
- Streamline creation of prefix announcements using ExaBGP or GoBGP
- Automate checking for the correct propagation of the updates
- Run emulations of possible failure scenarios or including uncommon features (e.g., long AS-paths, unassigned BGP attributes, large AS-sets, announcements with many large, extended, and normal BGP communities).

Improvements:

- DISCO: CLI (command line interface) to ... TODO
- ...
