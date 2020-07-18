package chassis

import "c6x.io/chassis/config"

const (
	//EnvPgTestDataFile import test data sql file
	EnvPgTestDataFile = "PG_TEST_DATA_FILE"
	//AppRunEnvLocal app run env "local"
	AppRunEnvLocal = "local"
	//AppRunEnvTest app run env "test"
	AppRunEnvTest = "test"
	//AppRunEnvProd app run env "prod"
	AppRunEnvProd = "prod"
)

//EnvIsProd   check env is production ? local and test will drop database & import test data
func EnvIsProd() bool {
	return !(config.App().Env == "test" || config.App().Env == "local")
}
