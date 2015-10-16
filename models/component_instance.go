package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	ComponentPropertie struct{}

	ComponentInstance struct {
		ID                   bson.ObjectId      `json:"id" bson:"_id,omitempty"`
		Component_properties ComponentPropertie `json:"component_properties"`
		Cartridge_name       string             `json:"cartridge_name"`

		Component_name    string        `json:"component_name"`
		Cartridge_vender  string        `json:"cartridge_vender"`
		Cartridge_id      bson.ObjectId `json:"cartridge_id" bson:"cartridge_id,omitempty"`
		Group_instance_id bson.ObjectId `json:"group_instance_id" bson:"group_instance_id,omitempty"`
		Create_at         time.Time     `json:"create_at"`
	}
)
