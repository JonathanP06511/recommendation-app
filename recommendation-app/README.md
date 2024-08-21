# Information API

This project is a Recommendation API built with Go. It provides an endpoint to retrieve a random health recommendation.

## Technologies Used

- **Go**: Programming language used for building the server-side application.
- **Net/HTTP**: Standard library package in Go for creating HTTP servers and handling requests.
- **JSON**: Format for exchanging data between the server and clients.
- **Rand**: Package used for generating random numbers.

## Installation

1. Clone this repository to your local machine:
    ```bash
    git clone <repository URL>
    ```

2. Navigate to the project directory:
    ```bash
    cd recommendation-app
    ```

3. Build the project:
    ```bash
    go build
    ```

## Usage

1. Start the server:
    ```bash
    go run main.go
    ```

2. The API will be available at `http://localhost:8080`.

## Project Structure

- `main.go`: Main file where the HTTP server is configured and the routes are set up.
- `routes/routes.go`: Defines the route for retrieving recommendations and contains the handler logic.
- `models/recommendation.go`: Defines the `Recommendation` model used for the API responses.

## Routes

- `GET /recommendations`: Endpoint to retrieve a random health recommendation.

## Docker

To run this project in a Docker container:

1. Build the Docker image:
    ```bash
    docker build -t recommendation-app .
    ```

2. Run the Docker container:
    ```bash
    docker run -p 8080:8080 recommendation-app
    ```

3. The API will be available at `http://localhost:8080` inside the Docker container.

## Documentation with Swagger

API documentation is generated with Swagger and served at the `/api-docs` endpoint. Swagger allows you to view and test the API endpoints interactively.

