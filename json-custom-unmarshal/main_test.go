package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtendedDBInstanceUnmarshal(t *testing.T) {
	mockResp := `
{
    "instance_id": "mysql-1",
    "instance_ip": "10.186.65.55",
    "instance_port": "3306",
    "instance_alias": "mysql-1-alias",
    "is_instance_health": true,
    "instance_role": "master",
    "error_msg": "ok"
}
`

	expectedDbInstance := &ExtendedDbInstance{
		InstanceId:       "mysql-1",
		InstanceIp:       "10.186.65.55",
		InstancePort:     "3306",
		InstanceAlias:    "mysql-1-alias",
		IsInstanceHealth: true,
		InstanceRole:     "master",
		ErrorMsg:         "ok",
	}

	actualDbInstance := &ExtendedDbInstance{}
	err := json.Unmarshal([]byte(mockResp), actualDbInstance)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedDbInstance, actualDbInstance)
}

func TestExtendedDBInstancePreVerUnmarshal(t *testing.T) {
	mockResp := `
	{
	    "instance_id": "mysql-2",
	    "instance_ip": "10.186.65.56",
	    "instance_port": "3307",
	    "instance_alias": "mysql-2-alias",
	    "is_instance_health": true,
	    "is_instance_role_master": true,
	    "error_msg": "fake_err"
	}
	`

	expectedDbInstance := &ExtendedDbInstance{
		InstanceId:       "mysql-2",
		InstanceIp:       "10.186.65.56",
		InstancePort:     "3307",
		InstanceAlias:    "mysql-2-alias",
		IsInstanceHealth: true,
		InstanceRole:     "master",
		ErrorMsg:         "fake_err",
	}

	actualDbInstance := &ExtendedDbInstance{}
	err := json.Unmarshal([]byte(mockResp), actualDbInstance)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedDbInstance, actualDbInstance)
}
