package models

import (
	"time"
	//3rd party package
	"gopkg.in/mgo.v2/bson"
)

type (
	UserCapabilities struct {
		Ha                                   bool     `json:"ha"`
		Subaccounts                          bool     `json:"subaccount"`
		Gear_sizes                           []string `json:"gear_sizes"`
		Max_domins                           int      `json:"max_domains"`
		Max_gears                            int      `json:"max_gears"`
		Max_teams                            int      `json:"max_teams"`
		View_global_team                     bool     `json:"view_global_team"`
		Max_untracked_addtl_storage_per_gear int      `json:"max_untracked_addtl_storage_per_gear"`
		Max_tracked_addtl_storage_per_gear   int      `json:"max_tracked_addtl_storage_per_gear"`
	}

	UserSshkey struct {
		ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Type    string        `json:"type"`
		_Type   string        `json:"_type"`
		Name    string        `json:"name"`
		Content string        `json:"content"`
	}

	UserPlanHistory    struct{}
	UserPendingOpGroup struct{}

	CloudUsers struct {
		ID                bson.ObjectId        `json:"id" bson:"_id,omitempty"`
		Capabilities      UserCapabilities     `json:"capabilities"`
		Consumed_gears    int                  `json:"consumed_gears"`
		Created_at        time.Time            `json:"create_at"`
		Login             string               `json:"login" bson:"login"`
		Pending_op_groups []UserPendingOpGroup `json:"pending_op_groups"`
		Plan_history      []UserPlanHistory    `json:"plan_history"`
		Ssh_keys          []UserSshkey         `json:"ssh_keys"`
		Update_at         time.Time            `json:"update_at"`
	}
)
