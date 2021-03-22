package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestReadOnlyGroupRole(t *testing.T) {
	t.Parallel()

	expectedNamespace := "test"
	expectedRoleName := "read-only-role"
	expectedRoleBindingName := expectedRoleName + "-binding"

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/read-only-group-role",
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

	// Check the k8s cluster Role object
	k8sOptions := k8s.NewKubectlOptions("", "", expectedNamespace)
	role := k8s.GetRole(t, k8sOptions, expectedRoleName)

	// Check that there are 3 rules
	assert.Equal(t, 3, len(role.Rules))

	// Check rules have expected values
	apiGroups := [3]string{"", "extensions", "apps"}
	for index, element := range apiGroups {
		assert.Equal(t, element, role.Rules[index].APIGroups[0])
		assert.Equal(t, "*", role.Rules[index].Resources[0])
		assert.Equal(t, []string{"list", "watch", "get"}, role.Rules[index].Verbs)
	}

	// Check the RoleBinding.  No Terratest object exists, so we need to use the Kubernetes client
	roleBinding := getRoleBinding(t, k8sOptions, expectedNamespace, expectedRoleBindingName)

	assert.Equal(t, "Role", roleBinding.RoleRef.Kind)
	assert.Equal(t, expectedRoleName, roleBinding.RoleRef.Name)
	assert.Equal(t, "rbac.authorization.k8s.io", roleBinding.RoleRef.APIGroup)

	assert.Equal(t, 1, len(roleBinding.Subjects))
	assert.Equal(t, "Group", roleBinding.Subjects[0].Kind)
	assert.Equal(t, "Readers", roleBinding.Subjects[0].Name)
	assert.Equal(t, "rbac.authorization.k8s.io", roleBinding.Subjects[0].APIGroup)

}
