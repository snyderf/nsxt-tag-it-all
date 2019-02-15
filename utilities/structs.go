package utilities

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type NSXObjTags struct {
	Objtype     string       `yaml:"objtype"`
	DisplayName string       `yaml:"displayName"`
	Tags        []common.Tag `yaml:"tags"`
}

type YamlNSXTags struct {
	NsxObject []NSXObjTags `yaml:"nsxObject"`
}
