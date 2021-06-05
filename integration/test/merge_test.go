package test

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeJsonMulti(t *testing.T) {
	terraformOptions := makeTerraformOptions(t, "../fixtures/merge_json_multi")
	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "json_multi")
	assert.Equal(t, "{\"hello\":\"galaxy\"}", output)
}

func TestMergeJsonSingle(t *testing.T) {
	terraformOptions := makeTerraformOptions(t, "../fixtures/merge_json_single")
	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "json_single")
	assert.Equal(t, "{\"hello\":\"world\"}", output)
}

func TestMergeYamlMulti(t *testing.T) {
	terraformOptions := makeTerraformOptions(t, "../fixtures/merge_yaml_multi")
	defer terraform.Destroy(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "yaml_multi")
	assert.Equal(t, "hello: galaxy\n", output)
}

func TestMergeYamlSingle(t *testing.T) {
	terraformOptions := makeTerraformOptions(t, "../fixtures/merge_yaml_single")
	defer terraform.Apply(t, terraformOptions)

	terraform.Apply(t, terraformOptions)
	output := terraform.Output(t, terraformOptions, "yaml_single")
	assert.Equal(t, "hello: world\n", output)
}
