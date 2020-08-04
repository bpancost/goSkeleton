# Golang Skeleton

A skeleton project in Golang with some boilerplate already implemented.
This is intended to make it easier to start a new project without needing a particular framework.
It is still an opinionated starting point, just less rigid in some regards.
Examples of a CLI project and a RESTful server are included and are intended to be adapted for new projects.

## General Philosophy

The way that this project is laid out is heavily inspired by the
[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and
[Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)). I also used
[Clean Architecture in Go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1) as a starting point.
The core concept is that the use cases or business logic should only concern themselves with the actual logic.
Other details such as how the business logic is invoked or how to use a database should be separate implementation
details. The use cases own and define the interfaces for how it interacts with other parts of the software, but the
implementation (or which implementation) should be transparent. This allows us to split up the application to make it
easier to develop and to unit test at each level.

## Directory Structure

### `/adapters`
Implementations for the triggers that activate the use cases, such as REST based web calls, GRPC calls, or asynchronous
messaging. This directory also includes implementations for data repositories (databases or other services) that satisfy
the use case interfaces for data storage and retrieval.
### `/app`
The scaffolding for the application which instantiates the concrete types. This can be considered the "dirty" part of
the application which actually needs to know which type is really behind the interfaces and plug them together so
everything runs. This directory should not include *any* business logic.
### `/cmd`
The entry points for the executables that need to be generated for the project. Each executable should have its own
subdirectory. In general these will call the scaffolding within the `/app` directory and should be kept relatively
light. Config files for development might also be in these directories.
### `/docs`
General project documentation, diagrams, and other information.
### `/domain`
Structs which describe sets of data, such as records within a database or returned from another service. These can
include convenience functions to make it easier to handle or transform the data it contains.
### `/internal`
Any common services which should not be available or used outside of this service. Depending on how these are used, they
may be candidates to move out to a separate repo and reference as a library to avoid copying into every service.
### `/proto`
Any protocol files which will be used for gRPC communication with this service.
### `/usecases`
The business logic of the service. This includes an interface covering each use case (for triggering adapters) as well
as the interfaces forming a contract for any repository or remote service implementation. There should be no direct
calls to any other service or data storage within the use cases.

## Usage

### Bazel

Bazel has been included to make it easier to incrementally build and test the services, as well as to reliably make
docker container images. This is best for quickly checking that a build compiles, passes tests, and can run in a
container.

All the scripts below should be run from the root project directory.

#### Building

To build everything together, use
```shell script
bazel build //...
```

Only the changed libraries and executable get built, so it will almost always make sense to build everything in the
project.

#### Testing

To run all the unit tests, use
```shell script
bazel test //...
```

To run a particular package's tests, you can indicate the directory and build target (generally `go_default_test`) such
as below
```shell script
bazel test //adapters/repository/person:go_default_test
```

#### Running

If a `BUILD.bazel` file has a container_image definition, then that can be run as a bazel run target. An 
example of this for the `/cmd/people_server` executable is below.
```shell script
bazel run //cmd/people_server:image_dev
```

This creates the docker image and tags it as `bazel/{directory}:{target}`. Using the same example as before,
it can be run in docker like below.
```shell script
docker run --rm -it -p8080:8080 bazel/cmd/people_server:image_dev
```

There are currently two targets in the project, both for the `people_server` example project. The first is `image_dev`
which copies in the `config.yml` and all of its development configuration. The second is `image` which relies only on
environment variable definitions. In a production use case, the environment variables would not be included here but
injected when deployed by your preferred continuous delivery system.

The container_image definitions are not generated by Gazelle and have to be manually created.

##### gRPC Testing

It is highly recommended to use [Evans](https://github.com/ktr0731/evans) to test the gRPC portion of this service.
To run evans for the current protocol file, use this command
```shell script
evans ./proto/people.proto
```

#### Updating the Build Files

Bazel describes the project by explicitly exhaustively describing the files in each package as a library and the
dependencies between each package/library. If a new file is added or dependencies between the packages change, the build
files will need to be updated. While this can be done manually, it is easiest to use Gazelle to automatically update the
build files.

I would suggest having Gazelle installed with Go if you haven't already, though this is optional since you can run
Gazelle with Bazel as well. If running with Go, then in the following always chose the commands starting with `gazelle`.
```shell script
go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle
```

If the external dependencies have been updated in go.mod, then run one of the following depending on whether gazelle
is installed locally or not
```shell script
bazel run //:gazelle -- update-repos -from_file=go.mod
```
```shell script
gazelle update-repos -from_file=go.mod
```

To update the build files, run one of the following
```shell script
gazelle update
```
```shell script
bazel run //:gazelle
```

### Jetbrains GoLand

If you are using [GoLand](https://www.jetbrains.com/go/), project files have been included and can be imported. This
includes run configurations and an HTTP scratch pad for testing the example People service. The run configurations will
run the application directly and are best used for debugging the service.

If the protocol definitions have been updated, make sure to generate a new copy of the Go files. While the Bazel build
automatically generates an injects these files, this is not the case whenever running directly.

## Libraries and Tools Used

### Go Libraries
* [GoMock](https://github.com/golang/mock)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Golang Protobuf](https://github.com/golang/protobuf)
* [Google Golang Protobuf](https://godoc.org/google.golang.org/protobuf)
* [Google gRPC](https://godoc.org/google.golang.org/grpc)
* [Go gRPC Middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)
* [Package Errors](https://github.com/pkg/errors)
* [Go UUID](https://github.com/satori/go.uuid)
* [Logrus](https://github.com/sirupsen/logrus)
* [GoConvey](https://github.com/smartystreets/goconvey)
* [Viper](https://github.com/spf13/viper)
* [CLI](https://github.com/urfave/cli)

### Bazel Build
* [Bazel](https://bazel.build/)
* [Bazel Go Rules](https://github.com/bazelbuild/rules_go)
* [Bazel Docker Image Rules](https://github.com/bazelbuild/rules_docker)
* [Gazelle Bazel Build File Generator](https://github.com/bazelbuild/bazel-gazelle)