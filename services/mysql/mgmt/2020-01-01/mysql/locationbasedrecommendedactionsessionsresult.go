package mysql

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LocationBasedRecommendedActionSessionsResultClient is the the Microsoft Azure management API provides create, read,
// update, and delete functionality for Azure MySQL resources including servers, databases, firewall rules, VNET rules,
// log files and configurations with new business model.
type LocationBasedRecommendedActionSessionsResultClient struct {
	BaseClient
}

// NewLocationBasedRecommendedActionSessionsResultClient creates an instance of the
// LocationBasedRecommendedActionSessionsResultClient client.
func NewLocationBasedRecommendedActionSessionsResultClient(subscriptionID string) LocationBasedRecommendedActionSessionsResultClient {
	return NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI creates an instance of the
// LocationBasedRecommendedActionSessionsResultClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedRecommendedActionSessionsResultClient {
	return LocationBasedRecommendedActionSessionsResultClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List recommendation action session operation result.
// Parameters:
// locationName - the name of the location.
// operationID - the operation identifier.
func (client LocationBasedRecommendedActionSessionsResultClient) List(ctx context.Context, locationName string, operationID string) (result RecommendationActionsResultListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationBasedRecommendedActionSessionsResultClient.List")
		defer func() {
			sc := -1
			if result.rarl.Response.Response != nil {
				sc = result.rarl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mysql.LocationBasedRecommendedActionSessionsResultClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, locationName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsResultClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rarl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsResultClient", "List", resp, "Failure sending request")
		return
	}

	result.rarl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsResultClient", "List", resp, "Failure responding to request")
	}
	if result.rarl.hasNextLink() && result.rarl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client LocationBasedRecommendedActionSessionsResultClient) ListPreparer(ctx context.Context, locationName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/recommendedActionSessionsOperationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocationBasedRecommendedActionSessionsResultClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocationBasedRecommendedActionSessionsResultClient) ListResponder(resp *http.Response) (result RecommendationActionsResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LocationBasedRecommendedActionSessionsResultClient) listNextResults(ctx context.Context, lastResults RecommendationActionsResultList) (result RecommendationActionsResultList, err error) {
	req, err := lastResults.recommendationActionsResultListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsResultClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsResultClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.LocationBasedRecommendedActionSessionsResultClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocationBasedRecommendedActionSessionsResultClient) ListComplete(ctx context.Context, locationName string, operationID string) (result RecommendationActionsResultListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationBasedRecommendedActionSessionsResultClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, locationName, operationID)
	return
}