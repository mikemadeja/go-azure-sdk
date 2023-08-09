package migrations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationsClient struct {
	Client *resourcemanager.Client
}

func NewMigrationsClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migrations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrationsClient: %+v", err)
	}

	return &MigrationsClient{
		Client: client,
	}, nil
}
