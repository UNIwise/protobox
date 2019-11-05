<div align="center">
<img src="resources/logo.png" height="300">
</div>

# protobox

protobox simplifies gRPC source generation and dependency management for your projects.

Keep track of proto files needed from folders or external git repos. Don't have protoc and all the plugins needed to generate your source files? No problem! protobox will use a docker image as a build host so you don't have to worry about the protoc build chain.

## Installation

```bash
go get github.com/UNIwise/protobox
```

## Quick start

```bash
# Display the help page
protobox help

# Initialize protobox.yaml
protobox init

# Add a proto dependency, target language 'go' and out directory 'gen'
protobox add https://github.com/UNIwise/protobox/blob/master/examples/service.proto go gen

# Add a proto local dependency, target language 'go' and out directory 'gen'
protobox add examples/service.proto go gen

# Generate your source files specified in protobox.yaml
protobox generate

# Use your local protoc binaries instead of the docker build host
protobox generate --local

# Check if your system have the needed dependencies for protobox to work
protobox check

# Lint the configuration file
protobox lint
```

## Configuration

All protobox project dependencies and settings are specified in a `protobox.yaml` in the root of your project.

A simple configuration with a local proto file:

```yaml
syntax: v1

services:
  - proto: service.proto
    out: 
      - language: go
        path: gen/
```

Proto files from git:

```yaml
syntax: v1

services:
  - repo: git@github.com:UNIwise/protobox
    ref: master # optional, a branch or commit hash
    proto: examples/service.proto
    out: 
      - language: go
        path: gen/
```

Multiple target languages:

```yaml
syntax: v1

services:
  - proto: service.proto
    out: 
      - language: go
        path: gen/go
      - language: php
        path: gen/php
```

Multiple services:

```yaml
syntax: v1

services:
  - proto: service1.proto
    out: 
      - language: go
        path: gen/service1
  - proto: service2.proto
    out: 
      - language: go
        path: gen/service2

```

And of course all examples above can be combined as you please! 

## Syntax



Available languages are: `go`, `php`, `ts`, `js`, `cpp`, `python`, `java`, `ruby`.