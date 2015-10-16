package models

import

//3rd party package
(
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	AnalyticsProperties struct {
		User_agent string `json:"user_agent"`
	}
	ConfigHash     struct{}
	GroupOverride  struct{}
	MetaProperties struct{}

	Application struct {
		ID                    bson.ObjectId       `json:"id" bson:"_id,omitempty"`
		Analytics             AnalyticsProperties `json:"analytics"`
		Builder_id            string              `json:"builder_id"`
		Canonical_name        string              `json:"canonica_name"`
		Component_instances   []ComponentInstance `json:"component_instances"`
		Config                ConfigHash          `json:"config"`
		Create_at             time.Time           `json:"create_at"`
		Default_gear_size     string              `json:"default_gear_size"`
		Deployments           Deployment          `json:"deployments"`
		Domain_id             bson.ObjectId       `json:"domain_id" bson:"domain_id,omitempty"`
		Domain_namespace      string              `json:"domain_namespace"`
		Gears                 []Gear              `json:"gears"`
		Group_instances       []GroupInstance     `json:"group_instances"`
		Group_overrides       []GroupOverride     `json:"group_override"`
		Ha                    bool                `json:"ha"`
		Init_git_url          string              `json:"init_git_url"`
		Members               []Member            `json:members"`
		Name                  string              `json:"name"`
		Owner_id              bson.ObjectId       `json:"owner_id" bson:"owner_id,omitempty"`
		Pending_app_op_groups []PendingAppOpGroup `json:"pending_op_groups"`
		Scalable              bool                `json:"scalable"`
		Secret_token          string              `json:"secret_token"`
		Update_at             time.Time           `json:"update_at"`
		Meta                  MetaProperties      `json:"meta"`
	}
)
