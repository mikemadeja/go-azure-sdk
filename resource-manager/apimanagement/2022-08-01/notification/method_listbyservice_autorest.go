package notification

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]NotificationContract

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByServiceOperationResponse, error)
}

type ListByServiceCompleteResult struct {
	Items []NotificationContract
}

func (r ListByServiceOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByServiceOperationResponse) LoadMore(ctx context.Context) (resp ListByServiceOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByServiceOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListByServiceOperationOptions() ListByServiceOperationOptions {
	return ListByServiceOperationOptions{}
}

func (o ListByServiceOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByServiceOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByService ...
func (c NotificationClient) ListByService(ctx context.Context, id ServiceId, options ListByServiceOperationOptions) (resp ListByServiceOperationResponse, err error) {
	req, err := c.preparerForListByService(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notification.NotificationClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "notification.NotificationClient", "ListByService", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByService(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "notification.NotificationClient", "ListByService", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByService prepares the ListByService request.
func (c NotificationClient) preparerForListByService(ctx context.Context, id ServiceId, options ListByServiceOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/notifications", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByServiceWithNextLink prepares the ListByService request with the given nextLink token.
func (c NotificationClient) preparerForListByServiceWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByService handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (c NotificationClient) responderForListByService(resp *http.Response) (result ListByServiceOperationResponse, err error) {
	type page struct {
		Values   []NotificationContract `json:"value"`
		NextLink *string                `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByServiceOperationResponse, err error) {
			req, err := c.preparerForListByServiceWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "notification.NotificationClient", "ListByService", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "notification.NotificationClient", "ListByService", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByService(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "notification.NotificationClient", "ListByService", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByServiceComplete retrieves all of the results into a single object
func (c NotificationClient) ListByServiceComplete(ctx context.Context, id ServiceId, options ListByServiceOperationOptions) (ListByServiceCompleteResult, error) {
	return c.ListByServiceCompleteMatchingPredicate(ctx, id, options, NotificationContractOperationPredicate{})
}

// ListByServiceCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c NotificationClient) ListByServiceCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ListByServiceOperationOptions, predicate NotificationContractOperationPredicate) (resp ListByServiceCompleteResult, err error) {
	items := make([]NotificationContract, 0)

	page, err := c.ListByService(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByServiceCompleteResult{
		Items: items,
	}
	return out, nil
}
