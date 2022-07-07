# go-cleanarchitecture-app

Example of a Go application using a clean architecture.

### Install
Check out the source

```shell
cd $WORKDIR
git clone https://github.com/mormorbump/GoCleanArchitectureApp.git
```

### Setup the package

```shell
cd GoCleanArchitectureApp/src/app
go mod tidy
go mod vendor
```

### Running with Docker

```shell
cd $WORKDIR/GoCleanArchitectureApp
docker-compose up
```

Access the web endpoint at http://localhost:8080/users