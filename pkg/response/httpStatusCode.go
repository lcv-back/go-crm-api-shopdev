package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeMissUserId   = 20002 // Miss user id
	ErrCodeParamInvalid = 20003 // Email is invalid
)

// message map
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeMissUserId:   "miss user id",
	ErrCodeParamInvalid: "email is invalid",
}
