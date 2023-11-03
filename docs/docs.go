// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/l3af/configs/v1": {
            "get": {
                "description": "Returns details of the configuration of eBPF Programs for all interfaces on a node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Returns details of the configuration of eBPF Programs for all interfaces on a node",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/l3af/configs/v1/add": {
            "post": {
                "description": "Adds new eBPF Programs on node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adds new eBPF Programs on node",
                "parameters": [
                    {
                        "description": "BPF programs",
                        "name": "cfgs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.L3afBPFPrograms"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/l3af/configs/v1/delete": {
            "post": {
                "description": "Removes eBPF Programs on node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Removes eBPF Programs on node",
                "parameters": [
                    {
                        "description": "BPF program names",
                        "name": "cfgs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.L3afBPFProgramNames"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/l3af/configs/v1/update": {
            "post": {
                "description": "Update eBPF Programs configuration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update eBPF Programs configuration",
                "parameters": [
                    {
                        "description": "BPF programs",
                        "name": "cfgs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.L3afBPFPrograms"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/l3af/configs/v1/{iface}": {
            "get": {
                "description": "Returns details of the configuration of eBPF Programs for a given interface",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Returns details of the configuration of eBPF Programs for a given interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "interface name",
                        "name": "iface",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.BPFProgram": {
            "type": "object",
            "properties": {
                "admin_status": {
                    "description": "Program admin status enabled or disabled",
                    "type": "string"
                },
                "artifact": {
                    "description": "Artifact file name",
                    "type": "string"
                },
                "cfg_version": {
                    "description": "Config version",
                    "type": "integer"
                },
                "cmd_config": {
                    "description": "Program config providing command",
                    "type": "string"
                },
                "cmd_start": {
                    "description": "Program start command",
                    "type": "string"
                },
                "cmd_status": {
                    "description": "Program status command",
                    "type": "string"
                },
                "cmd_stop": {
                    "description": "Program stop command",
                    "type": "string"
                },
                "cmd_update": {
                    "description": "Program update config command",
                    "type": "string"
                },
                "config_args": {
                    "description": "Map of arguments to config command",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.L3afDNFArgs"
                        }
                    ]
                },
                "config_file_path": {
                    "description": "Config file location",
                    "type": "string"
                },
                "cpu": {
                    "description": "User program cpu limits",
                    "type": "integer"
                },
                "ebpf_package_repo_url": {
                    "description": "Download url for Program",
                    "type": "string"
                },
                "entry_function_name": {
                    "description": "BPF entry function name to load",
                    "type": "string"
                },
                "id": {
                    "description": "Program id",
                    "type": "integer"
                },
                "is_plugin": {
                    "description": "User program is plugin or not",
                    "type": "boolean"
                },
                "map_args": {
                    "description": "Config BPF Map of arguments",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.L3afDNFArgs"
                        }
                    ]
                },
                "map_name": {
                    "description": "BPF map to store next program fd",
                    "type": "string"
                },
                "memory": {
                    "description": "User program memory limits",
                    "type": "integer"
                },
                "monitor_maps": {
                    "description": "Metrics BPF maps",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.L3afDNFMetricsMap"
                    }
                },
                "name": {
                    "description": "Name of the BPF program package",
                    "type": "string"
                },
                "object_file": {
                    "description": "Object file contains BPF code",
                    "type": "string"
                },
                "prog_type": {
                    "description": "Program type XDP or TC",
                    "type": "string"
                },
                "rules": {
                    "description": "Config rules",
                    "type": "string"
                },
                "rules_file": {
                    "description": "Config rules file name",
                    "type": "string"
                },
                "seq_id": {
                    "description": "Sequence position in the chain",
                    "type": "integer"
                },
                "start_args": {
                    "description": "Map of arguments to start command",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.L3afDNFArgs"
                        }
                    ]
                },
                "status_args": {
                    "description": "Map of arguments to status command",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.L3afDNFArgs"
                        }
                    ]
                },
                "stop_args": {
                    "description": "Map of arguments to stop command",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.L3afDNFArgs"
                        }
                    ]
                },
                "update_args": {
                    "description": "Map of arguments to update command",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.L3afDNFArgs"
                        }
                    ]
                },
                "user_program_daemon": {
                    "description": "User program daemon or not",
                    "type": "boolean"
                },
                "version": {
                    "description": "Program version",
                    "type": "string"
                }
            }
        },
        "models.BPFProgramNames": {
            "type": "object",
            "properties": {
                "tc_egress": {
                    "description": "names of the TC egress eBPF programs",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tc_ingress": {
                    "description": "names of the TC ingress eBPF programs",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "xdp_ingress": {
                    "description": "names of the XDP ingress eBPF programs",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.BPFPrograms": {
            "type": "object",
            "properties": {
                "tc_egress": {
                    "description": "list of tc egress bpf programs",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BPFProgram"
                    }
                },
                "tc_ingress": {
                    "description": "list of tc ingress bpf programs",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BPFProgram"
                    }
                },
                "xdp_ingress": {
                    "description": "list of xdp ingress bpf programs",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.BPFProgram"
                    }
                }
            }
        },
        "models.L3afBPFProgramNames": {
            "type": "object",
            "properties": {
                "bpf_programs": {
                    "description": "List of eBPF program names to remove",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.BPFProgramNames"
                        }
                    ]
                },
                "host_name": {
                    "description": "Host name or pod name",
                    "type": "string"
                },
                "iface": {
                    "description": "Interface name",
                    "type": "string"
                }
            }
        },
        "models.L3afBPFPrograms": {
            "type": "object",
            "properties": {
                "bpf_programs": {
                    "description": "List of bpf programs",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.BPFPrograms"
                        }
                    ]
                },
                "host_name": {
                    "description": "Host name or pod name",
                    "type": "string"
                },
                "iface": {
                    "description": "Interface name",
                    "type": "string"
                }
            }
        },
        "models.L3afDNFArgs": {
            "type": "object",
            "additionalProperties": true
        },
        "models.L3afDNFMetricsMap": {
            "type": "object",
            "properties": {
                "aggregator": {
                    "description": "Aggregation function names",
                    "type": "string"
                },
                "key": {
                    "description": "Index of the bpf map",
                    "type": "integer"
                },
                "name": {
                    "description": "BPF map name",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "L3AFD APIs",
	Description:      "Configuration APIs to deploy and get the details of the eBPF Programs on the node",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
