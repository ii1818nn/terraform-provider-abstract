# abstract_datasource

Read arbitrary data by executing a shell command and capturing its output.

## Example Usage

```hcl
data "abstract_datasource" "example" {
  read_command = "aws secretsmanager get-secret-value --secret-id $SECRET_ID --query SecretString --output text"

  variables = {
    SECRET_ID = "my-secret"
  }

  secrets = {
    AWS_SECRET_ACCESS_KEY = var.aws_secret_key
  }
}

output "secret_value" {
  value     = data.abstract_datasource.example.output
  sensitive = true
}
```

## Argument Reference

- `read_command` - (Required, string) Shell command to execute. Stdout is captured as `output`.
- `variables` - (Optional, string map) Environment variables available during command execution. Default empty map.
- `secrets` - (Optional, string map) Same as `variables` but marked sensitive — values are redacted in logs and state output. In case of key collision, `secrets` takes precedence over `variables`. Default empty map.
- `triggers` - (Optional, string map) Arbitrary values that force the data source to re-read when changed. Default empty map.

## Attribute Reference

- `output` - (string) Stdout captured from the executed command.
- `exit_code` - (number) Exit code of the executed command.
