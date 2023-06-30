package usages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageList struct {
	NextLink *string  `json:"nextLink,omitempty"`
	Value    *[]Usage `json:"value,omitempty"`
}
