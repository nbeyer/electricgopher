package api

import (
	"fmt"
)

type SpeedLimitMode struct {
	Active          bool    `json:"active"`
	CurrentLimitMph float64 `json:"current_limit_mph"`
	MaxLimitMph     int64   `json:"max_limit_mph"`
	MinLimitMph     int64   `json:"min_limit_mph"`
	PinCodeSet      bool    `json:"pin_code_set"`
}

type VehicleState struct {
	ApiVersion              int            `json:"api_version"`
	AutoparkStateV2         string         `json:"autopark_state_v2"`
	AutoparkStyle           string         `json:"autopark_style"`
	CalendarSupported       bool           `json:"calendar_supported"`
	CarVersion              string         `json:"car_verson"`
	CenterDisplayState      int            `json:"center_display_state"`
	Df                      int64          `json:"df"`
	Dr                      int64          `json:"dr"`
	Ft                      int64          `json:"ft"`
	HomelinkNearby          bool           `json:"homelink_nearby"`
	LastAutoparkError       string         `json:"last_autopark_error"`
	Locked                  bool           `json:"locked"`
	NotificationsSupported  bool           `json:"notifications_supported"`
	Odometer                float64        `json:"odometer"`
	ParsedCalendarSupported bool           `json:"parsed_calendar_supported"`
	Pf                      int64          `json:"pf"`
	Pr                      int64          `json:"pr"`
	RemoteStart             bool           `json:"remote_start"`
	RemoteStartSupported    bool           `json:"remote_start_supported"`
	Rt                      int64          `json:"rt"`
	SpeedLimitMode          SpeedLimitMode `json:"speed_limit_mode"`
	SunRoofPercentOpen      int            `json:"sun_roof_percent_open"`
	SunRoofState            string         `json:"sun_roof_state"`
	Timestamp               int64          `json:"timestamp"`
	ValetMode               bool           `json:"valet_mode"`
	ValetPinNeeded          bool           `json:"valet_pin_needed"`
	VehicleName             string         `json:"vehicle_name"`
}

type GetVehicleStateOutput struct {
	Response VehicleState `json:"response"`
}

func (c *Client) GetVehicleState(id string) (*GetVehicleStateOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetVehicleState(): begin")
	defer c.logger.Debugf("electricgopher.api.Client.GetVehicleState(): end")

	out := &GetVehicleStateOutput{}
	err := c.doGet(fmt.Sprintf("/api/1/vehicles/%s/data_request/vehicle_state", id), out)
	return out, err
}
