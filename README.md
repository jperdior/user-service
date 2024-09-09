# Boilerplate for a DDD module in Golang

![test workflow](https://github.com/jperdior/golang-template/actions/workflows/test.yml/badge.svg)

This is a boilerplate for a Domain-Driven Design module in Golang using gin-gonic.

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