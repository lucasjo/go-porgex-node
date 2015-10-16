package models

import "gopkg.in/mgo.v2/bson"

type PendingAppOpGroup struct {
	Pending_ops         []PendingAppOp `json:"pending_ops"`
	Parent_op_id        bson.ObjectId  `json:"parent_op_id" bson:"parent_op_id, omitempty"`
	Num_gears_added     int            `json:"num_gears_added"`
	Num_gears_removed   int            `json:"num_gears_removed"`
	Num_gears_created   int            `json:"num_gears_created"`
	Num_gears_roll_back int            `json:"num_gears_roll_back"`
	User_agent          string         `json:"user_agent"`
	Rollback_blocked    bool           `json:"rollback_blocked"`
}
