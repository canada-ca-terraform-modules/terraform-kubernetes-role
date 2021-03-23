package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

// Test using test_structure, which allows sections of the tests to be skipped
// if a `SKIP_stage_name=true` environment variable is set.  This allows for
// selectively skipping long running test steps like terraform apply/destroy.
func TestAdminUserRole(t *testing.T) {
	t.Parallel()

	workingDir := "../examples/admin-user-role"

	expectedNamespace := "default"
	expectedRoleName := "admin-role"
	expectedRoleBindingName := expectedRoleName + "-binding"

	terraformOptions := &terraform.Options{
		TerraformDir: workingDir,
	}
	test_structure.SaveTerraformOptions(t, workingDir, terraformOptions)

	// Destory the role
	defer test_structure.RunTestStage(t, "destroy", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.Destroy(t, terraformOptions)
	})

	// Apply and check the output
	test_structure.RunTestStage(t, "apply", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, workingDir)
		terraform.InitAndApply(t, terraformOptions)

		roleName := terraform.OutputRequired(t, terraformOptions, "role_name")
		roleBindingName := terraform.OutputRequired(t, terraformOptions, "role_binding_name")

		assert.Equal(t, expectedRoleName, roleName)
		assert.Equal(t, expectedRoleBindingName, roleBindingName)
	})

	// Check the k8s cluster Role and RoleBinding objects
	test_structure.RunTestStage(t, "k8s_objects", func() {
		kubectlOptions := k8s.NewKubectlOptions("", "", expectedNamespace)

		// Check the Role
		role := k8s.GetRole(t, kubectlOptions, expectedRoleName)

		// Check that there is one rule
		assert.Equal(t, 1, len(role.Rules))
		assert.Equal(t, "*", role.Rules[0].APIGroups[0])
		assert.Equal(t, "*", role.Rules[0].Resources[0])
		assert.Equal(t, "*", role.Rules[0].Verbs[0])

		// Check the RoleBinding
		roleBinding := getRoleBinding(t, kubectlOptions, expectedNamespace, expectedRoleBindingName)

		assert.Equal(t, "Role", roleBinding.RoleRef.Kind)
		assert.Equal(t, expectedRoleName, roleBinding.RoleRef.Name)
		assert.Equal(t, "rbac.authorization.k8s.io", roleBinding.RoleRef.APIGroup)

		assert.Equal(t, 1, len(roleBinding.Subjects))
		assert.Equal(t, "User", roleBinding.Subjects[0].Kind)
		assert.Equal(t, "Admin", roleBinding.Subjects[0].Name)
		assert.Equal(t, "rbac.authorization.k8s.io", roleBinding.Subjects[0].APIGroup)
	})
}
