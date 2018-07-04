package api

import (
	"fmt"
)

type ChargeState struct {
	ChargingState               string  `json:"charging_state"`
	FastChargerType             string  `json:"fast_charger_type"`
	FastChargerBrand            string  `json:"fast_charger_brand"`
	ChargeLimitSoc              int64   `json:"charge_limit_soc"`
	ChargeLimitSocStd           int64   `json:"charge_limit_soc_std"`
	ChargeLimitSocMin           int64   `json:"charge_limit_soc_min"`
	ChargeLimitSocMax           int64   `json:"charge_limit_soc_max"`
	ChargeToMaxRange            bool    `json:"charge_to_max_range"`
	MaxRangeChargeCounter       int64   `json:"max_range_charge_counter"`
	FastChargerPresent          bool    `json:"fast_charger_present"`
	BatteryRange                float64 `json:"battery_range"`
	EstBatteryRange             float64 `json:"est_battery_range"`
	IdealBatteryRange           float64 `json:"ideal_battery_range"`
	BatteryLevel                int64   `json:"battery_level"`
	UsableBatteryLevel          int64   `json:"usable_battery_level"`
	ChargeEnergyAdded           float64 `json:"charge_energy_added"`
	ChargeMilesAddedRated       float64 `json:"charge_miles_added_rated"`
	ChargeMilesAddedIdeal       float64 `json:"charge_miles_added_ideal"`
	ChargerVoltage              int64   `json:"charger_voltage"`
	ChargerPilotCurrent         int64   `json:"charger_pilot_current"`
	ChargerActualCurrent        int64   `json:"charger_actual_current"`
	ChargerPower                int64   `json:"charger_power"`
	TimeToFullCharge            float64 `json:"time_to_full_charge"`
	TripCharging                bool    `json:"trip_charging"`
	ChargeRate                  float64 `json:"charge_rate"`
	ChargePortDoorOpen          bool    `json:"charge_port_door_open"`
	ConnChargeCable             string  `json:"conn_charge_cable"`
	ScheduledChargingStartTime  string  `json:"scheduled_charging_start_time"`
	ScheduledChargingPending    bool    `json:"scheduled_charging_pending"`
	UserChargeEnableRequest     string  `json:"user_charge_enable_request"`
	ChargeEnableRequest         bool    `json:"charge_enable_request"`
	ChargerPhases               string  `json:"charger_phases"`
	ChargePortLatch             string  `json:"charge_port_latch"`
	ChargeCurrentRequest        int64   `json:"charge_current_request"`
	ChargeCurrentRequestMax     int64   `json:"charge_current_request_max"`
	ManagedChargingActive       bool    `json:"managed_charging_active"`
	ManagedChargingUserCanceled bool    `json:"managed_charging_user_canceled"`
	ManagedChargingStartTime    string  `json:"managed_charging_start_time"`
	BatteryHeaterOn             bool    `json:"battery_heater_on"`
	NotEnoughPowerToHeat        bool    `json:"not_enough_power_to_heat"`
	Timestamp                   int64   `json:"timestamp"`
}

type GetChargeStateOutput struct {
	Response ChargeState `json:"response"`
}

func (c *Client) GetChargeState(id string) (*GetChargeStateOutput, error) {
	c.logger.Debugf("electricgopher.api.Client.GetChargeState(): begin")
	defer c.logger.Debugf("electricgopher.api.Client.GetChargeState(): end")

	out := &GetChargeStateOutput{}
	err := c.doGet(fmt.Sprintf("/api/1/vehicles/%s/data_request/charge_state", id), out)
	return out, err
}
