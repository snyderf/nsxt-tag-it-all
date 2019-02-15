package utilities

import (
	"fmt"
	"nsxt-go/views"

	"github.com/vmware/go-vmware-nsxt"
	"github.com/vmware/go-vmware-nsxt/common"
	"github.com/vmware/go-vmware-nsxt/manager"
)

func TagVM(nsxClient *nsxt.APIClient, vmName string, vmTags []common.Tag) {

	//Get the OutsideID of the VM to be tagged.
	vmObj := views.GetVMObj(nsxClient, vmName)

	//pull the existing tags and aggregate them to the new tags.

	vmTags = append(vmObj.Tags, vmTags...)

	vmTagUpdate := manager.VirtualMachineTagUpdate{
		ExternalId: vmObj.ExternalId,
		Tags:       vmTags,
	}

	resp, err := nsxClient.FabricApi.UpdateVirtualMachineTagsUpdateTags(nsxClient.Context, vmTagUpdate)

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
	}

	fmt.Println("Successfully added Tags to VM: ", vmName)
	return
}

func TagLS(nsxClient *nsxt.APIClient, lsName string, lsTags []common.Tag) {

	//Get the ID of the Logical switch to be tagged
	lsID := views.GetLSId(nsxClient, lsName)

	//Get the logical switch data structure from the API
	ls, resp, err := nsxClient.LogicalSwitchingApi.GetLogicalSwitch(nsxClient.Context, lsID)

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
	}

	//Appdend the lsTags to the logical switch
	tagUpdate := append(ls.Tags, lsTags...)

	ls.Tags = tagUpdate

	ls, resp, err = nsxClient.LogicalSwitchingApi.UpdateLogicalSwitch(nsxClient.Context, lsID, ls)

	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
	}

	fmt.Printf("Logical switch: %s has been updated.\n", lsName)
}
