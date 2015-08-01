package stack

import (
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

const (
	SERVICE_SCHEMA = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "id": "service.json",
    "type": "object",
    "additionalProperties": false,
    "required": [
        "image"
    ],
    "properties": {
        "image": {
            "type": "string"
        },
        "command": {
            "type": "string"
        },
        "links": {
            "items": {
                "type": "string",
             	 "pattern": "^[0-9a-zA-Z]{1,20}(:[0-9a-zA-Z]{1,20})?$"
            }
        },
        "ports": {
            "items": {
                "type": "string",
             	 "pattern": "^[1-9][0-9]{0,4}(:[1-9][0-9]{0,4})?$"
            }
        },
        "volumes": {
            "items": {
                "type": "string"
            }
        },
        "volumes_from": {
            "items": {
                "type": "string",
             	 "pattern": "^[0-9a-zA-Z]+$"
            }
        },
        "environment": {
            "items": {
                "type": "string",
             	 "pattern": "^[0-9a-zA-Z_]+=?"
            }
        },
        "entrypoint": {
            "type": "string"
        },
        "mem_limit": {
            "type": "string",
           	"pattern": "^"
        },
        "privileged": {
            "type": "boolean"
        },
        "restart": {
            "type": "string",
           	"pattern": "^(no|always|on-failure(:[0-9]+)?)$"
        },
        "strategy": {
            "type": "string",
            "enum": [ "balance", "every_node" ]
        },
        "scale": {
            "type": "integer",
			"minimum": 1
        },
        "tags": {
            "items": {
                "type": "string",
                "pattern": "^[0-9a-zA-Z/-]{1,20}$"
            }
        }
    }
}`
	NAME_SCHEMA = `
{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "id": "name.json",
    "type": "string",
   	"pattern": "^[0-9a-zA-Z]{1,25}$"
}`
)

func Validate(stackName string, filePath string) error {
	stack := Stack{Name: ""}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	yaml.Unmarshal(data, &stack.Definition)

	nameLoader := gojsonschema.NewStringLoader(NAME_SCHEMA)
	schemaLoader := gojsonschema.NewStringLoader(SERVICE_SCHEMA)

	documentLoader := gojsonschema.NewGoLoader(stackName)
	result, err := gojsonschema.Validate(nameLoader, documentLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		message := "Stack name not valid: "
		for _, desc := range result.Errors() {
			message += fmt.Sprintf("'%s' %s\n", desc.Value, desc.Description)
		}
		return errors.New(message)
	}

	stack.SetDefaults()

	for name, definition := range stack.Definition {
		documentLoader := gojsonschema.NewGoLoader(name)
		result, err := gojsonschema.Validate(nameLoader, documentLoader)
		if err != nil {
			return err
		}

		if !result.Valid() {
			var message string
			message = "Service name not valid: "
			for _, desc := range result.Errors() {
				message += fmt.Sprintf("'%s' %s\n", desc.Value, desc.Description)
			}
			return errors.New(message)
		}

		documentLoader = gojsonschema.NewGoLoader(definition)
		result, err = gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			return err
		}

		if !result.Valid() {
			var message string
			message = "Service " + name + " not valid. see errors:\n"
			for _, desc := range result.Errors() {
				message += fmt.Sprintf("%s: '%s' %s\n", desc.Context.String()[7:], desc.Value, desc.Description)
			}
			return errors.New(message)
		}

		for _, link := range definition.Links {
			linkAndAlias := strings.Split(link, ":")
			var linkedService string
			if len(linkAndAlias) == 2 {
				linkedService = linkAndAlias[1]
			} else {
				linkedService = linkAndAlias[0]
			}
			_, ok := stack.Definition[linkedService]
			if !ok {
				return fmt.Errorf("Linked service '%s' in service '%s' does not exist\n", linkedService, name)
			}
		}

		for _, volumesFrom := range definition.Volumes_from {
			_, ok := stack.Definition[volumesFrom]
			if !ok {
				return fmt.Errorf("VolumesFrom '%s' in service '%s' does not exist\n", volumesFrom, name)
			}
		}
	}
	return nil
}

//        "build": {
//            "type": "string"
//        },
//        "external_links": {
//            "items": {
//                "type": "string",
//             	 "pattern": "[0-9a-fA-F]+(:[0-9a-fA-F_]+)?]"
//            }
//        },
//        "expose": {
//            "items": {
//                "type": "string",
//             	 "pattern": "[0-9]+"
//            }
//        },
//        "env_file": {
//            "type": "string"
//        },
//        "extends": {
//            "type": "object",
//		    "additionalProperties": false,
//		    "required": [
//		        "file",
//		        "service"
//		    ],
//		    "properties": {
//		        "file": {
//		            "type": "string"
//		        },
//		        "service": {
//		            "type": "string"
//		        }
//			}
//        },
//        "net": {
//            "type": "string"
//        },
//        "dns": {
//            "type": "string"
//        },
//        "cap_add": {
//            "type": "string"
//        },
//        "cap_drop": {
//            "type": "string"
//        },
//        "dns_search": {
//            "type": "string"
//        },
//        "working_dir": {
//            "type": "string"
//        },
//        "user": {
//            "type": "string"
//        },
//        "hostname": {
//            "type": "string"
//        },
//        "domainname": {
//            "type": "string"
//        },
//        "stdin_open": {
//            "type": "boolean"
//        },
//        "tty": {
//            "type": "boolean"
//        },
//        "cpu_shares": {
//            "type": "integer"
//        },
