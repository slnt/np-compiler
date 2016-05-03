package print

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

// A MarshalFunc marshals an object into a formartz
type MarshalFunc func(interface{}) ([]byte, error)

// JSON prints too stdout as json
func JSON(obj interface{}) {
	print(obj, json.Marshal)
}

// YAML prints to stdout as yaml
func YAML(obj interface{}) {
	print(obj, yaml.Marshal)
}

func print(obj interface{}, marshaller MarshalFunc) {
	bytes, err := marshaller(obj)
	if err == nil {
		fmt.Println(string(bytes))
	}

}
