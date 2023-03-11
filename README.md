# Terraform Provider for Postman

- [Documentation](https://registry.terraform.io/providers/jonnydgreen/postman/latest/docs)
- Tutorials:
  [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)

The Terraform Postman provider is a plugin that allows
[Terraform](https://www.terraform.io) to manage resources in Postman.

**WARNING:** The Terraform Provider for Postman makes use of the Postman API.
Before proceeding, please ensure you have checked your Postman API usage plans
on your personal/team account
[resource usage page](https://web.postman.co/billing/add-ons/overview) within
Postman.

## Quick Starts

```terraform
terraform {
  required_providers {
    postman = {
      version = "0.2"
      source  = "jonnydgreen/postman"
    }
  }
}

provider "postman" {}

resource "postman_workspace" "example" {
  name = "Example"
  type = "personal"
}

resource "postman_environment" "example" {
  name      = "Example"
  workspace = postman_workspace.example.id
  values = [
    {
      key   = "hello"
      value = "there"
    },
    {
      key     = "foo"
      value   = "bar"
      enabled = false
      type    = "secret"
    },
  ]
}
```

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- [Go](https://golang.org/doc/install) >= 1.18

## Supported Features

| Feature      | Resource           | Data Source        | Import             |
| ------------ | ------------------ | ------------------ | ------------------ |
| Workspaces   | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Environments | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Apis         | :x:                | :x:                | :x:                |
| Collections  | :x:                | :x:                | :x:                |
| Mocks        | :x:                | :x:                | :x:                |
| Monitors     | :x:                | :x:                | :x:                |

## Using the provider

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

Contributions are very welcome! :)

See the [contributing guide](./CONTRIBUTING.md) for more details.

## Future work

- [ ] Support for APIs
- [ ] Support for Collections
- [ ] Support for Mocks
- [ ] Support for Monitors
- [ ] Support for input validators
- [ ] Support for automated acceptance testing

### Acceptance tests ideas

- Manual Dispatch
- On main only
- Or both of the above
