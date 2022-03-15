// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package handlers

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/mgmt/keyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/gofrs/uuid"
	"github.com/project-radius/radius/pkg/azure/armauth"
	"github.com/project-radius/radius/pkg/azure/azresources"
	"github.com/project-radius/radius/pkg/azure/clients"
	"github.com/project-radius/radius/pkg/healthcontract"
	"github.com/project-radius/radius/pkg/keys"
	"github.com/project-radius/radius/pkg/resourcemodel"
)

const (
	KeyVaultURIKey  = "uri"
	KeyVaultNameKey = "keyvaultname"
	KeyVaultIDKey   = "keyvaultid"
)

func NewAzureKeyVaultHandler(arm *armauth.ArmConfig) ResourceHandler {
	return &azureKeyVaultHandler{arm: arm}
}

type azureKeyVaultHandler struct {
	arm *armauth.ArmConfig
}

func (handler *azureKeyVaultHandler) Put(ctx context.Context, options *PutOptions) (map[string]string, error) {
	properties := mergeProperties(*options.Resource, options.ExistingOutputResource)

	// This assertion is important so we don't start creating/modifying a resource
	err := ValidateResourceIDsForResource(properties, KeyVaultIDKey)
	if err != nil {
		return nil, err
	}

	// This is mostly called for the side-effect of verifying that the keyvault exists.
	kv, err := handler.GetKeyVaultByID(ctx, properties[KeyVaultIDKey])
	if err != nil {
		return nil, err
	}

	properties[KeyVaultURIKey] = *kv.Properties.VaultURI

	if options.Resource.Deployed {
		return properties, nil
	}

	options.Resource.Identity = resourcemodel.NewARMIdentity(properties[KeyVaultIDKey], clients.GetAPIVersionFromUserAgent(keyvault.UserAgent()))
	return properties, nil
}

func (handler *azureKeyVaultHandler) Delete(ctx context.Context, options DeleteOptions) error {
	return nil
}

func (handler *azureKeyVaultHandler) GetKeyVaultByID(ctx context.Context, id string) (*keyvault.Vault, error) {
	parsed, err := azresources.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse KeyVault resource id: %w", err)
	}

	kvc := clients.NewVaultsClient(handler.arm.SubscriptionID, handler.arm.Auth)

	kv, err := kvc.Get(ctx, parsed.ResourceGroup, parsed.Types[0].Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get KeyVault: %w", err)
	}

	return &kv, nil
}

func (handler *azureKeyVaultHandler) CreateKeyVault(ctx context.Context, vaultName string, options PutOptions) (*keyvault.Vault, error) {
	kvc := clients.NewVaultsClient(handler.arm.SubscriptionID, handler.arm.Auth)

	sc := clients.NewSubscriptionsClient(handler.arm.Auth)

	s, err := sc.Get(ctx, handler.arm.SubscriptionID)
	if err != nil {
		return nil, fmt.Errorf("unable to find subscription: %w", err)
	}
	tenantID, err := uuid.FromString(*s.TenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert tenantID to UUID: %w", err)
	}

	location, err := clients.GetResourceGroupLocation(ctx, *handler.arm)
	if err != nil {
		return nil, err
	}

	vaultsFuture, err := kvc.CreateOrUpdate(
		ctx,
		handler.arm.ResourceGroup,
		vaultName,
		keyvault.VaultCreateOrUpdateParameters{
			Location: location,
			Properties: &keyvault.VaultProperties{
				TenantID: &tenantID,
				Sku: &keyvault.Sku{
					Family: to.StringPtr("A"),
					Name:   keyvault.Standard,
				},
				EnableRbacAuthorization: to.BoolPtr(true),
			},
			Tags: keys.MakeTagsForRadiusResource(options.ApplicationName, options.ResourceName),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to PUT keyvault: %w", err)
	}

	err = vaultsFuture.WaitForCompletionRef(ctx, kvc.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to PUT keyvault: %w", err)
	}

	kv, err := vaultsFuture.Result(kvc)
	if err != nil {
		return nil, fmt.Errorf("failed to PUT keyvault: %w", err)
	}

	return &kv, nil
}

func (handler *azureKeyVaultHandler) DeleteKeyVault(ctx context.Context, vaultName string) error {
	kvClient := clients.NewVaultsClient(handler.arm.SubscriptionID, handler.arm.Auth)

	_, err := kvClient.Delete(ctx, handler.arm.ResourceGroup, vaultName)
	if err != nil {
		return fmt.Errorf("failed to DELETE keyvault: %w", err)
	}

	return nil
}

func NewAzureKeyVaultHealthHandler(arm *armauth.ArmConfig) HealthHandler {
	return &azureKeyVaultHealthHandler{arm: arm}
}

type azureKeyVaultHealthHandler struct {
	arm *armauth.ArmConfig
}

func (handler *azureKeyVaultHealthHandler) GetHealthOptions(ctx context.Context) healthcontract.HealthCheckOptions {
	return healthcontract.HealthCheckOptions{}
}
