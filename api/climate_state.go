package api

import (
	"fmt"
)

type ClimateState struct {
	InsideTemp              float64 `json:"inside_temp"`
	OutsideTemp             float64 `json:"outside_temp"`
	DriverTempSetting       float64 `json:"driver_temp_setting"`
	PassengerTempSetting    float64 `json:"passenger_temp_setting"`
	LeftTempDirection       int64   `json:"left_temp_direction"`
	RightTempDirection      int64   `json:"right_temp_direction"`
	IsFrontDefrosterOn      bool    `json:"is_front_defroster_on"`
	IsRearDefrosterOn       bool    `json:"is_rear_defroster_on"`
	FanStatus               int64   `json:"fan_status"`
	IsClimateOn             bool    `json:"is_climate_on"`
	MinAvailTemp            float64 `json:"min_avail_temp"`
	MaxAvailTemp            float64 `json:"max_avail_temp"`
	SeatHeaterLeft          bool    `json:"seat_heater_left"`
	SeatHeaterRight         bool    `json:"seat_heater_right"`
	SeatHeaterRearLeft      bool    `json:"seat_heater_rear_left"`
	SeatHeaterRearRight     bool    `json:"seat_heater_rear_right"`
	SeatHeaterRearCenter    bool    `json:"seat_heater_rear_center"`
	SeatHeaterRearLeftBack  int64   `json:"seat_heater_rear_left_back"`
	SeatHeaterRearRightBack int64   `json:"seat_heater_rear_right_back"`
	BatteryHeater           bool    `json:"battery_heater"`
	BatteryHeaterNoPower    bool    `json:"battery_heater_no_power"`
	SteeringWheelHeater     bool    `json:"steering_wheel_heater"`
	WipperBladeHeater       bool    `json:"wipper_blade_heater"`
	SideMirrorHeaters       bool    `json:"side_mirror_heaters"`
	IsPreconditioning       bool    `json:"is_preconditioning"`
	SmartPreconditioning    bool    `json:"smart_preconditioning"`
	IsAutoConditioningOn    bool    `json:"is_auto_conditioning_on"`
	Timestamp               int64   `json:"timestamp"`
}

type GetClimateStateOutput struct {
	Response ClimateState `json:"response"`
}

func (c *Client) GetClimateState(id string) (*GetClimateStateOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetClimateState(): begin")
	defer c.logger.Debugf("electricgopher.api.Client.GetClimateState(): end")

	out := &GetClimateStateOutput{}
	err := c.doGet(fmt.Sprintf("/api/1/vehicles/%s/data_request/climate_state", id), out)
	return out, err
}
