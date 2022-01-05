/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vmssvmclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"

	"sigs.k8s.io/cloud-provider-azure/pkg/retry"
)

const (
	// APIVersion is the API version for VMSS.
	APIVersion = "2020-12-01"
	// AzureStackCloudAPIVersion is the API version for Azure Stack
	AzureStackCloudAPIVersion = "2019-07-01"
	// AzureStackCloudName is the cloud name of Azure Stack
	AzureStackCloudName = "AZURESTACKCLOUD"
)

// Interface is the client interface for VirtualMachineScaleSetVM.
// Don't forget to run the following command to generate the mock client:
// mockgen -source=$GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/vmssvmclient/interface.go -package=mockvmssvmclient Interface > $GOPATH/src/sigs.k8s.io/cloud-provider-azure/pkg/azureclients/vmssvmclient/mockvmssvmclient/interface.go
type Interface interface {
	// Get gets a VirtualMachineScaleSetVM.
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, expand compute.InstanceViewTypes) (compute.VirtualMachineScaleSetVM, *retry.Error)

	// List gets a list of VirtualMachineScaleSetVMs in the virtualMachineScaleSet.
	List(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, expand string) ([]compute.VirtualMachineScaleSetVM, *retry.Error)

	// Update updates a VirtualMachineScaleSetVM.
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.VirtualMachineScaleSetVM, source string) *retry.Error

	// UpdateVMs updates a list of VirtualMachineScaleSetVM from map[instanceID]compute.VirtualMachineScaleSetVM.
	UpdateVMs(ctx context.Context, resourceGroupName string, VMScaleSetName string, instances map[string]compute.VirtualMachineScaleSetVM, source string, batchSize int) *retry.Error
}
