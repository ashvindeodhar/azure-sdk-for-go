package vmwarecloudsimpleapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/vmwarecloudsimple/mgmt/2019-04-01/vmwarecloudsimple"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetOperationResultByRegion(ctx context.Context, operationID string) (result vmwarecloudsimple.OperationResource, err error)
	GetPrivateCloud(ctx context.Context, pcName string) (result vmwarecloudsimple.PrivateCloud, err error)
}

var _ BaseClientAPI = (*vmwarecloudsimple.BaseClient)(nil)

// AvailableOperationsClientAPI contains the set of methods on the AvailableOperationsClient type.
type AvailableOperationsClientAPI interface {
	List(ctx context.Context) (result vmwarecloudsimple.AvailableOperationsListResponsePage, err error)
}

var _ AvailableOperationsClientAPI = (*vmwarecloudsimple.AvailableOperationsClient)(nil)

// DedicatedCloudNodeClientAPI contains the set of methods on the DedicatedCloudNodeClient type.
type DedicatedCloudNodeClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest vmwarecloudsimple.DedicatedCloudNode) (result vmwarecloudsimple.DedicatedCloudNodeCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string) (result vmwarecloudsimple.DedicatedCloudNode, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result vmwarecloudsimple.DedicatedCloudNodeListResponsePage, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32, skipToken string) (result vmwarecloudsimple.DedicatedCloudNodeListResponsePage, err error)
	Update(ctx context.Context, resourceGroupName string, dedicatedCloudNodeName string, dedicatedCloudNodeRequest vmwarecloudsimple.PatchPayload) (result vmwarecloudsimple.DedicatedCloudNode, err error)
}

var _ DedicatedCloudNodeClientAPI = (*vmwarecloudsimple.DedicatedCloudNodeClient)(nil)

// DedicatedCloudServiceClientAPI contains the set of methods on the DedicatedCloudServiceClient type.
type DedicatedCloudServiceClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest vmwarecloudsimple.DedicatedCloudService) (result vmwarecloudsimple.DedicatedCloudService, err error)
	Delete(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string) (result vmwarecloudsimple.DedicatedCloudServiceDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string) (result vmwarecloudsimple.DedicatedCloudService, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result vmwarecloudsimple.DedicatedCloudServiceListResponsePage, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32, skipToken string) (result vmwarecloudsimple.DedicatedCloudServiceListResponsePage, err error)
	Update(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest vmwarecloudsimple.PatchPayload) (result vmwarecloudsimple.DedicatedCloudService, err error)
}

var _ DedicatedCloudServiceClientAPI = (*vmwarecloudsimple.DedicatedCloudServiceClient)(nil)

// SkusAvailabilityWithinRegionClientAPI contains the set of methods on the SkusAvailabilityWithinRegionClient type.
type SkusAvailabilityWithinRegionClientAPI interface {
	List(ctx context.Context, skuID string) (result vmwarecloudsimple.SkuAvailabilityListResponsePage, err error)
}

var _ SkusAvailabilityWithinRegionClientAPI = (*vmwarecloudsimple.SkusAvailabilityWithinRegionClient)(nil)

// PrivateCloudByRegionClientAPI contains the set of methods on the PrivateCloudByRegionClient type.
type PrivateCloudByRegionClientAPI interface {
	List(ctx context.Context) (result vmwarecloudsimple.PrivateCloudListPage, err error)
}

var _ PrivateCloudByRegionClientAPI = (*vmwarecloudsimple.PrivateCloudByRegionClient)(nil)

// ResourcePoolsByPCClientAPI contains the set of methods on the ResourcePoolsByPCClient type.
type ResourcePoolsByPCClientAPI interface {
	List(ctx context.Context, pcName string) (result vmwarecloudsimple.ResourcePoolsListResponsePage, err error)
}

var _ ResourcePoolsByPCClientAPI = (*vmwarecloudsimple.ResourcePoolsByPCClient)(nil)

// ResourcePoolByPCClientAPI contains the set of methods on the ResourcePoolByPCClient type.
type ResourcePoolByPCClientAPI interface {
	Get(ctx context.Context, pcName string, resourcePoolName string) (result vmwarecloudsimple.ResourcePool, err error)
}

var _ ResourcePoolByPCClientAPI = (*vmwarecloudsimple.ResourcePoolByPCClient)(nil)

// VirtualMachineTemplatesByPCClientAPI contains the set of methods on the VirtualMachineTemplatesByPCClient type.
type VirtualMachineTemplatesByPCClientAPI interface {
	List(ctx context.Context, pcName string, resourcePoolName string) (result vmwarecloudsimple.VirtualMachineTemplateListResponsePage, err error)
}

var _ VirtualMachineTemplatesByPCClientAPI = (*vmwarecloudsimple.VirtualMachineTemplatesByPCClient)(nil)

// VirtualMachineTemplateByPCClientAPI contains the set of methods on the VirtualMachineTemplateByPCClient type.
type VirtualMachineTemplateByPCClientAPI interface {
	Get(ctx context.Context, pcName string, virtualMachineTemplateName string) (result vmwarecloudsimple.VirtualMachineTemplate, err error)
}

var _ VirtualMachineTemplateByPCClientAPI = (*vmwarecloudsimple.VirtualMachineTemplateByPCClient)(nil)

// VirtualNetworksByPCClientAPI contains the set of methods on the VirtualNetworksByPCClient type.
type VirtualNetworksByPCClientAPI interface {
	List(ctx context.Context, pcName string, resourcePoolName string) (result vmwarecloudsimple.VirtualNetworkListResponsePage, err error)
}

var _ VirtualNetworksByPCClientAPI = (*vmwarecloudsimple.VirtualNetworksByPCClient)(nil)

// VirtualNetworkByPCClientAPI contains the set of methods on the VirtualNetworkByPCClient type.
type VirtualNetworkByPCClientAPI interface {
	Get(ctx context.Context, pcName string, virtualNetworkName string) (result vmwarecloudsimple.VirtualNetwork, err error)
}

var _ VirtualNetworkByPCClientAPI = (*vmwarecloudsimple.VirtualNetworkByPCClient)(nil)

// UsagesWithinRegionClientAPI contains the set of methods on the UsagesWithinRegionClient type.
type UsagesWithinRegionClientAPI interface {
	List(ctx context.Context, filter string) (result vmwarecloudsimple.UsageListResponsePage, err error)
}

var _ UsagesWithinRegionClientAPI = (*vmwarecloudsimple.UsagesWithinRegionClient)(nil)

// VirtualMachineClientAPI contains the set of methods on the VirtualMachineClient type.
type VirtualMachineClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineRequest vmwarecloudsimple.VirtualMachine) (result vmwarecloudsimple.VirtualMachineCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualMachineName string) (result vmwarecloudsimple.VirtualMachineDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualMachineName string) (result vmwarecloudsimple.VirtualMachine, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skipToken string) (result vmwarecloudsimple.VirtualMachineListResponsePage, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32, skipToken string) (result vmwarecloudsimple.VirtualMachineListResponsePage, err error)
	Start(ctx context.Context, resourceGroupName string, virtualMachineName string) (result vmwarecloudsimple.VirtualMachineStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, virtualMachineName string, mParameter *vmwarecloudsimple.VirtualMachineStopMode, mode vmwarecloudsimple.StopMode) (result vmwarecloudsimple.VirtualMachineStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, virtualMachineName string, virtualMachineRequest vmwarecloudsimple.PatchPayload) (result vmwarecloudsimple.VirtualMachineUpdateFuture, err error)
}

var _ VirtualMachineClientAPI = (*vmwarecloudsimple.VirtualMachineClient)(nil)