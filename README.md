# Twitter API with Golang

<p align="center">
    <img src= "https://camo.githubusercontent.com/b58f4fe5bb32f7d78931303d9a98c0fc6dac1a798fad41043fec52d8b066bbf5/68747470733a2f2f73746f726167652e676f6f676c65617069732e636f6d2f6467687562626c652f676f706865722d6f6e2d626972642e706e67" 
    width="200" height="200"/><br>
</p>

---
## Motivation

Build an API very similar to Twitter, but it has far fewer features.Just have only two pages, the homepage and the user profile page and you can create your user and posts.

---
## Dependencies

### Mandatory

- Golang
- Docker
- Make
- PostgreSQL

### Opcional

- [x] VsCode extension OpenAPI (swagger) Editor to view API Documentation
- [x] VsCode extension Draw.io to view the Modelling Documentation
---

## How to run

### Running standalone

run only the API 

```sh
make run
```

### Running with Database + API

To run the API together with the database:
```sh
drun-docker
```
### How to test
Run tests and the coverage
```sh
make run-tests
```
---
### Documentation available

All our documentation is present inside the [docs](./docs) folder.

Inside the docs are two folders, api and modeling, inside [api](./docs/api/) you have all our available endpoints documented in the form of an [OpenAPI swagger](./docs/api/swagger.yaml). Inside the [modeling](./docs/api/) folder there are draw.io files with the [data modeling](./docs/modelling/DatabaseModelling.drawio) and the [creation roadmap](./docs/modelling/StriderBackendProjectRoadmap.drawio).

---

# Critique

This project has several technical debts that need to be corrected throughout its implementation, where the initial scope would not be enough for it.

The first point would be a greater test coverage with an end-to-end process, validating the entire user creation process, post, reply and finally the generation of the homepage and user profile.

Another point would be the scalability for searching and paginating the Posts and Homepage. If the process scales significantly, it would be valid to separate a service just for the homepage, not hindering the application with so much response time and search in the database.

Adding a cache like redis would gladly increase the response speed for information that has low rate of change like the userprofile.

Thinking about the folder architecture, if the process scales to a monolith with several developers, this architecture would not be so friendly. A transition to a folder architecture using interfaces and handlers would be important for scalability and better code maintainability.