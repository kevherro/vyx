[![Github Action CI](https://github.com/kevherro/vyx/workflows/ci/badge.svg)](https://github.com/kevherro/vyx/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/kevherro/vyx.svg)](https://pkg.go.dev/github.com/kevherro/vyx)

# Introduction

# Building vyx

Prerequisites:

- Go development kit of a [supported version](https://golang.org/doc/devel/release.html#policy).
  Follow [these instructions](http://golang.org/doc/code.html) to prepare
  the environment.

- OpenAI API key. Set it as an environment variable named `OPENAI_API_KEY`.

To build and install it:

    go install github.com/kevherro/vyx@latest

The binary will be installed in `$GOPATH/bin` (`$HOME/go/bin` by default).

# Basic Usage

vyx runs in interactive mode. It accepts interactive discourse:

```
% ./vyx # Start vyx
```

This will open a simple shell. Type 'help' for available commands/options.

```
% (vyx) who am i?
```
