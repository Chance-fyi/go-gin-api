package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
	"strings"
)

var cfg *viper.Viper

func Init() {
	dirname := "config"
	cfg = viper.New()

	viper.AddConfigPath(dirname)

	for _, cf := range readConfigFile(dirname) {
		viper.SetConfigType(cf["ext"])
		viper.SetConfigName(cf["name"])
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		s := viper.AllSettings()
		if len(s) > 0 {
			cfg.Set(cf["name"], s)
		}
	}
}

func readConfigFile(dirname string) (configFile []map[string]string) {
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range dir {
		fileName := fileInfo.Name()
		ext := path.Ext(fileName)

		configFile = append(configFile, map[string]string{
			"name": fileName[0 : len(fileName)-len(ext)],
			"ext":  strings.Trim(ext, "."),
		})
	}
	return
}

func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) {
	_ = cfg.UnmarshalKey(key, rawVal, opts...)
}
