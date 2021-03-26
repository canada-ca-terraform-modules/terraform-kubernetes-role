package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

// Test using test_structure, which allows sections of the tests to be skipped
// if a `SKIP_stage_name=true` environment variable is set.  This allows for
// selectively skipping long running test steps like terraform apply/destroy.
func TestNoCreateRole(t *testing.T) {
	t.Parallel()

	workingDir := "../examples/no-create-role"

	terraformOptions := &terraform.Options{
		TerraformDir: workingDir,
	}
	test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

	// Destory the role
	defer test_structure.RunTestStage(t, "destroy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.Destroy(t, terraformOptions)
	})

	// Apply and check that nothing was created
	test_structure.RunTestStage(t, "apply", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		counts := terraform.GetResourceCount(t, terraform.InitAndApply(t, terraformOptions))
		assert.Equal(t, 0, counts.Add)
		assert.Equal(t, 0, counts.Change)
		assert.Equal(t, 0, counts.Destroy)
	})
}
