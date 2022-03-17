package config

import "github.com/spf13/viper"

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) {
	_ = cfg.UnmarshalKey(key, rawVal, opts...)
}

func AllSettings() map[string]interface{} {
	return cfg.AllSettings()
}
