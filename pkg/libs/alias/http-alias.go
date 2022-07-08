package alias

import (
	"time"
)

const (
	HTTP_OK                   = 200
	HTTP_CREATED              = 201
	HTTP_NOCONTENT            = 204
	HTTP_RESETCONTENT         = 205
	HTTP_PARTIALCONTENT       = 206
	HTTP_BADREQUEST           = 400
	HTTP_UNAUTHORIZED         = 401
	HTTP_PAYMENTREQUIRED      = 402
	HTTP_FORBIDDEN            = 403
	HTTP_NOTFOUND             = 404
	HTTP_REQUESTTIMEOUT       = 408
	HTTP_GONE                 = 410
	HTTP_UNSUPPORTEDMEDIATYPE = 415
	HTTP_LOGIN_TIME_OUT       = 440
	HTTP_INTERNALSERVERERROR  = 500
	HTTP_SERVICEUNAVAILABLE   = 503 // Service Time Out
	HTTP_ISNTSOBAD            = 504
	TIME_FORMAT               = "2006-01-02 15:04:05.000000"
	TIME_FORMAT_REPORTING     = "2006/01/02 15:04:05"
	DateFormat                = "2006-01-02"
	GMT                       = 7
	RFC                       = time.RFC3339
	GCSTimeFormat             = "20060102"
)

const (
	ForbiddenMessage = "USERNAME_INVALID"
	SuccessMsg       = "SUCCESS"
)