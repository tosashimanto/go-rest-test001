package service

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var (
	APP_NAME = `Gsurvey-dummy`

	APP_VERSION                   = "1.0.0"
	DEV_ENV                       = "dev"
	STAGING_ENV                   = "staging"
	PROD_ENV                      = "prod"
	REPLACER    *strings.Replacer = strings.NewReplacer(".", "_")

	APP_CONFIG_PREFIX = `APP`
	APP_CONFIG_NAME   = `application`
	DB_CONFIG_PREFIX  = `DB`
	DB_CONFIG_NAME    = `database`

	CONFIG_PATH      = `config`
	CONFIG_FILE_TYPE = `yaml`
)

var BUCKETEER_AWS_ACCESS_KEY_ID string
var BUCKETEER_AWS_REGION string
var BUCKETEER_AWS_SECRET_ACCESS_KEY string
var BUCKETEER_BUCKET_NAME string

func SetConfig() {

	//viper.SetConfigName("controller")    // no need to include file extension
	//viper.AddConfigPath("config") // set the path of your config file
	var (
		appConfig *viper.Viper
		err       error
	)
	appConfig = viper.New()
	appConfig.SetEnvPrefix(APP_CONFIG_PREFIX)
	appConfig.AutomaticEnv()
	appConfig.SetConfigName(APP_CONFIG_NAME)
	appConfig.SetConfigType(CONFIG_FILE_TYPE)
	appConfig.AddConfigPath(CONFIG_PATH)
	err = appConfig.ReadInConfig()

	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		key := appConfig.GetString("BUCKETEER_AWS_ACCESS_KEY_ID")
		BUCKETEER_AWS_ACCESS_KEY_ID = key
		fmt.Println("key=" + key)

		BUCKETEER_AWS_REGION = appConfig.GetString("BUCKETEER_AWS_REGION")
		BUCKETEER_AWS_SECRET_ACCESS_KEY = appConfig.GetString("BUCKETEER_AWS_SECRET_ACCESS_KEY")
		BUCKETEER_BUCKET_NAME = appConfig.GetString("BUCKETEER_BUCKET_NAME")
	}

}
