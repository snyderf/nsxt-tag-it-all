package utilities

import (
	"fmt"

	"github.com/vmware/go-vmware-nsxt"
)

func GetTZId(nsxClient *nsxt.APIClient, tzname string) string {

	tzList, resp, err := nsxClient.NetworkTransportApi.ListTransportZones(nsxClient.Context, nil)

	//validate that there are no errors in getting the transport zone list
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return ""
	}

	//loop through the transport zone list to find the ID that matches the transport zone ID
	for _, tz := range tzList.Results {
		if tz.DisplayName == tzname {
			return tz.Id
		}
	}

	return ""
}

func GetVMId(nsxClient *nsxt.APIClient, vmName string) string {

	searchOptions := map[string]interface{}{
		"displayName": vmName,
	}

	vmList, resp, err := nsxClient.FabricApi.ListVirtualMachines(nsxClient.Context, searchOptions)

	//validate that there are no errors in getting the VM list
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return ""
	}

	for _, vm := range vmList.Results {
		if vm.DisplayName == vmName {
			return vm.ExternalId
		}
	}

	return ""
}

func GetLSId(nsxClient *nsxt.APIClient, lsName string) string {

	lsList, resp, err := nsxClient.LogicalSwitchingApi.ListLogicalSwitches(nsxClient.Context, nil)

	//validate that there are no errors in getting the Logical Switch List
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return ""
	}

	// search through each logical switch that was returned and match the first one that has a Display Name equal to lsName
	for _, ls := range lsList.Results {
		if ls.DisplayName == lsName {
			return ls.Id
		}
	}

	return ""

}

func GetRtId(nsxClient *nsxt.APIClient, rtName string) string {

	rtList, resp, err := nsxClient.LogicalSwitchingApi.ListLogicalSwitches(nsxClient.Context, nil)

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return ""
	}

	for _, rt := range rtList.Results {
		if rt.DisplayName == rtName {
			return rt.Id
		}
	}

	return ""
}

func GetIPBlockId(nsxClient *nsxt.APIClient, blockName string) string {

	searchOptions := map[string]interface{}{
		"displayName": blockName,
	}

	netList, resp, err := nsxClient.PoolManagementApi.ListIpBlocks(nsxClient.Context, searchOptions)

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return ""
	}

	for _, net := range netList.Results {
		if net.DisplayName == blockName {
			return net.Id
		}
	}

	return ""

}

func GetNetBlockId(nsxClient *nsxt.APIClient, blockName string) string {

	searchOptions := map[string]interface{}{
		"displayName": blockName,
	}

	netList, resp, err := nsxClient.PoolManagementApi.ListIpBlockSubnets(nsxClient.Context, searchOptions)

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
		return ""
	}

	for _, net := range netList.Results {
		if net.DisplayName == blockName {
			return net.Id
		}
	}

	return ""

}
