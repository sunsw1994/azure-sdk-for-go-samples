// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package synapse

import (
	"context"
	"time"

	"synapse/2019-11-01-preview/monitoring"
	"synapse/2019-11-01-preview/spark"
	"synapse/2020-02-01-preview/accesscontrol"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/iam"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/util"
	"github.com/Azure/azure-sdk-for-go/services/preview/synapse/mgmt/2019-06-01-preview/synapse"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
)

func getMonitoring() error {
	monitoring := monitoring.NewClient("https://shangweisynapseworkspace.dev.azuresynapse.net")
	monitoring.Authorizer, _ = auth.NewAuthorizerFromEnvironmentWithResource("https://dev.azuresynapse.net")

	_, err := monitoring.GetSQLJobQueryString(context.Background(), "2019-11-01-preview", "", "", "", "")

	// accessclient := accesscontrol.NewClient("https://shangweisynapseworkspace.dev.azuresynapse.net")
	// accessclient.Authorizer, _ = auth.NewAuthorizerFromEnvironmentWithResource("https://dev.azuresynapse.net")

	// _, err := accessclient.GetRoleDefinitionByID(context.Background(), "6e4bf58a-b8e1-4cc3-bbf9-d73143322b78")

	// _, err := accessclient.CreateRoleAssignment(context.Background(), accesscontrol.RoleAssignmentOptions{
	// 	RoleID:      to.StringPtr("6e4bf58a-b8e1-4cc3-bbf9-d73143322b78"),
	// 	PrincipalID: to.StringPtr("cf19cce5-d9ad-4456-9522-d63c69f27fb2"),
	// })

	//_, err := accessclient.DeleteRoleAssignmentByID(context.Background(), "6e4bf58a-b8e1-4cc3-bbf9-d73143322b78-bf4f36df-eda0-4fc4-9602-69d6ff714d8c")
	return err
}

func getAccessControl() error {
	accessclient := accesscontrol.NewClient("https://shangweisynapseworkspace.dev.azuresynapse.net")
	accessclient.Authorizer, _ = auth.NewAuthorizerFromEnvironmentWithResource("https://dev.azuresynapse.net")

	_, err := accessclient.GetRoleDefinitionByID(context.Background(), "6e4bf58a-b8e1-4cc3-bbf9-d73143322b78")

	// _, err := accessclient.CreateRoleAssignment(context.Background(), accesscontrol.RoleAssignmentOptions{
	// 	RoleID:      to.StringPtr("6e4bf58a-b8e1-4cc3-bbf9-d73143322b78"),
	// 	PrincipalID: to.StringPtr("cf19cce5-d9ad-4456-9522-d63c69f27fb2"),
	// })

	//_, err := accessclient.DeleteRoleAssignmentByID(context.Background(), "6e4bf58a-b8e1-4cc3-bbf9-d73143322b78-bf4f36df-eda0-4fc4-9602-69d6ff714d8c")
	return err
}

func get() error {
	accessclient := accesscontrol.NewClient("https://shangweisynapseworkspace.dev.azuresynapse.net")
	accessclient.Authorizer, _ = auth.NewAuthorizerFromEnvironmentWithResource("https://dev.azuresynapse.net")

	_, err := accessclient.GetRoleDefinitionByID(context.Background(), "6e4bf58a-b8e1-4cc3-bbf9-d73143322b78")

	// _, err := accessclient.CreateRoleAssignment(context.Background(), accesscontrol.RoleAssignmentOptions{
	// 	RoleID:      to.StringPtr("6e4bf58a-b8e1-4cc3-bbf9-d73143322b78"),
	// 	PrincipalID: to.StringPtr("cf19cce5-d9ad-4456-9522-d63c69f27fb2"),
	// })

	//_, err := accessclient.DeleteRoleAssignmentByID(context.Background(), "6e4bf58a-b8e1-4cc3-bbf9-d73143322b78-bf4f36df-eda0-4fc4-9602-69d6ff714d8c")
	return err
}

func getSparkClient() (*spark.BatchJob, error) {

	client := spark.NewBatchClient("https://shangweisynapseworkspace.dev.azuresynapse.net", "sparkpool")

	client.Authorizer, _ = auth.NewAuthorizerFromEnvironmentWithResource("https://dev.azuresynapse.net")

	sparkbatch, err := client.CreateSparkBatchJob(context.Background(), spark.BatchJobOptions{
		File:           to.StringPtr("abfss://shangwei@shangweiadlsgen2.dfs.core.windows.net/synapse/workspaces/shangweisynapseworkspace/sparkpools/sparkpool/spark-examples.jar"),
		ClassName:      to.StringPtr("org.apache.spark.examples.SparkPi"),
		ExecutorCores:  to.Int32Ptr(2),
		ExecutorMemory: to.StringPtr("8g"),
		Arguments:      to.StringSlicePtr([]string{"10"}),
		Name:           to.StringPtr("gosdkpi"),
		ExecutorCount:  to.Int32Ptr(2),
		DriverCores:    to.Int32Ptr(2),
		DriverMemory:   to.StringPtr("8g"),
	}, to.BoolPtr(false))
	// sparkbatch, err := client.GetSparkBatchJob(context.Background(), 0, to.BoolPtr(true))

	return &sparkbatch, err
}

func getAccessControlClient() {
	accesscontrol.NewClient("")
}

func getWorkspaceClient() (*synapse.WorkspacesClient, error) {
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return nil, err
	}
	client := synapse.NewWorkspacesClient(config.SubscriptionID())
	client.Authorizer = a
	client.AddToUserAgent(config.UserAgent())
	return &client, nil
}

func CreateWorkspace(resourceGroup, workspaceName string) (*synapse.Workspace, error) {
	client, err := getWorkspaceClient()
	if err != nil {
		return nil, err
	}
	client.PollingDuration = 20 * time.Minute
	util.PrintAndLog("creating synapse workspace")

	future, err := client.CreateOrUpdate(context.Background(), resourceGroup, workspaceName, synapse.Workspace{
		Location: to.StringPtr(config.Location()),
		WorkspaceProperties: &synapse.WorkspaceProperties{
			DefaultDataLakeStorage: &synapse.DataLakeStorageAccountDetails{
				AccountURL: to.StringPtr("https://shangweiadlsgen2.dfs.core.windows.net/"),
				Filesystem: to.StringPtr("shangwei"),
			},
			SQLAdministratorLoginPassword: to.StringPtr("Password1!"),
			SQLAdministratorLogin:         to.StringPtr("sqladminuser"),
		},
		Identity: &synapse.ManagedIdentity{
			Type: synapse.ResourceIdentityTypeSystemAssigned,
		},
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to create workspace")
	}
	util.PrintAndLog("waiting for synapse workspace to finish deploying, this will take a while...")
	err = future.WaitForCompletionRef(context.Background(), client.Client)
	if err != nil {
		return nil, errors.Wrap(err, "failed waiting for synapse workspace creation")
	}
	c, err := future.Result(*client)
	return &c, err
}
