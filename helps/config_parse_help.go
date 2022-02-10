package helps

import (
	"flag"
	"fmt"
	"worker-transaction/consts"
)

func ParseConfigPath(env string) string {
	if env == consts.ENV_DEV_K8S {
		return fmt.Sprintf("%s/conf/app_dev_k8s.conf", GetCurrentPath())
	} else if env == consts.ENV_PROD_K8S {
		return fmt.Sprintf("%s/conf/app_prod_k8s.conf", GetCurrentPath())
	} else {
		return consts.PATH_CONFIG_FILE_LOCAL_DEV
	}
}

func ReadFlags(env *string) {
	flag.StringVar(env, "env", "", "environment deploy [k8s_dev, k8s_prod, local dev]")
	flag.Parse()
}
