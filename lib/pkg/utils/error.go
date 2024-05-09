package utils

import "gateway-grpc/lib/helper"

func errorType(err error) (int, error) {
	return helper.CommonError(err)
}
