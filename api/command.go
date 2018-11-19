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

func (c *Client) Unlock(id string) (*UnlockOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.Unlock(%v): begin", id)
	defer c.logger.Debugf("electricgopher.api.Client.Unlock(%v): end", id)
	out := &UnlockOutput{}
	err := c.doPost(fmt.Sprintf("/api/1/vehicles/%s/command/door_unlock", id), out, bytes.NewBufferString(""))
	return out, err
}

type UnlockOutput struct {
	Response UnlockPayload `json:"response"`
}

type UnlockPayload struct {
	Reason string `json:"reason"`
	Result bool   `json:"result"`
}
