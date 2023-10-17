# stock-exchange-sim

### Building

To build both binaries into `bin/` directory, use:

`go build -o bin/ ./...`

Or with [just](https://github.com/casey/just):

`just b` / `just build`

### Running

#### Simulator

`go run ./cmd/simulator <file> <timeout>`

Or with just:

`just s <file> <timeout>`

#### Checker

`go run ./cmd/checker <file> <log_file>`

Or just just:

`just c <file> <log_file>`

#### Get a png of the project

`go run cmd/simulator/main.go examples/complex && dot -Tpng output.dot -o build.png`
