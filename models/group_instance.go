package models

type GroupInstance struct {
	Platform    string `json:"platform"`
	Gear_size   string `json:"gear_size"`
	Addtl_fs_gb int    `json:"addtl_fs_gb"`
}
