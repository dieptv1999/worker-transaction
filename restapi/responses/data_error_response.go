package responses

import (
	"errors"
	"net/http"
)

var (
	Success               = errors.New("Success")
	UnSuccess             = errors.New("UnSuccess")
	ErrUnknown            = errors.New("Unknown")
	BadRequest            = errors.New("Bad Request")
	ErrNotInt             = errors.New("Err Int")
	SessionExpired        = errors.New("Session Expired")
	VerifyCodeExpired     = errors.New("Verify Code Expired")
	PhoneUnVerified       = errors.New("Phone Unverified")
	Existed               = errors.New("Existed")
	NotAdmin              = errors.New("Not Admin")
	NotPermission         = errors.New("Not Permission")
	UnAuthorized          = errors.New("UnAuthorized")
	StatusTooManyRequests = errors.New("StatusTooManyRequests")
	ErrLogin              = errors.New("Login Error")
	ErrSystem             = errors.New("System Error")
	NotExisted            = errors.New("Data Not Existed !")
	ProgramExpired        = errors.New("Program expired")
	ProgramNotStartYet    = errors.New("Program not start yet")
	ProgramInsufficient   = errors.New("Remain amount drop is insufficient")
	ErrChangePass         = errors.New("Change Password Error !")
	NotMore               = errors.New("No More Data")
	ExistedPosUser        = errors.New("Existed Pos User")
	LimitCharacter        = errors.New("Limit character")
	CannotEmpty           = errors.New("Cannot empty")
	ErrSig                = errors.New("Sig error")
	ExchangeStop          = errors.New("Exchange stop")
	TurnExceeded          = errors.New("Exceeded")
	InvalidMethod         = errors.New("Invalid Method")
	OverTime              = errors.New("Time is over")
	MethodNotAllowed      = errors.New("Method Not Allowed")
	NotAcceptable         = errors.New("Not Acceptable")
	NotAllowed            = errors.New("Not Allowed")
	UserNameExisted       = errors.New("Username existed")

	MapDescription = map[error]string{
		Success:               "Success!",
		ErrUnknown:            "Unknown error!",
		UnAuthorized:          "UnAuthorized",
		StatusTooManyRequests: "StatusTooManyRequests",
		ProgramExpired:        ProgramExpired.Error(),
		ProgramNotStartYet:    ProgramNotStartYet.Error(),
		ProgramInsufficient:   ProgramInsufficient.Error(),
		PhoneUnVerified:       PhoneUnVerified.Error(),
		BadRequest:            "Bad Request!",
		ErrNotInt:             "Param not int!",
		UnSuccess:             "Unsuccess!",
		VerifyCodeExpired:     "Verify Code Expired",
		SessionExpired:        "SessionExpired!",
		Existed:               "Existed !",
		NotAdmin:              "Not Admin !",
		NotPermission:         "Not Permission !",
		ErrLogin:              "Wrong sign. ",
		ErrSystem:             "The system is having problems.",
		NotExisted:            "Data Not Existed!",
		ErrChangePass:         "Wrong username, password.",
		NotMore:               "No more data",
		ExistedPosUser:        "Existed position organization",
		ErrSig:                "Sig error !",
		CannotEmpty:           "Data empty !",
		ExchangeStop:          "The exchange is suspended!",
		TurnExceeded:          "You have no more turn",
		InvalidMethod:         "Invalid Method",
		OverTime:              "Time is over",
		MethodNotAllowed:      "Method Not Allowed",
		NotAcceptable:         "Not Acceptable",
		NotAllowed:            "Not Allowed",
		UserNameExisted:       "Username Existed",
	}
	MapErrorCode = map[error]int64{
		Success:               200,
		UnSuccess:             201,
		ErrNotInt:             302,
		SessionExpired:        303,
		NotExisted:            304,
		Existed:               408,
		ErrChangePass:         306,
		NotAdmin:              307,
		NotPermission:         308,
		NotMore:               309,
		ExistedPosUser:        310,
		CannotEmpty:           311,
		LimitCharacter:        312,
		ErrSig:                316,
		BadRequest:            400,
		ErrUnknown:            401,
		ErrLogin:              402,
		ErrSystem:             403,
		UnAuthorized:          405,
		StatusTooManyRequests: http.StatusTooManyRequests,
		VerifyCodeExpired:     406,
		ExchangeStop:          413,
		TurnExceeded:          414,
		ProgramExpired:        407,
		ProgramInsufficient:   407,
		ProgramNotStartYet:    407,
		PhoneUnVerified:       407,
		InvalidMethod:         413,
		OverTime:              414,
		MethodNotAllowed:      415,
		NotAcceptable:         416,
		NotAllowed:            417,
		UserNameExisted:       418,
	}
)

// Returns a error.
// swagger:response Err
type Err struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message string `json:"message"`
}

type ValidateErr struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message map[string]string `json:"message"`
}

func NewErr(err error) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: MapDescription[err],
	}
}

func NewErrByText(err error, text string) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: text,
	}
}
