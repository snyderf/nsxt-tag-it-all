package utilities

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

/*
Import YAMLTagsFromFile will take a filename of a yaml file (*.yaml) and return an array of NSXObjTags
filename must be something ****.yaml for this work and the file itself must be in YAML format.
The return of NSXObjTags is designed to be a common return struct that can be easily parsed into any other function to
output the minimum needed parameters to TAG an NSX Object
*/
func ImportYAMLTagsFromFile(filename string) []NSXObjTags {

	var config YamlNSXTags

	src, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", src)
	err = yaml.Unmarshal(src, &config)
	if err != nil {
		panic(err)
	}

	//fmt.Println("value: ", config.NsxObject[0].Tags)

	return config.NsxObject

}
