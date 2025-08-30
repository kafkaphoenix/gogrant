# gogrant

## Description
This repository contains the implementation of a service that demonstrates how to deploy a Golang API with Casbin, Keycloak and MongoDB using Docker Compose.

## Requirements
To run this project, please follow the steps listed below:

### 1. Go installation
Follow the official [Go installation guide](https://go.dev/doc/install) to install Go 1.24.

### 2. Install Docker & Docker Compose
Ensure that Docker and Docker Compose are installed:
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Docker Compose Installation](https://docs.docker.com/compose/install/)
> Docker has to be running to start the application.

## Makefile
- lint: Runs golangci-lint to check and fix code style issues.
- tests: Runs the tests.
- swagger: Generates Swagger documentation for the API.
- mocks: Generates mocks for interfaces using `mockery`.
- up: Starts the server using Docker Compose.
- logs: Displays the logs of the server.
- clean: Cleans up Docker resources.
- deploy: Command for development purposes, mixing previous commands.

## Try it out!

### 1. Start service:
```sh
make up
```

### 2. Clean any docker resource:
```sh
make clean
```

### 3. View application logs:
```sh
make logs
```

### 4. Test:
After starting the service, you have two options:
- Open your browser and navigate to `http://localhost:8080/api/v0/swagger/index.html`.
- Use a tool like [Bruno](https://www.usebruno.com/) to interact with the API. You can find the collection in the `e2e` folder.

## API Endpoints
### 0. Health Check

```bash
curl -X 'GET' \
  'http://localhost:8080/api/v0/health' \
  -H 'accept: application/json'
```

### 1. Create Document

```bash
curl -X 'POST' \
  'http://localhost:8080/api/v0/documents' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "author": "Author Name",
  "content": "This is the content of the document.",
  "description": "This is a sample document description.",
  "format": "markdown",
  "tags": [
    "tag1",
    "tag2"
  ],
  "title": "Document Title"
}'
```

### 2. Get Document by ID

```bash
curl -X 'GET' \
  'http://localhost:8080/api/v0/documents/{id}' \
  -H 'accept: application/json'
```
> Replace `{id}` with the actual document ID.

### 3. List documents given filters

```bash
curl -X 'GET' \
  'http://localhost:8080/api/v0/documents?author=Lovecraft&tags=tag1%2Ctag2&format=markdown' \
  -H 'accept: application/json'
```
> You can use the query parameters `author`, `tags`, and `format` to filter the documents.
