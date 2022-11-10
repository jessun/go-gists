package main

import (
	"encoding/json"
)

type ExtendedDbInstance struct {
	InstanceId       string `json:"instance_id"`
	InstanceIp       string `json:"instance_ip"`
	InstancePort     string `json:"instance_port"`
	InstanceAlias    string `json:"instance_alias"`
	IsInstanceHealth bool   `json:"is_instance_health"`
	InstanceRole     string `json:"instance_role"`
	ErrorMsg         string `json:"error_msg"`
}

// UnmarshalJSON used to be compatible with role judgment based on "is_instance_role_master" field. (DMP-14327)
func (t *ExtendedDbInstance) UnmarshalJSON(b []byte) error {
	type Alias ExtendedDbInstance
	aux := &struct {
		IsInstanceRoleMaster *bool `json:"is_instance_role_master"` // WARN: use InstanceRole instead
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.IsInstanceRoleMaster == nil {
		return nil
	}

	if *aux.IsInstanceRoleMaster {
		t.InstanceRole = "master"
		return nil
	}

	t.InstanceRole = "slave"

	return nil
}
