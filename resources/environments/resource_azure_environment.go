// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"time"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource = &azureEnvironmentResource{}
)

type azureEnvironmentResource struct {
	client *cdp.Client
}

func NewAzureEnvironmentResource() resource.Resource {
	return &azureEnvironmentResource{}
}

func (r *azureEnvironmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_azure_environment"
}

func (r *azureEnvironmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = AzureEnvironmentSchema
}

func (r *azureEnvironmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *azureEnvironmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data azureEnvironmentResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to set plan")
		return
	}

	client := r.client.Environments

	params := operations.NewCreateAzureEnvironmentParamsWithContext(ctx)
	params.WithInput(ToAzureEnvrionmentRequest(ctx, &data))

	responseOk, err := client.Operations.CreateAzureEnvironment(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Azure Environment",
			"Got error while creating Azure Environment: "+err.Error(),
		)
		return
	}

	envResp := responseOk.Payload.Environment
	toAzureEnvrionmentResource(ctx, envResp, &data)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	timeout := time.Hour * 1
	if err := waitForEnvironmentToBeAvailable(data.ID.ValueString(), timeout, client, ctx); err != nil {
		resp.Diagnostics.AddError(
			"Error creating Azure Environment",
			"Failed to poll creating Azure Environment: "+err.Error(),
		)
		return
	}

	environmentName := data.EnvironmentName.ValueString()
	descParams := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	descParams.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(descParams)
	if err != nil {
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": data.ID.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error creating Azure Environment",
			"Could not read Azure Environment: "+data.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	toAzureEnvrionmentResource(ctx, descEnvResp.GetPayload().Environment, &data)
	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *azureEnvironmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state azureEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentName := state.EnvironmentName.ValueString()
	params := operations.NewDescribeEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DescribeEnvironmentRequest{
		EnvironmentName: &environmentName,
	})
	descEnvResp, err := r.client.Environments.Operations.DescribeEnvironment(params)
	if err != nil {
		if isEnvNotFoundError(err) {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Environment not found, removing from state.")
			tflog.Warn(ctx, "Environment not found, removing from state", map[string]interface{}{
				"id": state.ID.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading Azure Environment",
			"Could not read Azure Environment: "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	toAzureEnvrionmentResource(ctx, descEnvResp.GetPayload().Environment, &state)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func toAzureEnvrionmentResource(ctx context.Context, env *environmentsmodels.Environment, model *azureEnvironmentResourceModel) {
	model.ID = types.StringPointerValue(env.Crn)
	model.Crn = types.StringPointerValue(env.Crn)
	model.CredentialName = types.StringPointerValue(env.CredentialName)
	model.Description = types.StringValue(env.Description)
	model.EnableTunnel = types.BoolValue(env.TunnelEnabled)
	model.EnvironmentName = types.StringPointerValue(env.EnvironmentName)
	var freeIpaRecipes types.Set
	if env.Freeipa != nil {
		freeIpaRecipes, _ = types.SetValueFrom(ctx, types.StringType, env.Freeipa.Recipes)
	} else {
		freeIpaRecipes = types.SetNull(types.StringType)
	}
	model.FreeIpa, _ = types.ObjectValueFrom(ctx, map[string]attr.Type{
		"catalog":                 types.StringType,
		"image_id":                types.StringType,
		"instance_count_by_group": types.Int64Type,
		"instance_type":           types.StringType,
		"recipes":                 types.SetType{ElemType: types.StringType},
	}, &azureFreeIpaDetails{
		Recipes: freeIpaRecipes,
	})
	if env.LogStorage != nil {
		if env.LogStorage.AzureDetails != nil {
			model.LogStorage = &azureLogStorage{
				ManagedIdentity:     types.StringValue(env.LogStorage.AzureDetails.ManagedIdentity),
				StorageLocationBase: types.StringValue(env.LogStorage.AzureDetails.StorageLocationBase),
			}
			if env.BackupStorage != nil {
				if env.BackupStorage.AzureDetails != nil {
					model.LogStorage.BackupStorageLocationBase = types.StringValue(env.BackupStorage.AzureDetails.StorageLocationBase)
				}

			}
		}
	}
	if env.Network != nil {
		model.NewNetworkParams, _ = types.ObjectValueFrom(ctx, map[string]attr.Type{
			"network_cidr": types.StringType,
		}, &newNetworkParams{
			NetworkCidr: types.StringValue(env.Network.NetworkCidr),
		})
		if env.Network.Azure != nil {
			subnetIds, _ := types.SetValueFrom(ctx, types.StringType, env.Network.SubnetIds)
			model.ExistingNetworkParams, _ = types.ObjectValueFrom(ctx, map[string]attr.Type{
				"aks_private_dns_zone_id":      types.StringType,
				"database_private_dns_zone_id": types.StringType,
				"network_id":                   types.StringType,
				"resource_group_name":          types.StringType,
				"subnet_ids": types.SetType{
					ElemType: types.StringType,
				},
			}, &existingAzureNetwork{
				AksPrivateDNSZoneID:      types.StringValue(env.Network.Azure.AksPrivateDNSZoneID),
				DatabasePrivateDNSZoneID: types.StringValue(env.Network.Azure.DatabasePrivateDNSZoneID),
				NetworkID:                types.StringPointerValue(env.Network.Azure.NetworkID),
				ResourceGroupName:        types.StringPointerValue(env.Network.Azure.ResourceGroupName),
				SubnetIds:                subnetIds,
			})
			model.UsePublicIP = types.BoolPointerValue(env.Network.Azure.UsePublicIP)
		}
	}
	if env.ProxyConfig != nil {
		model.ProxyConfigName = types.StringPointerValue(env.ProxyConfig.ProxyConfigName)
	}
	model.PublicKey = types.StringValue(env.Authentication.PublicKey)
	model.Region = types.StringPointerValue(env.Region)
	model.ReportDeploymentLogs = types.BoolValue(env.ReportDeploymentLogs)
	if env.SecurityAccess != nil {
		var dsgIDs types.Set
		if model.SecurityAccess == nil || model.SecurityAccess.DefaultSecurityGroupIDs.IsUnknown() {
			dsgIDs = types.SetNull(types.StringType)
		} else {
			dsgIDs = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		var sgIDsknox types.Set
		if model.SecurityAccess == nil || model.SecurityAccess.SecurityGroupIDsForKnox.IsUnknown() {
			sgIDsknox = types.SetNull(types.StringType)
		} else {
			sgIDsknox = model.SecurityAccess.DefaultSecurityGroupIDs
		}
		model.SecurityAccess = &SecurityAccess{
			Cidr:                    types.StringValue(env.SecurityAccess.Cidr),
			DefaultSecurityGroupID:  types.StringValue(env.SecurityAccess.DefaultSecurityGroupID),
			DefaultSecurityGroupIDs: dsgIDs,
			SecurityGroupIDForKnox:  types.StringValue(env.SecurityAccess.SecurityGroupIDForKnox),
			SecurityGroupIDsForKnox: sgIDsknox,
		}
	}
	model.Status = types.StringPointerValue(env.Status)
	model.StatusReason = types.StringValue(env.StatusReason)
	if env.Tags != nil {
		merged := env.Tags.Defaults
		for k, v := range env.Tags.UserDefined {
			merged[k] = v
		}
		tagMap, _ := types.MapValueFrom(ctx, types.StringType, merged)
		model.Tags = tagMap
	}
	model.WorkloadAnalytics = types.BoolValue(env.WorkloadAnalytics)
}

func (r *azureEnvironmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *azureEnvironmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state azureEnvironmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentName := state.EnvironmentName.ValueString()
	params := operations.NewDeleteEnvironmentParamsWithContext(ctx)
	params.WithInput(&environmentsmodels.DeleteEnvironmentRequest{EnvironmentName: &environmentName})
	_, err := r.client.Environments.Operations.DeleteEnvironment(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Azure Environment",
			"Could not delete Azure Environment, unexpected error: "+err.Error(),
		)
		return
	}

	timeout := time.Hour * 1
	err = waitForEnvironmentToBeDeleted(environmentName, timeout, r.client.Environments, ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Azure Environment",
			"Failed to poll delete Azure Environment, unexpected error: "+err.Error(),
		)
		return
	}
}
