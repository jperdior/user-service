# User service in DDD made with Golang

![test workflow](https://github.com/jperdior/user-service/actions/workflows/test.yml/badge.svg)

This is a self learning project to implement a User service in Domain-Driven Design using Golang. It's also my intention to even be able to use it in future projects.

I am creating issues to guide the development of this project, so if you want to contribute, feel free to take a look at the issues or contribute with your own ideas.

## Requirements

- Go
- Docker
- Docker Compose
- Make
- Available ports 9091

## Running the project

To run the project, you can use the following command:

```bash
make start
```

**The first time you run this command it will ask you for a project name and that will update all the files with the new project name.**

## Other commands

To restart the project to see changes applied, you can use the following command:

```bash
make restart
```

To stop the project, you can use the following command:

```bash
make stop
```

To run the tests, you can use the following command:

```bash
make test
```

To run the analysis, you can use the following command:

```bash
make analysis
```
To view the API documentation you can run the following command:

```bash
make open-docs
```

## Developing new endpoints

To refresh the openapi documentation you can run the following command:

```bash
make refresh-openapi
```