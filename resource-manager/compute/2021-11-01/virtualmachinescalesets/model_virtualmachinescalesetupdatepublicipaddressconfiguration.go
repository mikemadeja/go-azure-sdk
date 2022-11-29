package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetUpdatePublicIPAddressConfiguration struct {
	Name       *string                                                             `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties `json:"properties,omitempty"`
}
