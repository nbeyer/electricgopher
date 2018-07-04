package api

import (
	"fmt"
)

type DriveState struct {
	ShiftState              string  `json:"shift_state"`
	Speed                   int64   `json:"speed"`
	Power                   int64   `json:"power"`
	Latitude                float64 `json:"latitude"`
	Longitude               float64 `json:"longitude"`
	Heading                 int64   `json:"heading"`
	GpsAsOf                 int64   `json:"gps_as_of"`
	NativeLocationSupported int64   `json:"native_location_supported"`
	NativeLatitude          float64 `json:"native_latitude"`
	NativeLongitude         float64 `json:"native_longitude"`
	NativeType              string  `json:"native_type"`
	Timestamp               int64   `json:"timestamp"`
}

type GetDriveStateOutput struct {
	Response DriveState `json:"response"`
}

func (c *Client) GetDriveState(id string) (*GetDriveStateOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetDriveState(): begin")
	defer c.logger.Debugf("electricgopher.api.Client.GetDriveState(): end")

	out := &GetDriveStateOutput{}
	err := c.doGet(fmt.Sprintf("/api/1/vehicles/%s/data_request/drive_state", id), out)
	return out, err
}
