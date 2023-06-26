package protecteditems

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectedItem = AzureVMWorkloadSAPAseDatabaseProtectedItem{}

type AzureVMWorkloadSAPAseDatabaseProtectedItem struct {
	ExtendedInfo              *AzureVMWorkloadProtectedItemExtendedInfo `json:"extendedInfo,omitempty"`
	FriendlyName              *string                                   `json:"friendlyName,omitempty"`
	KpisHealths               *map[string]KPIResourceHealthDetails      `json:"kpisHealths,omitempty"`
	LastBackupErrorDetail     *ErrorDetail                              `json:"lastBackupErrorDetail,omitempty"`
	LastBackupStatus          *LastBackupStatus                         `json:"lastBackupStatus,omitempty"`
	LastBackupTime            *string                                   `json:"lastBackupTime,omitempty"`
	NodesList                 *[]DistributedNodesInfo                   `json:"nodesList,omitempty"`
	ParentName                *string                                   `json:"parentName,omitempty"`
	ParentType                *string                                   `json:"parentType,omitempty"`
	ProtectedItemDataSourceId *string                                   `json:"protectedItemDataSourceId,omitempty"`
	ProtectedItemHealthStatus *ProtectedItemHealthStatus                `json:"protectedItemHealthStatus,omitempty"`
	ProtectionState           *ProtectionState                          `json:"protectionState,omitempty"`
	ProtectionStatus          *string                                   `json:"protectionStatus,omitempty"`
	ServerName                *string                                   `json:"serverName,omitempty"`

	// Fields inherited from ProtectedItem
	BackupManagementType             *BackupManagementType `json:"backupManagementType,omitempty"`
	BackupSetName                    *string               `json:"backupSetName,omitempty"`
	ContainerName                    *string               `json:"containerName,omitempty"`
	CreateMode                       *CreateMode           `json:"createMode,omitempty"`
	DeferredDeleteTimeInUTC          *string               `json:"deferredDeleteTimeInUTC,omitempty"`
	DeferredDeleteTimeRemaining      *string               `json:"deferredDeleteTimeRemaining,omitempty"`
	IsArchiveEnabled                 *bool                 `json:"isArchiveEnabled,omitempty"`
	IsDeferredDeleteScheduleUpcoming *bool                 `json:"isDeferredDeleteScheduleUpcoming,omitempty"`
	IsRehydrate                      *bool                 `json:"isRehydrate,omitempty"`
	IsScheduledForDeferredDelete     *bool                 `json:"isScheduledForDeferredDelete,omitempty"`
	LastRecoveryPoint                *string               `json:"lastRecoveryPoint,omitempty"`
	PolicyId                         *string               `json:"policyId,omitempty"`
	PolicyName                       *string               `json:"policyName,omitempty"`
	ResourceGuardOperationRequests   *[]string             `json:"resourceGuardOperationRequests,omitempty"`
	SoftDeleteRetentionPeriodInDays  *int64                `json:"softDeleteRetentionPeriodInDays,omitempty"`
	SourceResourceId                 *string               `json:"sourceResourceId,omitempty"`
	WorkloadType                     *DataSourceType       `json:"workloadType,omitempty"`
}

func (o *AzureVMWorkloadSAPAseDatabaseProtectedItem) GetDeferredDeleteTimeInUTCAsTime() (*time.Time, error) {
	if o.DeferredDeleteTimeInUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeferredDeleteTimeInUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureVMWorkloadSAPAseDatabaseProtectedItem) SetDeferredDeleteTimeInUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeferredDeleteTimeInUTC = &formatted
}

func (o *AzureVMWorkloadSAPAseDatabaseProtectedItem) GetLastRecoveryPointAsTime() (*time.Time, error) {
	if o.LastRecoveryPoint == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRecoveryPoint, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureVMWorkloadSAPAseDatabaseProtectedItem) SetLastRecoveryPointAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRecoveryPoint = &formatted
}

var _ json.Marshaler = AzureVMWorkloadSAPAseDatabaseProtectedItem{}

func (s AzureVMWorkloadSAPAseDatabaseProtectedItem) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMWorkloadSAPAseDatabaseProtectedItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMWorkloadSAPAseDatabaseProtectedItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMWorkloadSAPAseDatabaseProtectedItem: %+v", err)
	}
	decoded["protectedItemType"] = "AzureVmWorkloadSAPAseDatabase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMWorkloadSAPAseDatabaseProtectedItem: %+v", err)
	}

	return encoded, nil
}
