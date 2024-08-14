# ASCII-ART-WEB-DOCKERIZE
* This project demonstrates how to create a simple web server using Go, containerize it with Docker, and manage Docker images and containers. It follows Docker best practices and applies metadata to Docker objects while also managing unused objects efficiently.

## Features
- **Containerization**: Docker
- **Components**:
  - **Dockerfile**: Defines how to build the Docker image.
  - **Docker Image**: A snapshot of the application environment.
  - **Docker Container**: A running instance of the Docker image.

## Instructions to run locally

To clone this repository, copy the command below on your terminal:

```bash
git clone https://learn.zone01kisumu.ke/git/hanapiko/ascii-art-web-dockerize.git
```

Go to the project directory
```bash
cd ascii-art-web-dockerize
```

## Usage

- To build the image, use the command below;

```bash
docker build -t <imageName> .
```
- To create and run the container use the command below;
```bash
docker run -d -p hostport:containerport --name <containerName> <imageName>
```

## AUTHORS
- [hanapiko](https://learn.zone01kisumu.ke/git/hanapiko)

- [weakinyi](https://learn.zone01kisumu.ke/git/weakinyi)

- [somotto](https://learn.zone01kisumu.ke/git/somotto)

## LICENSE
This project is licensed under the terms of the [MIT License](./LICENSE).

