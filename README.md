> ⚠️ **Incoming holidays**: The whole team will be absent from the 23rd of December to the 7th of January. Our involvement in GitHub issues during this time will be limited. We will look out for the critical issues, though. Merry Christmas and a Happy New Year!

## Support

For official support and urgent, production-impacting issues, please [contact Snowflake Support](https://community.snowflake.com/s/article/How-To-Submit-a-Support-Case-in-Snowflake-Lodge).

> ⚠️ **Keep in mind** that the official support starts with the [v2.0.0](https://registry.terraform.io/providers/snowflakedb/snowflake/2.0.0) for stable resources only. All previous versions and preview resources are not officially supported.

Please follow [creating issues guidelines](CREATING_ISSUES.md), [FAQ](FAQ.md), and [known issues](KNOWN_ISSUES.md) before submitting an issue on GitHub or directly to Snowflake Support.

# Snowflakier Terraform Provider

This is a fork of the official [Snowflake Terraform Provider](https://github.com/Snowflake-Labs/terraform-provider-snowflake) with additional functionality while maintaining compatibility with the base provider.

> ⚠️ **Note**: This is an unofficial fork maintained by [@henryupton](https://github.com/henryupton). For official support, use the [official Snowflake provider](https://registry.terraform.io/providers/Snowflake-Labs/snowflake).

This is a terraform provider for managing [Snowflake](https://www.snowflake.com/) resources with extended capabilities.

## Table of contents
<!-- TOC -->
* [Snowflake Terraform Provider](#snowflake-terraform-provider)
  * [Table of contents](#table-of-contents)
  * [Getting started](#getting-started)
  * [Supported Architectures](#supported-architectures)
  * [Migration guide](#migration-guide)
  * [Roadmap](#roadmap)
  * [Getting Help](#getting-help)
  * [Would you like to create an issue?](#would-you-like-to-create-an-issue)
  * [Additional SQL Client configuration](#additional-sql-client-configuration)
  * [Contributing](#contributing)
  * [Releases](#releases)
<!-- TOC -->

## Getting started

Install the Snowflakier Terraform provider by adding a requirement block and a provider block to your Terraform codebase:

```hcl
terraform {
  required_providers {
    snowflakier = {
      source  = "henryupton/snowflakier"
      version = ">= 0.1.0"
    }
    # You can use both providers side-by-side
    snowflake = {
      source  = "Snowflake-Labs/snowflake"
      version = ">= 1.0.0"
    }
  }
}

provider "snowflakier" {
  organization_name = "organization_name"
  account_name      = "account_name"
  user              = "johndoe"
  password          = var.password
  role              = "PUBLIC"
}

# Remember to provide the password securely.
variable "password" {
  type      = string
  sensitive = true
}
```

This fork maintains resource compatibility with the base Snowflake provider, so you can use them together.

Don't forget to run `terraform init` and you're ready to go! 🚀

Start browsing the [registry docs](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs) to find resources and data sources to use.

## Supported Architectures

We have compiled a list to clarify which binaries are officially supported and which are provided additionally but not officially supported.
The lists are based on what the underlying [gosnowflake driver](https://github.com/snowflakedb/gosnowflake) supports and what [HashiCorp recommends for Terraform providers](https://developer.hashicorp.com/terraform/registry/providers/os-arch).

The Snowflake Terraform provider supports the following architectures:
- Windows: amd64
- Linux: amd64 and arm64
- Darwin: amd64 and arm64

We also provide additional configurations, but they are not officially supported, and we do not prioritize fixes for them:
- Windows: arm64 and 386
- Linux: 386
- Darwin: 386
- Freebsd: any architecture

## Migration guide

Please check the [migration guide](./MIGRATION_GUIDE.md) when changing the version of the provider.

## Roadmap

Check [Roadmap](./ROADMAP.md).

## Getting Help

Some links that might help you:

- The [introductory tutorial](https://guides.snowflake.com/guide/terraforming_snowflake/#0) shows how to set up your Snowflake account for Terraform (service user, role, authentication, etc) and how to create your first resources in Terraform.
- The [docs on the Terraform registry](https://registry.terraform.io/providers/snowflakedb/snowflake/latest) are a complete reference of all resources and data sources supported and contain more advanced examples.
- The [discussions area](https://github.com/snowflakedb/terraform-provider-snowflake/discussions) of this repo, we use this forum to discuss new features and changes to the provider.
- **If you are an enterprise customer**, reach out to your account team. This helps us prioritize issues.
- The [issues section](https://github.com/snowflakedb/terraform-provider-snowflake/issues) might already have an issue addressing your question.

## Would you like to create an issue?
If you would like to create a GitHub issue, please read our [guide](./CREATING_ISSUES.md) first.
It contains useful links, FAQ, and commonly known issues with solutions that may already solve your case.

## Additional SQL Client configuration
The provider uses the underlying [gosnowflake](https://github.com/snowflakedb/gosnowflake) driver to send SQL commands to Snowflake.

By default, the underlying driver is set to error level logging. It can be changed by setting `driver_tracing` field in the configuration to one of (from most to least verbose):
- `trace`
- `debug`
- `info`
- `print`
- `warning`
- `error`
- `fatal`
- `panic`

Read more in [provider configuration docs](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs#schema).

## Contributing

Check [Contributing](./CONTRIBUTING.md).

## Releases

Releases will be performed as needed, typically once every 2 weeks.

Releases are published to [the terraform registry](https://registry.terraform.io/providers/snowflakedb/snowflake/latest). Each change has its own release notes (e.g. https://github.com/snowflakedb/terraform-provider-snowflake/releases/tag/v0.89.0) and migration guide if needed (e.g. https://github.com/snowflakedb/terraform-provider-snowflake/blob/main/MIGRATION_GUIDE.md#v0880--v0890).
