package constants

const (
	OK                              = 0
	ERR_SIGN                        = -1
	ERR_ACCESS_TOKEN                = -2
	ERR_EMPTY_DATA                  = -3
	ERR_SREVICE_UNAVAILABLE         = -4
	ERR_SREVICE_UNKNOWN             = -5
	ERR_JSON_UNMARSHAL              = -6
	ERR_EMAIL_ALREADY_EXIST         = -7
	ERR_PHONE_ALREADY_EXIST         = -8
	ERR_CREATE_USER                 = -9
	ERR_GET_USER                    = -10
	ERR_CREATE_ACCESS_TOKEN         = -11
	ERR_WRONG_PARAMS                = -12
	ERR_WRONG_PASSWORD              = -13
	ERR_UNKNOWN_METHOD              = -14
	ERR_VERIFY_CODE_NOT_FOUND       = -15
	ERR_VERIFY_CODE_EXPIRED         = -16
	ERR_SEND_CODE                   = -17
	ERR_UID_EMPTY                   = -18
	ERR_UPDATE_PROFILE              = -19
	ERR_UPDATE_PASSWORD             = -20
	ERR_USER_NOT_FOUND              = -21
	ERR_GET_PASSWORD                = -22
	ERR_USER_LOGOUT                 = -23
	ERR_UPLOAD_AVATAR               = -24
	ERR_UPDATE_AVATAR               = -25
	ERR_EMPTY_PASSWORD              = -26
	ERR_EMPTY_EMAIL                 = -27
	ERR_EMPTY_PHONE                 = -28
	ERR_CREATE_CODE                 = -29
	ERR_ADMIN_ACCOUNT_ALREADY_EXIST = -30
	ERR_ADMIN_ROLE_NAME_EXIST       = -31
	ERR_EMAIL_VERIFICATION          = -32
	ERR_ADMIN_ROLE_NOT_EXIST        = -33
	ERR_GET_COURSE_CATEGORIES       = -34
	ERR_DATABASE_OPERATION          = -35
	ERR_CREATE_EXPERTISE            = -36
	ERR_UPDATE_EXPERTISE            = -37
	ERR_EXPERTISE_NOT_FOUND         = -121
	ERR_CREATE_INTEREST             = -122
	ERR_UPDATE_INTEREST             = -123
	ERR_INTEREST_NOT_FOUND          = -124
	ERR_TEACHER_PROFILE_FAIL        = -140
	ERR_CREATE_TEACHER_PROFILE      = -141
	ERR_UPDATE_TEACHER_PROFILE      = -142
	ERR_TEACHER_PROFILE_NOT_FOUND   = -143
	ERR_NOT_CREATOR                 = -301
	ERR_COURSE_NOT_FOUND            = -302
	ERR_DELETE_SCHEDULES            = -303
	ERR_CREATE_COURSE_REVIEW        = -304
	ERR_COURSE_REVIEW_NOT_FOUND     = -305
	ERR_NOT_YOUR_COURSE_REVIEW      = -306
	ERR_UPDATE_COURSE_REVIEW        = -307
	ERR_CREATE_USER_COURSE          = -320
	ERR_UPDATE_USER_COURSE          = -321
	ERR_USER_COURSE_NOT_FOUND       = -322
	ERR_DELETE_USER_COURSE          = -323
	ERR_CREATE_ORDER                = -600
	ERR_UPDATE_ORDER                = -601
	ERR_ORDER_NOT_FOUND             = -602
	ERR_DELETE_ORDER                = -603
	ERR_CREATE_CART                 = -610
	ERR_UPDATE_CART                 = -611
	ERR_CART_NOT_FOUND              = -612
	ERR_DELETE_CART                 = -613
)

var codeErr = map[int]string{
	OK:                              "success",
	ERR_SIGN:                        "verify sign error",
	ERR_ACCESS_TOKEN:                "verify access_token error or expired",
	ERR_EMPTY_DATA:                  "empty data",
	ERR_SREVICE_UNAVAILABLE:         "service unavaliable",
	ERR_SREVICE_UNKNOWN:             "unknown service",
	ERR_JSON_UNMARSHAL:              "json unmarshal error",
	ERR_EMAIL_ALREADY_EXIST:         "email address already exist",
	ERR_PHONE_ALREADY_EXIST:         "phone number already exist",
	ERR_CREATE_USER:                 "signup user error",
	ERR_GET_USER:                    "find user error",
	ERR_CREATE_ACCESS_TOKEN:         "create token error",
	ERR_WRONG_PARAMS:                "wrong params",
	ERR_WRONG_PASSWORD:              "wrong password",
	ERR_UNKNOWN_METHOD:              "unknown method",
	ERR_VERIFY_CODE_NOT_FOUND:       "verification code not found",
	ERR_VERIFY_CODE_EXPIRED:         "verification code expired",
	ERR_SEND_CODE:                   "send verification code fail",
	ERR_UID_EMPTY:                   "uid can't be empty",
	ERR_UPDATE_PROFILE:              "update profile fail",
	ERR_UPDATE_PASSWORD:             "update password fail",
	ERR_USER_NOT_FOUND:              "user not found",
	ERR_GET_PASSWORD:                "get password fail",
	ERR_USER_LOGOUT:                 "logout fail",
	ERR_UPLOAD_AVATAR:               "upload avatar fail",
	ERR_UPDATE_AVATAR:               "update avatar fail",
	ERR_EMPTY_PASSWORD:              "password can't be empty",
	ERR_EMPTY_EMAIL:                 "email can't be empty",
	ERR_EMPTY_PHONE:                 "phone can't be empty",
	ERR_CREATE_CODE:                 "create verification code fail",
	ERR_ADMIN_ACCOUNT_ALREADY_EXIST: "admin account already exist",
	ERR_ADMIN_ROLE_NAME_EXIST:       "role name exist",
	ERR_EMAIL_VERIFICATION:          "email verification fail",
	ERR_ADMIN_ROLE_NOT_EXIST:        "role not exists",
	ERR_GET_COURSE_CATEGORIES:       "get course categories",
	ERR_DATABASE_OPERATION:          "database operation fail",
	ERR_CREATE_EXPERTISE:            "create expertise fail",
	ERR_UPDATE_EXPERTISE:            "update expertise fail",
	ERR_EXPERTISE_NOT_FOUND:         "expertise not found",
	ERR_CREATE_INTEREST:             "create interest fail",
	ERR_UPDATE_INTEREST:             "update interest fail",
	ERR_INTEREST_NOT_FOUND:          "interest not found",
	ERR_TEACHER_PROFILE_FAIL:        "teacher profile fail",
	ERR_CREATE_TEACHER_PROFILE:      "create teacher profile fail",
	ERR_UPDATE_TEACHER_PROFILE:      "update teacher profile fail",
	ERR_TEACHER_PROFILE_NOT_FOUND:   "teacher profile not found",
	ERR_NOT_CREATOR:                 "not the course creator",
	ERR_COURSE_NOT_FOUND:            "course not found",
	ERR_DELETE_SCHEDULES:            "delete course schedules fail",
	ERR_CREATE_COURSE_REVIEW:        "create course review fail",
	ERR_COURSE_REVIEW_NOT_FOUND:     "course review not found",
	ERR_NOT_YOUR_COURSE_REVIEW:      "not your course review",
	ERR_UPDATE_COURSE_REVIEW:        "update course review fail",
	ERR_CREATE_USER_COURSE:          "create user course fail",
	ERR_UPDATE_USER_COURSE:          "update user course fail",
	ERR_USER_COURSE_NOT_FOUND:       "user course not found",
	ERR_DELETE_USER_COURSE:          "delete user course",
	ERR_CREATE_ORDER:                "create order fail",
	ERR_UPDATE_ORDER:                "update order fail",
	ERR_ORDER_NOT_FOUND:             "order not found",
	ERR_DELETE_ORDER:                "delete order fail",
	ERR_CREATE_CART:                 "create shopping cart fail",
	ERR_UPDATE_CART:                 "update shopping cart fail",
	ERR_CART_NOT_FOUND:              "shopping cart not found",
	ERR_DELETE_CART:                 "delete shopping cart fail",
}

func ErrMsg(errCode int) string {
	if msg, ok := codeErr[errCode]; ok {
		return msg
	}
	return ""
}
