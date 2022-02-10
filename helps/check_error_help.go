package helps

import (
	"strings"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

func IsError(err error) bool {
	return err != nil && !strings.Contains(err.Error(), generic.TErrorCode_EGood.String())
}

func IsIncorrectFrameSize(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Incorrect frame size")
}

func IsNotExisted(err error) bool {
	return err != nil && strings.Contains(err.Error(), generic.TErrorCode_EItemNotExisted.String())
}

func IsUnknown(err error) bool {
	return err != nil && strings.Contains(err.Error(), generic.TErrorCode_EUnknownException.String())
}
