package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "3b36c431-a92e-4ac8-927b-17a0f4b30054"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "<your-college-id>",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))
}

// GetVirtualMachineNics checks if the NIC exists in the VM.
// This function would fail the test if there is an error.
//func GetVirtualMachineNics(t *testing.T, vmName string, resGroupName string, subscriptionID string) []string {
//	nicList, err := azure.GetVirtualMachineNicsE(vmName, resGroupName, subscriptionID)
//	require.NoError(err)

//	return nicList
//}

// GetVirtualMachineNicsE checks if the VM runs the correct Ubuntu version.
/*func GetVirtualMachineNicsE(vmName string, resGroupName string, subscriptionID string) ([]string, error) {

	// Get VM Object
	vm, err := azure.GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get VM NIC(s); value always present, no nil checks needed.
	vmNICs := *vm.NetworkProfile.NetworkInterfaces

	nics := make([]string, len(vmNICs))
	for i, nic := range vmNICs {
		// Get ID from resource string.
		nicName := azure.GetNameFromResourceID(*nic.ID)
		if err == nil {
			nics[i] = nicName
		}
	}
	return nics, nil
}

// ComfirmIfVMRunningAndCorrectUbuntu checks if the VM runs the correct Ubuntu version.
func ComfirmIfVMRunningAndCorrectUbuntu(t *testing.T, vmName string, resGroupName string, subscriptionID string) ([]string, error) {
	// Get VM Object
	vm, err := azure.GetVirtualMachineE(vmName, resGroupName, subscriptionID)
	if err != nil {
		return nil, err
	}

	require.Equal(t, "Canonical", *vm.StorageProfile.ImageReference.Publisher)
    require.Equal(t, "UbuntuServer", *vm.StorageProfile.ImageReference.Offer)
    require.Equal(t, "18.04-LTS", *vm.StorageProfile.ImageReference.Sku) // Adjust as needed

	return nil, err
}*/