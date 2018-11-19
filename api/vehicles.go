package api

import ()

type Vehicle struct {
	Id              int64    `json:"id"`
	IdS             string   `json:"id_s"`
	VehicleId       int64    `json:"vehicle_id"`
	Vin             string   `json:"vin"`
	DisplayName     string   `json:"display_name"`
	OptionCodes     string   `json:"option_codes"`
	Color           string   `json:"color"`
	Tokens          []string `json:"tokens"`
	State           string   `json:"state"`
	InService       bool     `json:"in_service"`
	CalendarEnabled bool     `json:"calendar_enabled"`
}

type GetVehiclesOutput struct {
	Response []Vehicle `json:"response"`
	Count    int       `json:"count"`
}

func (c *Client) GetVehicles() (*GetVehiclesOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetVehicles(): begin")
	defer c.logger.Debugf("electricgopher.api.Client.GetVehicles(): end")

	out := &GetVehiclesOutput{}
	err := c.doGet("/api/1/vehicles", out)
	return out, err
}
