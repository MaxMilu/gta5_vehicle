package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	HTTPS  bool `default:"false" env:"HTTPS"`
	Port   uint `default:"7000" env:"PORT"`
	System struct {
		ResourceRoot               string `env:"SystemResourceRoot" default:""`
		MaxThreadQuantity          int    `env:"MaxThreadQuantity" default:"10"`
		MaxConnectionQuantity      int    `env:"MaxConnectionQuantity" default:"10000"`
		MaxInnerConnectionQuantity int    `env:"MaxInnerConnectionQuantity" default:"10000"`
	}
	DB struct {
		Adapter       string `env:"DBAdapter" default:"mysql"`
		Name          string `env:"DBName" default:"qor_example"`
		Host          string `env:"DBHost" default:"localhost"`
		Port          string `env:"DBPort" default:"3306"`
		User          string `env:"DBUser"`
		Password      string `env:"DBPassword"`
		MachineID     uint16 `env:"DBWorkID"`
		HistorySchema string `env:"DBHistorySchema" default:"qor_example_history"`
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/database.example.yml"); err != nil {
		println(err.Error())
		return
	}
}
