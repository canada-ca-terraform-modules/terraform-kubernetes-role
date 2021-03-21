package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAdminUserRole(t *testing.T) {
	t.Parallel()

	expectedRoleName := "admin-role"
	expectedRoleBindingName := expectedRoleName + "-binding"

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/admin-user-role",
	}

	// Clean up after the test
	defer terraform.Destroy(t, terraformOptions)

	// Create the Role and RoleBinding
	terraform.InitAndApply(t, terraformOptions)

	// Get the output
	roleName := terraform.Output(t, terraformOptions, "role_name")
	roleBindingName := terraform.Output(t, terraformOptions, "role_binding_name")

	// Check the results
	assert.Equal(t, expectedRoleName, roleName)
	assert.Equal(t, expectedRoleBindingName, roleBindingName)

}
