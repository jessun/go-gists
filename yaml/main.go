package main

import (
	"encoding/json"
	"fmt"

	"github.com/actiontech/sqle/sqle/config"

	"gopkg.in/yaml.v2"
)

type Cfg struct {
	YamlConfig  config.Config `json:"yaml.config"`
	Test        string        `json:"test"`
	YamlConfig2 config.Config `json:"yaml.Config"`
}

func main() {

	cfg := Cfg{
		YamlConfig: config.Config{
			Server: config.Server{
				SqleCnf: config.SqleConfig{
					SqleServerPort: 10000,
					SecretKey:      "secret_key",
				},
			},
		},
		Test: "test",
	}

	bs, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("=====================")

	bs, err = yaml.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
