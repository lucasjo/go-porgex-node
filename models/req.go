package models

import "encoding/json"

type Request struct {
	Service  string          `json:"service"`
	Fromhost string          `json:"fromhost"`
	Data     json.RawMessage `json:"data"`
}
