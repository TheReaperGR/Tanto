package config

import "github.com/spf13/viper"

type Config struct {
	Proxmox      ProxmoxConfig  `mapstructure:"proxmox"`
	Bridges      []BridgeConfig `mapstructure:"bridges"`
	VMs          []VMConfig     `mapstructure:"vms"`
	Containerlab ClabConfig     `mapstructure:"containerlab"`
}

// …and the sub-structs: ProxmoxConfig, BridgeConfig, VMConfig, InterfaceConfig, ClabConfig…

// Load reads a YAML file into Config
func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
