package targetcomputesizes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetComputeSizesClient struct {
	Client *resourcemanager.Client
}

func NewTargetComputeSizesClientWithBaseURI(sdkApi sdkEnv.Api) (*TargetComputeSizesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "targetcomputesizes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TargetComputeSizesClient: %+v", err)
	}

	return &TargetComputeSizesClient{
		Client: client,
	}, nil
}
