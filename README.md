# go-snake
Implementing a simple snake game using the GUI module I made here: https://github.com/gowhale/LED-MATRIX-GOLANG

## Raspberry Pi Demo 

If an LED segment display has been wired up to the Raspberry Pi correctly. After running the following command: 
`go run .`


## Terminal Demo

To run the code without using an Raspberry Pi and to get the output to print onto the terminal use the following command:

`go run . -debug` The debug flag specifies to just print on the terminal. 


## Actions used in this repo:

### Testing

The pkg-cov workflow runs all go tests and ensures pkg coverage is above 80%.

![example event parameter](https://github.com/gowhale/go-snake/actions/workflows/pkg-cov.yml/badge.svg?event=push)

The pages workflow publishes a test coverage website everytime there is a push to the main branch. The website can be found here: https://gowhale.github.io/go-snake/#file0

![example event parameter](https://github.com/gowhale/go-snake/actions/workflows/pages.yml/badge.svg?event=push)

### Linters

The revive workflow is executed to statically analsye go files: https://github.com/mgechev/revive

![example event parameter](https://github.com/gowhale/go-snake/actions/workflows/revive.yml/badge.svg?event=push)

The golangci-lint workflow runs the golangci-lint linter: https://github.com/golangci/golangci-lint

![example event parameter](https://github.com/gowhale/go-snake/actions/workflows/golangci-lint.yml/badge.svg?event=push)
