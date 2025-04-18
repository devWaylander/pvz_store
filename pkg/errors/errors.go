package errors

const (
	// ===================-  COMMON  -===================
	ErrDecodeCtx = "ERR_FAILED_TO_DECODE_CONTEXT_CLAIMS"
	// ===================-  USER  -===================
	ErrUserNotFound  = "ERR_USER_NOT_FOUND"
	ErrUserExist     = "ERR_USER_ALREADY_EXIST"
	ErrForbiddenRole = "ACCESS_IS_FORBIDDEN_FOR_CURRENT_ROLE"
	// ===================-  AUTH  -===================
	ErrEncodeJWT           = "ERR_FAILED_TO_ENCODE_JWT"
	ErrGenUUID             = "ERR_FAILED_TO_GEN_RANDOM_UUID"
	ErrWrongPassword       = "ERR_WRONG_PASSWORD"
	ErrWrongPasswordFormat = "ERR_WRONG_PASSWORD_FORMAT"
	ErrInvalidToken        = "ERR_INVALID_AUTH_TOKEN"
	ErrInvalidClaims       = "ERR_CANNOT_PARSE_CLAIMS"
	ErrUnauthenticated     = "ERR_UNAUTHENTICATED"
	// ===================-  PVZ  -===================
	ErrWrongRegDate = "ERR_DATE_FROM_FUTURE_FOR_REGISTRATION_DATE"
	ErrPVZExist     = "ERR_PVZ_ALREADY_EXIST"
	// ===================-  RECEPTION  -===================
	ErrPVZDoesntExist       = "ERR_PVZ_ID_DOESNT_EXIST"
	ErrReceptionDoesntExist = "ERR_RECEPTION_DOESNT_EXIST"
	ErrReceptionExist       = "ERR_RECEPTION_ALREADY_IN_PROGRESS_STATUS"
	ErrWrongReceptionStatus = "ERR_RECEPTION_ALREADY_IS_CLOSED"
	// ===================-  PRODUCT  -===================
	ErrNoProductsToDelete = "ERR_NO_PRODUCTS_TO_DELETE"
)
