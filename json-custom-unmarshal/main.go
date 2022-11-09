package main

import (
	"encoding/json"
	"fmt"
)

type ExtendedDbInstanceSrc struct {
	InstanceId           string `json:"instance_id"`
	InstanceIp           string `json:"instance_ip"`
	InstancePort         string `json:"instance_port"`
	InstanceAlias        string `json:"instance_alias"`
	IsInstanceHealth     bool   `json:"is_instance_health"`
	ErrorMsg             string `json:"error_msg"`
	IsInstanceRoleMaster bool   `json:"is_instance_role_master"`
}

type ExtendedDbInstanceTarget struct {
	InstanceId       string `json:"instance_id"`
	InstanceIp       string `json:"instance_ip"`
	InstancePort     string `json:"instance_port"`
	InstanceAlias    string `json:"instance_alias"`
	IsInstanceHealth bool   `json:"is_instance_health"`
	ErrorMsg         string `json:"error_msg"`
	InstanceRole     string `json:"instance_role"` // ["master", "slave", "no_role", ""]
}

func (t *ExtendedDbInstanceTarget) UnmarshalJSON(b []byte) error {
	type Alias ExtendedDbInstanceTarget
	aux := &struct {
		IsInstanceRoleMaster bool `json:"is_instance_role_master"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.IsInstanceRoleMaster {
		t.InstanceRole = "master"
	} else {
		t.InstanceRole = "slave"
	}

	return nil
}

func main() {
	src := &ExtendedDbInstanceSrc{
		InstanceId:       "instance_id",
		InstanceIp:       "instance_ip",
		InstancePort:     "instance_port",
		InstanceAlias:    "instance_alias",
		ErrorMsg:         "error_msg",
		IsInstanceHealth: true,
	}

	b, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}

	tgt := &ExtendedDbInstanceTarget{}

	err = json.Unmarshal(b, tgt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", tgt)
}
