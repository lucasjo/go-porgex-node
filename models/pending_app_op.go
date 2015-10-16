package models

import "gopkg.in/mgo.v2/bson"

type PendingAppOp struct {
	State             string        `json:"state"`
	Prereq            []string      `json:"prereq"`
	Retry_count       int           `json:"retry_count"`
	Retry_rollback_op bson.ObjectId `json:"retry_rollback_op" bson:"retry_rollback_op,omitempty"`
	Skip_rollback     bool          `json:"skip_rollback"`
}
