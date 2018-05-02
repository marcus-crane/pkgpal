# pkgpal

pkgpal (name pending) is a CLI tool written in Go. It's designed to be the successor to my side project [pkgparse-node](https://github.com/marcus-crane/pkgparse-node) but where that was Node specific, this version is language agnostic.

# What is this?

The main functionality is that you can easily query packages from various package managers and also automatically parse, and better understand package manager files.

For example, with pkgparse-node, you could execute `pkgparse feast` in a directory with a `package.json` file and you would be returned a list of dependencies and their descriptions. It was useful for figuring out what dependenices a project has, and also usually what functionality you could expect from it.

This project strives to do the same but in a way that you could run eg; `pkgpal feast` in any directory and it'll automatically detect what files are present. Let's say it finds a requirements.txt file then it would parse that and query pypi by way of the [pkgparse server](https://github.com/marcus-crane/pkgparse).

# Development

At the moment, I'm just writing tests and executing them using VS Code. The project isn't in any usage state, nor does it have the CLI part set up. I'll be likely using [urfave/cli](https://github.com/urfave/cli) v2 but I'll do all that once the functionality is in place first.

You can install everything needed by cloning the project into your Go path and running `go get -t ./...`.

I'm pretty sure that works anyway. Let me know if it doesn't.

# Tests

 I'm trying to do this project using TDD which is interesting and hard because I'm still fairly new to go.

 You can run the current tests with `go test ./...`.
