package main

import (
	"encoding/json"
	"fmt"
)

type ExtendedDbInstance struct {
	InstanceId           string `json:"instance_id"`
	InstanceIp           string `json:"instance_ip"`
	InstancePort         string `json:"instance_port"`
	InstanceAlias        string `json:"instance_alias"`
	IsInstanceHealth     bool   `json:"is_instance_health"`
	IsInstanceRoleMaster bool   `json:"is_instance_role_master,omitempty"`
	ErrorMsg             string `json:"error_msg"`
}

func main() {
	s := `
{
    "instance_id": "i-1234567890",
    "instance_ip": "127.0.0.1",
    "instance_port": "3306",
    "instance_alias": "test",
    "is_instance_health": true,
    "error_msg": "no"
}
`

	out := &ExtendedDbInstance{}

	err := json.Unmarshal([]byte(s), out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", out)

}
