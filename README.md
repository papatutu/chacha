# Single-page application server

This is a simple Go Fiber application that serves a single-page application (React, Vue, Flutter, Angular)  with support for custom routes. The application uses Viper for configuration management and supports environment variables, configuration files, and command-line flags.

## Features

- Serves static files from a specified directory
- Catches all routes and serves `index.html` for SPA routing
- Configurable via environment variables, configuration files, and command-line flags

## Example Command Line Usage

To run the application with custom settings using command-line flags, you can use the following example:

```sh
chacha --port=8080 --serveDir=./dist
```

This command will start the server on port `8080` and serve static files from the `./dist` directory.

## Configuration

The application can be configured using environment variables, or command-line flags.

### Environment Variables

- `PORT`: The port to run the server on (default: `9000`)
- `SERVE_DIR`: The directory to serve static files from (default: `./public`)

### Configuration File

Create a `config.env` file in the root directory with the following content:

```
PORT=9000 SERVE_DIR=./public
```

### Command-Line Flags

- `--port`: The port to run the server on
- `--serveDir`: The directory to serve static files from

### Configuration Priority

The priority of configuration settings is as follows:

1. Command-line flags
2. Configuration file
3. Environment variables

Command-line flags will overwrite settings from the configuration file, and settings from the configuration file will overwrite environment variables if both are set.

## Docker

You can easily run this application using Docker. The Docker image is available on Docker Hub under the repository `papatutu/chacha`. To pull the image and run the container, use the following command:

```sh
docker run -p 9000:9000 -e SERVE_DIR=./build/web berberin/chacha:latest
```


## Docker Compose

To run this application using Docker Compose, create a `docker-compose.yml` file in the root directory with the following content:

```yaml
services:
    chacha:
        image: papatutu/chacha:latest
        ports:
            - "9000:9000"
        environment:
            - PORT=9000
            - SERVE_DIR=./build/web
        volumes:
            - ./build:/app/build
```

Then, start the application with:

```sh
docker-compose up
```

This command will start the server on port `9000` and serve static files from the `./build/web` directory.


# License
This project is licensed under the MIT License.