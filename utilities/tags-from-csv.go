package utilities

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/vmware/go-vmware-nsxt/common"
)

func ImportCSVTagsFromFile(filename string) []NSXObjTags {

	var config []NSXObjTags

	src, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	lines, err := csv.NewReader(src).ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", lines)

	for _, line := range lines {

		lineTags := []common.Tag{{Scope: line[2], Tag: line[3]}}

		lineconfig := NSXObjTags{
			Objtype:     line[0],
			DisplayName: line[1],
			Tags:        lineTags,
		}

		config = append(config, lineconfig)
	}

	//fmt.Println("value: ", config.NsxObject[0].Tags)

	return config

}
