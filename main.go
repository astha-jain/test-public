package main

import (
	"fmt"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
)

var URL = "https://us-south.iaas.cloud.ibm.com/v1"

const IAMURL = "https://iam.cloud.ibm.com/identity/token"

var APIKey = "YOUR_KEY_here"

func main() {
	// Gen1 service
	_, vpcServiceErr := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
		URL: URL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: APIKey,
			URL:    IAMURL,
		},
	})
	if vpcServiceErr != nil {
		fmt.Println("Unable to create service for gen1.")
		fmt.Println("ServiceError -", vpcServiceErr)
	}
}
