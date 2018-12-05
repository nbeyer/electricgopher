package api

import (
	"bytes"
	"fmt"
)

func (c *Client) WakeUp(id string) (*WakeUpOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.WakeUp(%v): begin", id)
	defer c.logger.Debugf("electricgopher.api.Client.WakeUp%v): end", id)

	out := &WakeUpOutput{}
	err := c.doPost(fmt.Sprintf("/api/1/vehicles/%s/wake_up", id), out, bytes.NewBufferString(""))
	return out, err
}

type WakeUpOutput struct {
	Response Vehicle `json:"response"`
}

func (c *Client) Unlock(id string) (*CommandResponse, error) {
	c.logger.Debugf("electricgopher.api.Client.Unlock(%v): begin", id)
	defer c.logger.Debugf("electricgopher.api.Client.Unlock(%v): end", id)
	out := &ResponseEnvelope{}
	err := c.doPost(fmt.Sprintf("/api/1/vehicles/%s/command/door_unlock", id), out, bytes.NewBufferString(""))
	return &out.Response, err
}

func (c *Client) OpenFrunk(id string) (*CommandResponse, error) {
	c.logger.Debugf("electricgopher.api.Client.OpenFrunk(%v): begin", id)
	defer c.logger.Debugf("electricgopher.api.Client.OpenFrunk(%v): end", id)
	return c.actuateTrunk(id, "front")
}

func (c *Client) OpenTrunk(id string) (*CommandResponse, error) {
	c.logger.Debugf("electricgopher.api.Client.OpenTrunk(%v): begin", id)
	defer c.logger.Debugf("electricgopher.api.Client.OpenTrunk(%v): end", id)
	return c.actuateTrunk(id, "rear")
}

func (c *Client) actuateTrunk(id string, which string) (*CommandResponse, error) {
	c.logger.Debugf("electricgopher.api.Client.OpenFrunk(%v): begin", id)
	defer c.logger.Debugf("electricgopher.api.Client.OpenFrunk(%v): end", id)
	out := &ResponseEnvelope{}
	err := c.doPost(fmt.Sprintf("/api/1/vehicles/%s/command/actuate_trunk", id), out, bytes.NewBufferString(fmt.Sprintf("which_trunk=%s", which)))
	return &out.Response, err
}

type ResponseEnvelope struct {
	Response CommandResponse `json:"response"`
}

type CommandResponse struct {
	Reason string `json:"reason"`
	Result bool   `json:"result"`
}
