package trafficmanager

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// EndpointsClient is the client for the Endpoints methods of the Trafficmanager service.
type EndpointsClient struct {
	BaseClient
}

// NewEndpointsClient creates an instance of the EndpointsClient client.
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return NewEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEndpointsClientWithBaseURI creates an instance of the EndpointsClient client.
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return EndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Traffic Manager endpoint.
// Parameters:
// resourceGroupName - the name of the resource group containing the Traffic Manager endpoint to be created or
// updated.
// profileName - the name of the Traffic Manager profile.
// endpointType - the type of the Traffic Manager endpoint to be created or updated.
// endpointName - the name of the Traffic Manager endpoint to be created or updated.
// parameters - the Traffic Manager endpoint parameters supplied to the CreateOrUpdate operation.
func (client EndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (result Endpoint, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, profileName, endpointType, endpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EndpointsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EndpointsClient) CreateOrUpdateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Traffic Manager endpoint.
// Parameters:
// resourceGroupName - the name of the resource group containing the Traffic Manager endpoint to be deleted.
// profileName - the name of the Traffic Manager profile.
// endpointType - the type of the Traffic Manager endpoint to be deleted.
// endpointName - the name of the Traffic Manager endpoint to be deleted.
func (client EndpointsClient) Delete(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, profileName, endpointType, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EndpointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a Traffic Manager endpoint.
// Parameters:
// resourceGroupName - the name of the resource group containing the Traffic Manager endpoint.
// profileName - the name of the Traffic Manager profile.
// endpointType - the type of the Traffic Manager endpoint.
// endpointName - the name of the Traffic Manager endpoint.
func (client EndpointsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (result Endpoint, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, profileName, endpointType, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EndpointsClient) GetResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update a Traffic Manager endpoint.
// Parameters:
// resourceGroupName - the name of the resource group containing the Traffic Manager endpoint to be updated.
// profileName - the name of the Traffic Manager profile.
// endpointType - the type of the Traffic Manager endpoint to be updated.
// endpointName - the name of the Traffic Manager endpoint to be updated.
// parameters - the Traffic Manager endpoint parameters supplied to the Update operation.
func (client EndpointsClient) Update(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (result Endpoint, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, profileName, endpointType, endpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EndpointsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EndpointsClient) UpdateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
