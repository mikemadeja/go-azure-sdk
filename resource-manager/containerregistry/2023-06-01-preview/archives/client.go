package archives

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivesClient struct {
	Client *resourcemanager.Client
}

func NewArchivesClientWithBaseURI(sdkApi sdkEnv.Api) (*ArchivesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "archives", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ArchivesClient: %+v", err)
	}

	return &ArchivesClient{
		Client: client,
	}, nil
}
