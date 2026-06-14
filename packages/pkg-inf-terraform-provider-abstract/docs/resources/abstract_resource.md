# abstract_resource

Manage an arbitrary resource by specifying shell commands for each Terraform lifecycle phase.

## Example Usage

```hcl
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

  secrets = {
    AWS_SECRET_ACCESS_KEY = var.aws_secret_key
  }

  triggers = {
    bucket_name = "my-bucket"
  }
}
```

## Argument Reference

- `lifecycle_commands` - (Required) see [lifecycle_commands](#lifecycle_commands).
- `variables` - (Optional, string map) Environment variables available in each lifecycle command. Default empty map.
- `secrets` - (Optional, string map) Same as `variables` but marked sensitive — values are redacted in logs and state output. In case of key collision, `secrets` takes precedence over `variables`. Default empty map.
- `triggers` - (Optional, string map) Arbitrary values that force resource recreation when changed, similar to [`null_resource.triggers`](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource#triggers). Default empty map.

### lifecycle_commands

Block containing shell commands executed during each Terraform lifecycle phase.

- `create` - (Required, string) Command executed in the **Create** phase.
- `read` - (Optional, string) Command executed in the **Read** phase and after `create`/`update`. A non-zero exit code signals that the resource no longer exists and triggers recreation. Defaults to `exit 0`.
- `update` - (Optional, string) Command executed in the **Update** phase. Defaults to `exit 0` (no-op — only in-place update, no recreation).
- `delete` - (Required, string) Command executed in the **Delete** phase.

## Attribute Reference

- `output` - (string) Stdout captured from the last executed command.
- `exit_code` - (number) Exit code of the last executed command.
