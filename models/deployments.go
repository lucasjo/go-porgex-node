package models

import "time"

type Deployment struct {
	Deployment_id     string    `json:"deployment_id"`
	Create_at         time.Time `json:"create_at"`
	Hot_deploy        bool      `json:"hot_deploy"`
	Force_clean_build bool      `json:"force_clean_build"`
	Ref               string    `json:"ref"`
	Sha1              string    `json:"sha1"`
	Artifact_url      string    `json:"artifact_url"`
	Activations       []int     `json:"activations"`
}
