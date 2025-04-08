package models

import (
	"time"

	"github.com/gocql/gocql"
)

// Camera struct
type Camera struct {
	CameraID         gocql.UUID `json:"camera_id"`
	Location         string     `json:"location"`
	CameraType       string     `json:"camera_type"`
	CameraStatus     int        `json:"camera_status"`
	CameraLastUpdate time.Time  `json:"camera_last_update"`
	Discription      string     `json:"discription"`
}
