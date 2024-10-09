//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20231001preview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	internal *arm.Client
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
	cl, err := arm.NewClient(moduleName+".LocationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocationsClient{
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a location. The location resource represents a logical location where the resource
// provider operates.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - resourceProviderName - The resource provider name. This is also the resource provider namespace. Example: 'Applications.Datastores'.
//   - locationName - The location name.
//   - resource - Resource create parameters.
//   - options - LocationsClientBeginCreateOrUpdateOptions contains the optional parameters for the LocationsClient.BeginCreateOrUpdate
//     method.
func (client *LocationsClient) BeginCreateOrUpdate(ctx context.Context, planeName string, resourceProviderName string, locationName string, resource LocationResource, options *LocationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[LocationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, planeName, resourceProviderName, locationName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LocationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LocationsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a location. The location resource represents a logical location where the resource provider
// operates.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *LocationsClient) createOrUpdate(ctx context.Context, planeName string, resourceProviderName string, locationName string, resource LocationResource, options *LocationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, planeName, resourceProviderName, locationName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LocationsClient) createOrUpdateCreateRequest(ctx context.Context, planeName string, resourceProviderName string, locationName string, resource LocationResource, options *LocationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/planes/radius/{planeName}/providers/System.Resources/resourceproviders/{resourceProviderName}/locations/{locationName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Delete a location. The location resource represents a logical location where the resource provider operates.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - resourceProviderName - The resource provider name. This is also the resource provider namespace. Example: 'Applications.Datastores'.
//   - locationName - The location name.
//   - options - LocationsClientBeginDeleteOptions contains the optional parameters for the LocationsClient.BeginDelete method.
func (client *LocationsClient) BeginDelete(ctx context.Context, planeName string, resourceProviderName string, locationName string, options *LocationsClientBeginDeleteOptions) (*runtime.Poller[LocationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, planeName, resourceProviderName, locationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LocationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[LocationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a location. The location resource represents a logical location where the resource provider operates.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *LocationsClient) deleteOperation(ctx context.Context, planeName string, resourceProviderName string, locationName string, options *LocationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, planeName, resourceProviderName, locationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LocationsClient) deleteCreateRequest(ctx context.Context, planeName string, resourceProviderName string, locationName string, options *LocationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/planes/radius/{planeName}/providers/System.Resources/resourceproviders/{resourceProviderName}/locations/{locationName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified location. The location resource represents a logical location where the resource provider operates.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - resourceProviderName - The resource provider name. This is also the resource provider namespace. Example: 'Applications.Datastores'.
//   - locationName - The location name.
//   - options - LocationsClientGetOptions contains the optional parameters for the LocationsClient.Get method.
func (client *LocationsClient) Get(ctx context.Context, planeName string, resourceProviderName string, locationName string, options *LocationsClientGetOptions) (LocationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, planeName, resourceProviderName, locationName, options)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LocationsClient) getCreateRequest(ctx context.Context, planeName string, resourceProviderName string, locationName string, options *LocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/planes/radius/{planeName}/providers/System.Resources/resourceproviders/{resourceProviderName}/locations/{locationName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LocationsClient) getHandleResponse(resp *http.Response) (LocationsClientGetResponse, error) {
	result := LocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationResource); err != nil {
		return LocationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List available locations for the specified resource provider.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - resourceProviderName - The resource provider name. This is also the resource provider namespace. Example: 'Applications.Datastores'.
//   - options - LocationsClientListOptions contains the optional parameters for the LocationsClient.NewListPager method.
func (client *LocationsClient) NewListPager(planeName string, resourceProviderName string, options *LocationsClientListOptions) (*runtime.Pager[LocationsClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[LocationsClientListResponse]{
		More: func(page LocationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationsClientListResponse) (LocationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, planeName, resourceProviderName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LocationsClient) listCreateRequest(ctx context.Context, planeName string, resourceProviderName string, options *LocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/planes/radius/{planeName}/providers/System.Resources/resourceproviders/{resourceProviderName}/locations"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	if resourceProviderName == "" {
		return nil, errors.New("parameter resourceProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderName}", url.PathEscape(resourceProviderName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationsClient) listHandleResponse(resp *http.Response) (LocationsClientListResponse, error) {
	result := LocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationResourceListResult); err != nil {
		return LocationsClientListResponse{}, err
	}
	return result, nil
}

