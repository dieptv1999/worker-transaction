package responseservice

import (
	"worker-transaction/restapi/responses"
)

func GetCommonErrorResponseArray(err error) responses.ResponseCommonArray {
	return responses.ResponseCommonArray{
		Data:       make([]interface{}, 0),
		TotalCount: 0,
		Error:      responses.NewErr(err),
	}
}

func GetCommonErrorResponse(err error) responses.ResponseCommonSingle {
	return responses.ResponseCommonSingle{
		Data:  map[string]interface{}{},
		Error: responses.NewErr(err),
	}
}

// func parseErrors(errs validation.Errors) map[string]string {
// 	errors := map[string]string{}
// 	bData, err := errs.MarshalJSON()
// 	if err != nil {
// 		log.Println(err.Error(), "-- err.Error() services/responseservice/common_response.go:30")
// 		return map[string]string{}
// 	}

// 	err = json.Unmarshal(bData, &errors)
// 	if err != nil {
// 		log.Println(err.Error(), "-- err.Error() services/responseservice/common_response.go:36")
// 	}

// 	return errors
// }

// func GetCommonErrorResponseWithValidate(err validation.Errors) responses.ResponseCommonSingleWithValidate {
// 	return responses.ResponseCommonSingleWithValidate{
// 		Data: map[string]interface{}{},
// 		Error: &responses.ValidateErr{
// 			Code:    401,
// 			Message: parseErrors(err),
// 		},
// 	}
// }

func GetCommonSucceedResponseArray(data interface{}, totalCount int64) responses.ResponseCommonArray {
	return responses.ResponseCommonArray{
		Data:       data,
		TotalCount: totalCount,
		Error:      responses.NewErr(responses.Success),
	}
}

func GetCommonSucceedResponseWithData(data interface{}) responses.ResponseCommonSingle {
	return responses.ResponseCommonSingle{
		Data:  data,
		Error: responses.NewErr(responses.Success),
	}
}

func GetCommonSucceedResponse() responses.ResponseCommonSingle {
	return responses.ResponseCommonSingle{
		Data: map[string]interface{}{
			"message": "Succeed",
		},
		Error: responses.NewErr(responses.Success),
	}
}
