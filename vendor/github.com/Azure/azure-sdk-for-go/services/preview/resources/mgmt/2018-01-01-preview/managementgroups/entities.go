package managementgroups

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EntitiesClient is the the Azure Management Groups API enables consolidation of multiple
// subscriptions/resources into an organizational hierarchy and centrally
// manage access control, policies, alerting and reporting for those resources.
type EntitiesClient struct {
	BaseClient
}

// NewEntitiesClient creates an instance of the EntitiesClient client.
func NewEntitiesClient(operationResultID string, skiptoken string) EntitiesClient {
	return NewEntitiesClientWithBaseURI(DefaultBaseURI, operationResultID, skiptoken)
}

// NewEntitiesClientWithBaseURI creates an instance of the EntitiesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEntitiesClientWithBaseURI(baseURI string, operationResultID string, skiptoken string) EntitiesClient {
	return EntitiesClient{NewWithBaseURI(baseURI, operationResultID, skiptoken)}
}

// List list all entities (Management Groups, Subscriptions, etc.) for the authenticated user.
// Parameters:
// groupName - a filter which allows the call to be filtered for a specific group.
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client EntitiesClient) List(ctx context.Context, groupName string, cacheControl string) (result EntityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntitiesClient.List")
		defer func() {
			sc := -1
			if result.elr.Response.Response != nil {
				sc = result.elr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, groupName, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.elr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", resp, "Failure sending request")
		return
	}

	result.elr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client EntitiesClient) ListPreparer(ctx context.Context, groupName string, cacheControl string) (*http.Request, error) {
	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(client.Skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", client.Skiptoken)
	}
	if len(groupName) > 0 {
		queryParameters["groupName"] = autorest.Encode("query", groupName)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Management/getEntities"),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client EntitiesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client EntitiesClient) ListResponder(resp *http.Response) (result EntityListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client EntitiesClient) listNextResults(ctx context.Context, lastResults EntityListResult) (result EntityListResult, err error) {
	req, err := lastResults.entityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.EntitiesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client EntitiesClient) ListComplete(ctx context.Context, groupName string, cacheControl string) (result EntityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EntitiesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, groupName, cacheControl)
	return
}