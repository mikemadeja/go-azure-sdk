package listprivatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourcesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type PrivateLinkResourcesListCompleteResult struct {
	Items []PrivateLinkResource
}

// PrivateLinkResourcesList ...
func (c ListPrivateLinkResourcesClient) PrivateLinkResourcesList(ctx context.Context, id VaultId) (result PrivateLinkResourcesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/privateLinkResources", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateLinkResourcesListComplete retrieves all the results into a single object
func (c ListPrivateLinkResourcesClient) PrivateLinkResourcesListComplete(ctx context.Context, id VaultId) (PrivateLinkResourcesListCompleteResult, error) {
	return c.PrivateLinkResourcesListCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// PrivateLinkResourcesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ListPrivateLinkResourcesClient) PrivateLinkResourcesListCompleteMatchingPredicate(ctx context.Context, id VaultId, predicate PrivateLinkResourceOperationPredicate) (result PrivateLinkResourcesListCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.PrivateLinkResourcesList(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = PrivateLinkResourcesListCompleteResult{
		Items: items,
	}
	return
}
