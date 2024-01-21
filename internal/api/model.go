package api

import (
	"time"

	"sentinel/internal/dataset"
)

type lastValuesReq struct {
	ID string `json:"id"` // unique id of device
}

type saveValuesReq struct {
	ID     string  `json:"id"`     // unique ID of device
	Temp   float32 `json:"temp"`   // temperature sensor data
	Light  int     `json:"light"`  // light sensor data
	Motion bool    `json:"motion"` // motion sensor data
}

func (r saveValuesReq) toModel(stamp time.Time) dataset.Dataset {
	return dataset.Dataset{
		ID:        r.ID,
		Temp:      r.Temp,
		Light:     r.Light,
		Motion:    r.Motion,
		UpdatedAt: stamp,
	}
}
