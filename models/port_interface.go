package models

type PortInterface struct {
	Cartridge_name   string   `json:"cartridge_name"`
	External_port    string   `json:"external_port"`
	Internal_address string   `json:"internal_address"`
	Protocols        []string `json:"proptocols"`
	Type             []string `json:"type"`
	Mappings         []string `json:"mappings"`
}
