package utilities

import (
	"fmt"

	"github.com/vmware/go-vmware-nsxt"
)

func providerConnectivityCheck(nsxClient *nsxt.APIClient) error {
	// Connectivity check - get a random object to check connectivity and credentials
	// TODO(asarfaty): Use a list command which returns the full body, when the go vendor has one.
	_, httpResponse, err := nsxClient.ServicesApi.ReadLoadBalancerPool(nsxClient.Context, "Dummy")
	if err != nil {
		if httpResponse == nil || (httpResponse.StatusCode == 401 || httpResponse.StatusCode == 403) {
			return fmt.Errorf("NSXT provider connectivity check failed: %s", err)
		}
	}
	return nil
}

func ConnectNSXManager(host string, un string, pw string) *nsxt.APIClient {

	retriesConfig := nsxt.ClientRetriesConfiguration{
		MaxRetries:      3,
		RetryMinDelay:   20,
		RetryMaxDelay:   150,
		RetryOnStatuses: []int{429, 503},
	}

	cfg := nsxt.Configuration{
		BasePath:             "/api/v1",
		Host:                 host,
		Scheme:               "https",
		UserName:             un,
		Password:             pw,
		Insecure:             true,
		RetriesConfiguration: retriesConfig,
	}

	nsxClient, err := nsxt.NewAPIClient(&cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if providerConnectivityCheck(nsxClient) != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection to NSX-T Manager", host, "Successful")

	return nsxClient

}
