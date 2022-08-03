package metrics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictiveValue struct {
	TimeStamp string  `json:"timeStamp"`
	Value     float64 `json:"value"`
}

func (o *PredictiveValue) GetTimeStampAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.TimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *PredictiveValue) SetTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeStamp = formatted
}
