package main

import (
	"fmt"
	"io/ioutil"
	"nsxt-tag-it-all/utilities"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Need to include local packages as well as parameter input package.
/*
type importConfig struct {
	importConfig []nsxConfig `yaml:importconfig`
}
*/

type nsxConfig struct {
	Hostname   string `yaml:"hostname"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Importfile string `yaml:"importfile"`
}

func main() {
	//Need to validate that the config file is properly opened and parsed properly
	cfg := nsxConfig{}
	nsxTags := []utilities.NSXObjTags{}

	//Prepare regex parsers to determine the type of file that will be opened.
	yamlType, _ := regexp.Compile(".*yaml$")
	csvType, _ := regexp.Compile(".*csv$")

	cfgsrc, err := ioutil.ReadFile("config.yaml")

	//check to see if there was an error in opening the yaml file
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", cfgsrc)
	//Convert from yaml file to go struct
	err = yaml.Unmarshal([]byte(cfgsrc), &cfg)

	fmt.Println(cfg)
	//validate that the config yaml was able to be converted properly or exit with error.
	if err != nil {
		panic(err)
	}

	//cfg := config.importConfig[0]

	//Validate that the filetype in the config file is a ".yaml" or ".csv" file and if so import the file to the nsxTags Variable.
	fmt.Println("opening file: ", cfg.Importfile)

	if yamlType.MatchString(cfg.Importfile) {
		nsxTags = utilities.ImportYAMLTagsFromFile(cfg.Importfile)
	} else if csvType.MatchString(cfg.Importfile) {
		nsxTags = utilities.ImportCSVTagsFromFile(cfg.Importfile)
	} else {
		fmt.Println("The filename you have specified in the Config file does not mee the specific extensions supported by this program")
		fmt.Println("This program only supports \".yaml\" and \".csv\" file types at this time")
		fmt.Println("Please verify that your file is in the correct format and named properly.")
		panic("configuration error: file type in config.yaml:importfile invalid please correct.") //end the program if the file type is invalid
	}

	//Once the filetype has been validated and imported connect to the NSX manager
	nsxClient := utilities.ConnectNSXManager(cfg.Hostname, cfg.Username, cfg.Password)

	for _, objTags := range nsxTags {
		switch objTags.Objtype {
		case "Virtual Machine":
			utilities.TagVM(nsxClient, objTags.DisplayName, objTags.Tags)
		case "Logical Switch":
			utilities.TagLS(nsxClient, objTags.DisplayName, objTags.Tags)
		default:
			fmt.Printf("No NSX Object with type - %s - found\n", objTags.Objtype)
		}
	}
}
