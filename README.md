# Terraform Provider for Postman

- Tutorials:
  [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
- Documentation: TODO

The Terraform Postman provider is a plugin that allows
[Terraform](https://www.terraform.io) to manage resources in Postman.

## Quick Starts

TODO

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18

## Using the provider

TODO

Fill this in for each provider.

### Upgrading the provider

The Postman provider doesn't upgrade automatically once you've started using it.

After a new release you can run

```bash
terraform init -upgrade
```

to upgrade to the latest stable version of the Postman provider. See the
[Terraform website](https://www.terraform.io/docs/configuration/providers.html#provider-versions)
for more information on provider upgrades, and how to set version constraints on
your provider.

## Developing the Provider

See the [contributing guide](./CONTRIBUTING.md) for more details.

## Issues with Postman API definition

- Workspace query string for all requests is not `workspaceId` but `workspace`.
