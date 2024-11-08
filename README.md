
# Go Dependency Injection Example

This project demonstrates a simple example of dependency injection in Go using the `go.uber.org/fx` package. It includes an HTTP server with a single endpoint that echoes the request body and prints a pyramid pattern to the console.

## Project Structure

- `main.go`: The entry point of the application. It sets up the dependency injection and starts the HTTP server.
- `pkg/services/pyramid_service.go`: Contains the `PyramidService` which provides a method to print a pyramid pattern.
- `pkg/handlers/echo_handler.go`: Contains the `EchoHandler` which handles HTTP requests to the `/echo` endpoint.

## Dependencies

- `go.uber.org/fx`: A dependency injection framework for Go.
- `go.uber.org/zap`: A fast, structured logging library for Go.

## Running the Project

1. Clone the repository:
    ```sh
    git clone https://github.com/ericomonteiro/go-dependency-injection
    cd go-dependency-injection
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

4. Send a request to the `/echo` endpoint:
    ```sh
    curl -X POST http://localhost:8080/echo -d "Hello, World!"
    ```

## Example Output

When you send a request to the `/echo` endpoint, the server will echo the request body and print a pyramid pattern to the console:

```sh
Handling request path=/echo
     *
    * *
   * * * 
  * * * * 
 * * * * * 
```

## License

This project is licensed under the MIT License.