package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
)

func makeTerraformOptions(t *testing.T, terraformDir string) *terraform.Options {
	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		EnvVars: map[string]string{
			"TF_CLI_CONFIG_FILE": "../.terraformrc",
			"TF_LOG": "TRACE",
		},
	})
}
