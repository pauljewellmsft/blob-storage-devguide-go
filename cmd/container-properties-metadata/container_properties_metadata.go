package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// Blob dev guide container properties/metadata sample

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// <snippet_get_container_properties>
func getContainerProperties(client *azblob.Client, containerName string) {
	// Reference the container as a client object
	containerClient := client.ServiceClient().NewContainerClient(containerName)

	// Get the container properties
	resp, err := containerClient.GetProperties(context.TODO(), nil)
	handleError(err)

	// Print the container properties
	fmt.Printf("Blob public access: %v\n", *resp.BlobPublicAccess)
	fmt.Printf("Lease status: %v\n", *resp.LeaseStatus)
	fmt.Printf("Lease state: %v\n", *resp.LeaseState)
	fmt.Printf("Has immutability policy: %v\n", *resp.HasImmutabilityPolicy)
}

// </snippet_get_container_properties>

// <snippet_set_container_metadata>
func setContainerMetadata(client *azblob.Client, containerName string) {
	// Reference the container as a client object
	containerClient := client.ServiceClient().NewContainerClient(containerName)

	// Set the container metadata
	var metadata = make(map[string]*string)
	metadata["key1"] = to.Ptr("value1")
	metadata["key2"] = to.Ptr("value2")

	_, err := containerClient.SetMetadata(context.TODO(), nil)
	handleError(err)
}

// </snippet_set_container_metadata>

// <snippet_get_container_metadata>
func getContainerMetadata(client *azblob.Client, containerName string) {
	// Reference the container as a client object
	containerClient := client.ServiceClient().NewContainerClient(containerName)

	// Get the container properties, which includes metadata
	resp, err := containerClient.GetProperties(context.TODO(), nil)
	handleError(err)

	// Print the container metadata
	for k, v := range resp.Metadata {
		fmt.Printf("%v: %v\n", k, *v)
	}
}

// </snippet_get_container_metadata>

func main() {
	// TODO: replace <storage-account-name> with your actual storage account name
	url := "https://<storage-account-name>.blob.core.windows.net/"

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	client, err := azblob.NewClient(url, credential, nil)
	handleError(err)

	containerName := "sample-container"

	getContainerProperties(client, containerName)
	setContainerMetadata(client, containerName)
	getContainerMetadata(client, containerName)
}
