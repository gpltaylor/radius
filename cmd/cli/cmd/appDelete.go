// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package cmd

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/radius/cmd/cli/utils"
	"github.com/Azure/radius/pkg/radclient"
	"github.com/spf13/cobra"
)

// appDeleteCmd command to delete an application
var appDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete RAD application",
	Long:  "Delete the specified RAD application deployed in the default environment",
	RunE:  deleteApplication,
}

func init() {
	applicationCmd.AddCommand(appDeleteCmd)

	appDeleteCmd.Flags().StringP("name", "n", "", "The application name")
	if err := appDeleteCmd.MarkFlagRequired("name"); err != nil {
		fmt.Printf("Failed to mark the name flag required: %v", err)
	}
}

func deleteApplication(cmd *cobra.Command, args []string) error {
	applicationName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	env, err := validateEnvironment()
	if err != nil {
		return err
	}

	azcred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("Failed to obtain Azure credential: %w", err)
	}

	con := armcore.NewDefaultConnection(azcred, nil)

	// Delete deployments: An application can have multiple deployments in it that should be deleted before the application can be deleted.
	dc := radclient.NewDeploymentClient(con, env.SubscriptionID)

	// Retrieve all the deployments in the application
	response, err := dc.ListByApplication(cmd.Context(), env.ResourceGroup, applicationName, nil)
	if err != nil {
		return err
	}

	// Delete the deployments
	deploymentResources := *response.DeploymentList
	for _, deploymentResource := range *deploymentResources.Value {
		// This is needed until server side implementation is fixed https://github.com/Azure/radius/issues/159
		deploymentName := utils.GetResourceNameFromFullyQualifiedPath(*deploymentResource.Name)

		_, err := dc.Delete(cmd.Context(), env.ResourceGroup, applicationName, deploymentName, nil)
		if err != nil {
			return fmt.Errorf("Failed to delete the deployment %s, %w", deploymentName, err)
		}
		fmt.Printf("Deleted deployment '%s'\n", deploymentName)
	}

	// Delete application
	ac := radclient.NewApplicationClient(con, env.SubscriptionID)

	_, err = ac.Delete(cmd.Context(), env.ResourceGroup, applicationName, nil)
	if err != nil {
		return fmt.Errorf("Failed to delete the application %w", err)
	}
	fmt.Printf("Application '%s' has been deleted\n", applicationName)

	return err
}
