package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAdminUserRole(t *testing.T) {
	t.Parallel()

	expectedNamespace := "default"
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

	// Check the k8s cluster Role object
	k8sOptions := k8s.NewKubectlOptions("", "", expectedNamespace)
	role := k8s.GetRole(t, k8sOptions, expectedRoleName)

	// Check that there is one rule
	assert.Equal(t, 1, len(role.Rules))
	assert.Equal(t, "*", role.Rules[0].APIGroups[0])
	assert.Equal(t, "*", role.Rules[0].Resources[0])
	assert.Equal(t, "*", role.Rules[0].Verbs[0])

	// Check the RoleBinding.  No Terratest object exists, so we need to use the Kubernetes client
	roleBinding := getRoleBinding(t, k8sOptions, expectedNamespace, expectedRoleBindingName)

	assert.Equal(t, "Role", roleBinding.RoleRef.Kind)
	assert.Equal(t, expectedRoleName, roleBinding.RoleRef.Name)
	assert.Equal(t, "rbac.authorization.k8s.io", roleBinding.RoleRef.APIGroup)

	assert.Equal(t, 1, len(roleBinding.Subjects))
	assert.Equal(t, "User", roleBinding.Subjects[0].Kind)
	assert.Equal(t, "Admin", roleBinding.Subjects[0].Name)
	assert.Equal(t, "rbac.authorization.k8s.io", roleBinding.Subjects[0].APIGroup)

}

// Returns a RoleBinging object with the given name and namespace.
// Will fail the test if the RoleBinding does not exist.
func getRoleBinding(t *testing.T, options *k8s.KubectlOptions, namespace string, roleBindingName string) *v1.RoleBinding {
	clientset, err := k8s.GetKubernetesClientFromOptionsE(t, options)
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}

	roleBinding, err := clientset.RbacV1().
		RoleBindings(namespace).
		Get(context.Background(), roleBindingName, metav1.GetOptions{})
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}

	return roleBinding
}
