# Contributing

## Developing

If you wish to work on the provider, you'll first need
[Go](http://www.golang.org) installed on your machine (see
[Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put
the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider:

```sh
$ make install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using
Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform
provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Updating the API definition

Copy the updated Postman API definition into `openapi.yaml`. Then run:

```bash
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
  -i /local/openapi.yaml \
  -g go \
  -o /local/pkg/postman-sdk
```

## Local testing

```bash
(cd examples/testing && \
  rm -rf .terraform .terraform.lock.hcl && \
  terraform init && \
  terraform apply --auto-approve)
```
