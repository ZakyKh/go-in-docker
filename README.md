# Go in Docker

This repository is implementation of a very simple Go microservice with Docker-related files that allows it to be run in Docker.

### Installation

1. Clone the files in this repository
    `(follow repo instruction above)`
2. Configure environment variable: Edit `.env` as needed
    ```bash
	cp env.sample .env
	```
3. Install dependencies
    ```bash
	go mod download && go mod verify
	```
4. Run the microservice
    ```bash
	go run main.go
	```
5. Test if the microservice runs smoothly. Replace `APP_PORT` with configured port in `.env`
    ```bash
	curl localhost:<APP_PORT>/healthz
	```
	
### Running in Docker
Reference: [Official Docker Hub Page for Golang Image](https://hub.docker.com/_/golang)

1. Build the image
    ```bash
	docker build -t go-docker .
	```
2. Run the built image. Replace `EXPOSED_PORT` with port to be opened to host machine, and `PORT_IN_ENV` with port in env.sample
    ```bash
	docker run -it --rm --name go-in-docker -p <EXPOSED_PORT>:<PORT_IN_ENV> go-docker
	```

