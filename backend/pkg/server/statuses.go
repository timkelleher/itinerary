package backend

import (
	"net/http"
	"sort"
)

var statusMap = map[int]string{
	http.StatusContinue:                      "Continue",                      // 100
	http.StatusSwitchingProtocols:            "SwitchingProtocols",            // 101
	http.StatusProcessing:                    "Processing",                    // 102
	http.StatusEarlyHints:                    "EarlyHints",                    // 103
	http.StatusOK:                            "OK",                            // 200
	http.StatusCreated:                       "Created",                       // 201
	http.StatusAccepted:                      "Accepted",                      // 202
	http.StatusNonAuthoritativeInfo:          "NonAuthoritativeInfo",          // 203
	http.StatusNoContent:                     "NoContent",                     // 204
	http.StatusResetContent:                  "ResetContent",                  // 205
	http.StatusPartialContent:                "PartialContent",                // 206
	http.StatusMultiStatus:                   "MultiStatus",                   // 207
	http.StatusAlreadyReported:               "AlreadyReported",               // 208
	http.StatusIMUsed:                        "IMUsed",                        // 226
	http.StatusMultipleChoices:               "MultipleChoices",               // 300
	http.StatusMovedPermanently:              "MovedPermanently",              // 301
	http.StatusFound:                         "Found",                         // 302
	http.StatusSeeOther:                      "SeeOther",                      // 303
	http.StatusNotModified:                   "NotModified",                   // 304
	http.StatusUseProxy:                      "UseProxy",                      // 305
	http.StatusTemporaryRedirect:             "TemporaryRedirect",             // 307
	http.StatusPermanentRedirect:             "PermanentRedirect",             // 308
	http.StatusBadRequest:                    "BadRequest",                    // 400
	http.StatusUnauthorized:                  "Unauthorized",                  // 401
	http.StatusPaymentRequired:               "PaymentRequired",               // 402
	http.StatusForbidden:                     "Forbidden",                     // 403
	http.StatusNotFound:                      "NotFound",                      // 404
	http.StatusMethodNotAllowed:              "MethodNotAllowed",              // 405
	http.StatusNotAcceptable:                 "NotAcceptable",                 // 406
	http.StatusProxyAuthRequired:             "ProxyAuthRequired",             // 407
	http.StatusRequestTimeout:                "RequestTimeout",                // 408
	http.StatusConflict:                      "Conflict",                      // 409
	http.StatusGone:                          "Gone",                          // 410
	http.StatusLengthRequired:                "LengthRequired",                // 411
	http.StatusPreconditionFailed:            "PreconditionFailed",            // 412
	http.StatusRequestEntityTooLarge:         "RequestEntityTooLarge",         // 413
	http.StatusRequestURITooLong:             "RequestURITooLong",             // 414
	http.StatusUnsupportedMediaType:          "UnsupportedMediaType",          // 415
	http.StatusRequestedRangeNotSatisfiable:  "RequestedRangeNotSatisfiable",  // 416
	http.StatusExpectationFailed:             "ExpectationFailed",             // 417
	http.StatusTeapot:                        "Teapot",                        // 418
	http.StatusMisdirectedRequest:            "MisdirectedRequest",            // 421
	http.StatusUnprocessableEntity:           "UnprocessableEntity",           // 422
	http.StatusLocked:                        "Locked",                        // 423
	http.StatusFailedDependency:              "FailedDependency",              // 424
	http.StatusTooEarly:                      "TooEarly",                      // 425
	http.StatusUpgradeRequired:               "UpgradeRequired",               // 426
	http.StatusPreconditionRequired:          "PreconditionRequired",          // 428
	http.StatusTooManyRequests:               "TooManyRequests",               // 429
	http.StatusRequestHeaderFieldsTooLarge:   "RequestHeaderFieldsTooLarge",   // 431
	http.StatusUnavailableForLegalReasons:    "UnavailableForLegalReasons",    // 451
	http.StatusInternalServerError:           "InternalServerError",           // 500
	http.StatusNotImplemented:                "NotImplemented",                // 501
	http.StatusBadGateway:                    "BadGateway",                    // 502
	http.StatusServiceUnavailable:            "ServiceUnavailable",            // 503
	http.StatusGatewayTimeout:                "GatewayTimeout",                // 504
	http.StatusHTTPVersionNotSupported:       "HTTPVersionNotSupported",       // 505
	http.StatusVariantAlsoNegotiates:         "VariantAlsoNegotiates",         // 506
	http.StatusInsufficientStorage:           "InsufficientStorage",           // 507
	http.StatusLoopDetected:                  "LoopDetected",                  // 508
	http.StatusNotExtended:                   "NotExtended",                   // 510
	http.StatusNetworkAuthenticationRequired: "NetworkAuthenticationRequired", // 511
}

var statuses []Status

type Status struct {
	Code  int
	Label string
}

func Statuses() []Status {
	if len(statuses) == 0 {
		for code, label := range statusMap {
			statuses = append(statuses, Status{Code: code, Label: label})
		}
		sort.Slice(statuses, func(i, j int) bool {
			return statuses[i].Code < statuses[j].Code
		})
	}
	return statuses
}
