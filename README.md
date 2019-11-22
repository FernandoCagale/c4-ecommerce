**c4-ecommerce - docker**

```sh
$   docker build -t c4-ecommerce .
```

```sh
$   docker run -d --name c4-ecommerce -p 8080:8080 c4-ecommerce
```

**c4-ecommerce - local**


```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download wire "dependency injection"`

```sh
$   go get -u github.com/google/wire/cmd/wire
```

`generate wire_gen.go`

```sh
$   wire
```

`generate build`

```sh
$   go build -o bin/application
```


```sh
$   ./bin/application
```