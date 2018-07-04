package api

import (
	"fmt"
)

type GetMobileEnabledOutput struct {
	Response bool `json:"response"`
}

func (c *Client) GetMobileEnabled(id string) (*GetMobileEnabledOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetMobileEnabled(): begin")
	defer c.logger.Debugf("electricgopher.api.Client.GetMobileEnabled(): end")

	out := &GetMobileEnabledOutput{}
	err := c.doGet(fmt.Sprintf("/api/1/vehicles/%s/mobile_enabled", id), out)
	return out, err
}
