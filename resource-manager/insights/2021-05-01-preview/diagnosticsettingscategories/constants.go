package diagnosticsettingscategories

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategoryType string

const (
	CategoryTypeLogs    CategoryType = "Logs"
	CategoryTypeMetrics CategoryType = "Metrics"
)

func PossibleValuesForCategoryType() []string {
	return []string{
		string(CategoryTypeLogs),
		string(CategoryTypeMetrics),
	}
}

func parseCategoryType(input string) (*CategoryType, error) {
	vals := map[string]CategoryType{
		"logs":    CategoryTypeLogs,
		"metrics": CategoryTypeMetrics,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CategoryType(input)
	return &out, nil
}
