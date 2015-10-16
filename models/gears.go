package models

import "gopkg.in/mgo.v2/bson"

type Gear struct {
	Server_identity  string        `json:"server_identity"`
	Uuid             string        `json:"uuid"`
	Uid              int           `json:"uid"`
	Name             string        `json:"name"`
	Quarntined       bool          `json:"quarntined"`
	Removed          bool          `json:"removed"`
	Host_singletons  bool          `json:"host_singletons"`
	App_dns          bool          `json:"app_dns"`
	Sparse_carts     []string      `json:"sparse_carts"`
	Group_instanceId bson.ObjectId `json:"group_instance_id" bson:"group_instance_id,omitempty"`
	Port_interfaces  PortInterface `json:"port_interface"`
}
