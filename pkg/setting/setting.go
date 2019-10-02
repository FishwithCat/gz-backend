package setting

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

var (
	DbUser     string
	DbPassword string
	JwtSecret  string
)

func init() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	kvStore := client.KV()
	initConfig(kvStore, "config/database/user", &DbUser)
	initConfig(kvStore, "config/database/password", &DbPassword)
	initConfig(kvStore, "config/backend/jwtSecret", &JwtSecret)
}

func initConfig(kvStore *api.KV, keyPath string, variable *string) {
	pair, _, err := kvStore.Get(keyPath, nil)
	if err != nil {
		panic(err)
	}
	*variable = string(pair.Value)
}
