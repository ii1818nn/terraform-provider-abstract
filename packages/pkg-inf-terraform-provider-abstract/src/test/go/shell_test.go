package runner_test

import (
	"testing"

	"github.com/ii1818nn/pkg-inf-terraform-provider-abstract/src/main/go/runner"
)

func TestRun_Output(t *testing.T) {
	res, err := runner.Shell("echo hello", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Output != "hello" {
		t.Errorf("expected output %q, got %q", "hello", res.Output)
	}
	if res.ExitCode != 0 {
		t.Errorf("expected exit code 0, got %d", res.ExitCode)
	}
}

func TestRun_ExitCode(t *testing.T) {
	res, err := runner.Shell("exit 42", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.ExitCode != 42 {
		t.Errorf("expected exit code 42, got %d", res.ExitCode)
	}
}

func TestRun_EnvVariables(t *testing.T) {
	res, err := runner.Shell("echo $MY_VAR", map[string]string{"MY_VAR": "hello_env"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Output != "hello_env" {
		t.Errorf("expected output %q, got %q", "hello_env", res.Output)
	}
}

func TestRun_EnvSecretOverridesVariable(t *testing.T) {
	res, err := runner.Shell("echo $KEY", map[string]string{"KEY": "secret_value"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Output != "secret_value" {
		t.Errorf("expected output %q, got %q", "secret_value", res.Output)
	}
}

func TestRun_InvalidCommand(t *testing.T) {
	_, err := runner.Shell("", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
