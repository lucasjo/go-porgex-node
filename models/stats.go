package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CPUUsage stores All CPU stats aggregated since container inception.
type CPUUsage struct {
	// Total CPU time consumed.
	// Units: nanoseconds.
	Total_usage uint64 `json:"total_usage"`
	// Total CPU time consumed per core.
	// Units: nanoseconds.
	Percpu_usage []uint64 `json:"percpu_usage"`
	// Time spent by tasks of the cgroup in kernel mode.
	// Units: nanoseconds.
	Usage_in_sytemmode uint64 `json:"usage_in_kernelmode"`
	// Time spent by tasks of the cgroup in user mode.
	// Units: nanoseconds.
	Usage_in_usermode uint64 `json:"usage_in_usermode"`
}

// CPUStats aggregates and wraps all CPU related info of container
type CPUStats struct {
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AppId        string        `json:"appId"`
	Cpu_usage    CPUUsage      `json:"cpu_usage"`
	System_usage uint64        `json:"system_cpu_usage"`
	CreateAt     time.Time     `json:"create_at"`
}

type MemStats struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AppId         string        `json:"appId"`
	Max_usage     uint64        `json:"max_usage"`
	Limit_usage   uint64        `json:"limit_usage"`
	Current_usage uint64        `json:"current_usage"`
	Create_at     time.Time     `json:"create_at"`
}
