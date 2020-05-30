# c4-ecommerce

## Flow [c4-kustomize](https://github.com/FernandoCagale/c4-kustomize)

## Dependencies

`Docker Rabbitmq`

```sh
$ docker run --network host --name rabbit -d rabbitmq
```

## Build Docker

```sh
$   docker build -t c4-ecommerce .
```

```sh
$   docker run -d --name c4-ecommerce -p 8080:8080 c4-ecommerce
```

## Kubernetes [YAML](https://github.com/FernandoCagale/c4-kustomize/tree/master/c4-ecommerce/base)

    *   deployment.yaml
    *   service.yaml
    *   virtualservice.yaml

## Running local

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

```sh
$   ./scripts/start.sh
```