# Baz Service

This is the Baz service

Generated with

```
micro new --namespace=echo.micro --type=service baz
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: echo.micro.service.baz
- Type: service
- Alias: baz

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./baz-service
```

Build a docker image
```
make docker
```