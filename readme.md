# NSXT-Tag-It-All

## Introduction:
This is a set of scripts that utilize the [go-vmware-nsxt](https://github.com/vmware/go-vmware-nsxt) to provide a mechanism for bulk tagging of NSX-T objects. 

## Prerequisites
Please ensure that you have Go installed on your machine. 
Also follow the instructions listed in [go-vmware-nsxt](https://github.com/vmware/go-vmware-nsxt) to install the go-vmware-nsxt libraries that this script uses. 
Lastly make sure you have your $GOPATH environment varibale set prior to running this.

## Usage

To run the applicaiton first configure the config.yaml file to list your NSX Managers IP or FQDN, Username, and Password, also add the file name of the .yaml or.csv file that contains your list of tags and the objects they are to be assigned to. Once the config file is completed issue the `go run main.go` command from within the same folder as the **main.go** file is.

## Currently supported objects for tagging.

At the current time the following objects are available for bulk tag addition. 

* Virtual Manchines
* Logical Switches

## Future Support

As this project develops more objects will be added. Check back from time to time or follow the repository to see all the new additions. 
