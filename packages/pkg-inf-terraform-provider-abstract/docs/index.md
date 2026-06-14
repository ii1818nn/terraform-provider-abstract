# Abstract Provider

Provider for managing arbitrary resources and reading arbitrary data by executing local shell commands for each Terraform lifecycle phase.

## Example Usage

```hcl
provider "abstract" {}

resource "abstract_resource" "example" {
  lifecycle_commands {
    create = "aws s3 mb s3://$BUCKET_NAME"
    read   = "aws s3api head-bucket --bucket $BUCKET_NAME && echo exists"
    update = "echo no-op"
    delete = "aws s3 rb s3://$BUCKET_NAME --force"
  }

  variables = {
    BUCKET_NAME = "my-bucket"
  }

  triggers = {
    bucket_name = "my-bucket"
  }
}

data "abstract_datasource" "example" {
  read_command = "aws secretsmanager get-secret-value --secret-id $SECRET_ID --query SecretString --output text"

  variables = {
    SECRET_ID = "my-secret"
  }
}
```

## Resources

- [abstract_resource](resources/abstract_resource.md) — manage an arbitrary resource via shell lifecycle commands.

## Data Sources

- [abstract_datasource](data-sources/abstract_datasource.md) — read arbitrary data by executing a shell command.
