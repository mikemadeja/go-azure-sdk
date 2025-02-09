package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDefinitionSummary struct {
	Effect                      *string         `json:"effect,omitempty"`
	PolicyDefinitionGroupNames  *[]string       `json:"policyDefinitionGroupNames,omitempty"`
	PolicyDefinitionId          *string         `json:"policyDefinitionId,omitempty"`
	PolicyDefinitionReferenceId *string         `json:"policyDefinitionReferenceId,omitempty"`
	Results                     *SummaryResults `json:"results,omitempty"`
}
