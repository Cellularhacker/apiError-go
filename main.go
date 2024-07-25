package apiError

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	UnknownServerError                 = 5000
	Unauthorized                       = 5001
	EmailAlreadyRegistered             = 5002
	PasswordLessThanRequired           = 5003
	SomethingMissing                   = 5004
	WrongConfirmationCode              = 5005
	WrongAccessToken                   = 5006
	EmailNotRegistered                 = 5007
	EmailAlreadyConfirmed              = 5008
	WrongPassword                      = 5009
	DuplicateAPIKey                    = 5010
	WrongEmailOrPassword               = 5011
	NameAlreadyRegistered              = 5012
	UserDeleted                        = 5013
	ContentMissing                     = 5014
	NoLeaguePresent                    = 5015
	CantUnfollowLOT                    = 5016
	ReferralCodeMissing                = 5017
	EmailBelongsToAnotherAccount       = 5018
	PhoneNumberBelongsToAnotherAccount = 5019

	BotBalanceLessThanMinimum = 5031
	BotKeyCantDelete          = 5032
	BotCompanyError           = 5033
	BotDoesntExistError       = 5034

	FunctionNotSupportedYetForGivenExchange = 5035
	InvoicePaymentVerifyUnknownError        = 5036
	KeyHasNoPermissionForGivenOperation     = 5037

	LeagueRegisterByPassed      = 5040
	LeagueMinimumBalanceNotMet  = 5041
	LeagueRemovedDueToKeyDelete = 5042

	ErrorBotRatioDifferenceHigh      = 6001
	ErrorBotCopyingKeyInvalidated    = 6002
	ErrorBotOriginalKeyInvalidated   = 6003
	ErrorBotPositionMismatch         = 6004
	ErrorBotUnknown                  = 6005
	ErrorBotOpenOrderFound           = 6006
	ErrorKeyIsRegisteredInAnotherBot = 6007
	ErrorKeyIsNotCopyLeader          = 6008
	ErrorCopyLeaderConditionsNotMet  = 6009
	ErrorBotBalanceNotSufficient     = 6010
)

// Error APIError error returned by APIHandler
type Error struct {
	Error   error
	Message string
	Code    int
}

func New(message string, code int) *Error {
	return &Error{nil, message, code}
}

func DetectError(err error) *Error {
	errStr := strings.ToLower(err.Error())
	if strings.Contains(errStr, "invalid") {
		return &Error{nil, errStr, http.StatusBadRequest}
	} else if strings.Contains(errStr, "author") {
		return NotAuthorizedError(errStr)
	} else if strings.Contains(errStr, "found") {
		return NotFoundError(errStr)
	} else if strings.Contains(errStr, "legal") {
		return UnavailableLegalReason(errStr)
	}

	return InternalServerError(err)
}

func InternalServerErrorLoc(loc string, err error) *Error {
	return InternalServerError(fmt.Errorf("%s: %s", loc, err))
}

func InternalServerError(err error) *Error {
	return &Error{err, "Unknown error! Try again", UnknownServerError}
}

func BadRequestError(location string) *Error {
	return &Error{nil, fmt.Sprintf("'%s' is invalid or missing", location), http.StatusBadRequest}
}

func BadRequestMsgErr(msg string) *Error {
	return &Error{nil, msg, http.StatusBadRequest}
}

func NotAuthorizedError(message string) *Error {
	return &Error{nil, message, http.StatusUnauthorized}
}
func NotAuthorizedUser() *Error {
	return NotAuthorizedError("unauthorized")
}

func SomethingMissingError(message string) *Error {
	return &Error{nil, message, SomethingMissing}
}

func NotFoundError(message string) *Error {
	return &Error{nil, message, http.StatusNotFound}
}

func UnavailableLegalReason(message string) *Error {
	return &Error{nil, message, http.StatusUnavailableForLegalReasons}
}

func ConflictError(message string) *Error {
	return &Error{nil, message, http.StatusConflict}
}
