set positional-arguments

alias b := build
alias s := simulator
alias c := checker

@build:
  go build -o bin/ ./...

@checker *args:
  go run ./cmd/checker $@

@simulator *args:
  go run ./cmd/simulator $@
