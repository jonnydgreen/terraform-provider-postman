# Contributing

## Developing

### Setup

If you wish to work on the provider, you'll first need
[Go](http://www.golang.org) installed on your machine (see
[Requirements](#requirements) above).

Then,
[follow the instructions](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install)
to set up your local dev environment. For example:

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/jonnydgreen/postman" = "<path_to_go_bin>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

### Compile provider

To compile the provider, run:

```sh
go install .
```

This will build the provider and put the provider binary in the `$GOPATH/bin`
directory.

### Generate documentation

To generate or update documentation, run:

```sh
make gen
```

### Run acceptance tests

**WARNING:** The Terraform Provider for Postman makes use of the Postman API.
Before proceeding, please ensure you have checked your Postman API usage plans
on your personal/team account
[resource usage page](https://web.postman.co/billing/add-ons/overview) within
Postman.

Because of limits in Postman API usage, it is **not recommended** to run the full
suite of Acceptance tests. Instead, it is better to run them on a per-test basis
when needed by providing the relevant test name.

Before running, ensure the `POSTMAN_API_KEY` environment variable is set.
Otherwise, the tests will not run. This can be obtained in the Postman UI. Once
set, run the tests as follows:

```sh
make testacc TESTARGS='-run=TestAccWorkspaceResource__basic'
```

### Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using
Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform
provider:

```sh
go get github.com/author/dependency
make vendor
```

Then commit the changes to `go.mod` and `go.sum`.

### Updating the API definition

Copy the updated Postman API definition into `openapi.yaml`. Then run:

```bash
make client
```

### Local dev

Setup git commit hooks by running the following command at the root of the repo:

```sh
git config core.hooksPath .githooks
```

## Submitting a pull request

When submitting a pull request, please provide the following:

- Description of the changes in the PR
- Documentation
- Acceptance test (including shell output from running it)
- Any other useful information

## Notes

### Issues with Postman API definition

- Workspace query string for all requests is not `workspaceId` but `workspace`.
- Environment values has double nested array in both request and responses
