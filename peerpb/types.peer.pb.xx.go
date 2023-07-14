package peerpb

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthErrorType string

const (
	AuthErrorTypeExpired AuthErrorType = "expired"
	AuthErrorTypeInvalid AuthErrorType = "invalid"
	AuthErrorTypeReplace AuthErrorType = "replace"
)

const (
	ServerError         = "server_error"
	ParamError          = "param_error"
	AuthFailed          = "auth_failed"
	PublicKeyError      = "public_key_error"
	RsaDecryptError     = "rsa_decrypt_error"
	RobotNotAllowCreate = "robot_not_allow_create"
	//"robot_nickname_required"
	RobotNicknameRequired = "robot_nickname_required"
	//"robot_avatar_required"
	RobotAvatarRequired = "robot_avatar_required"
	//robot_id_exist
	RobotIdExist = "robot_id_exist"

	//captcha_error
	CaptchaError = "captcha_error"
	//"platform_not_allow"
	PlatformNotAllow = "platform_not_allow"
	// "password_invalid"
	PasswordInvalid = "password_invalid"
	//"sms_code_invalid"
	SmsCodeInvalid = "sms_code_invalid"
	//"phone_invalid"
	PhoneInvalid = "phone_invalid"
	//"email_invalid"
	EmailInvalid = "email_invalid"
	// "email_code_invalid"
	EmailCodeInvalid = "email_code_invalid"
	//"login_failed"
	LoginFailed = "login_failed"

	//username_required
	UsernameRequired = "username_required"
	//username_format_error
	UsernameFormatError = "username_format_error"
	//username_lock_error
	UsernameLockError = "username_lock_error"
	//username_already_exists
	UsernameAlreadyExists = "username_already_exists"
	//password_salt_required
	PasswordSaltRequired = "password_salt_required"
	//password_required
	PasswordRequired = "password_required"
	//phone_required
	PhoneRequired = "phone_required"
	//phone_code_required
	PhoneCodeRequired = "phone_code_required"
	//phone_format_error
	PhoneFormatError = "phone_format_error"
	//phone_code_error
	PhoneCodeError = "phone_code_error"
	//sms_code_required
	SmsCodeRequired = "sms_code_required"
	//sms_code_error
	SmsCodeError = "sms_code_error"
	//phone_lock_error
	PhoneLockError = "phone_lock_error"
	//phone_already_exists
	PhoneAlreadyExists = "phone_already_exists"
	//email_required
	EmailRequired = "email_required"
	//email_format_error
	EmailFormatError = "email_format_error"
	//email_code_required
	EmailCodeRequired = "email_code_required"
	//email_code_error
	EmailCodeError = "email_code_error"
	//email_lock_error
	EmailLockError = "email_lock_error"
	//email_already_exists
	EmailAlreadyExists = "email_already_exists"
	//nickname_required
	NicknameRequired = "nickname_required"
	//avatar_required
	AvatarRequired = "avatar_required"
	//captcha_required
	CaptchaRequired = "captcha_required"

	//friend_apply_already
	FriendApplyAlready = "friend_apply_already"
	//friend_apply_from_user_status_error
	FriendApplyFromUserStatusError = "friend_apply_from_user_status_error"
	//friend_apply_from_user_status_error
	FriendApplyToUserStatusError = "friend_apply_to_user_status_error"
	//friend_apply_to_user_destroy
	FriendApplyToUserDestroy = "friend_apply_to_user_destroy"
	//friend_apply_from_user_role_error
	FriendApplyFromUserRoleError = "friend_apply_from_user_role_error"
	//friend_apply_to_user_role_error
	FriendApplyToUserRoleError = "friend_apply_to_user_role_error"
	//friend_apply_answer_question_error
	FriendApplyAnswerQuestionError = "friend_apply_answer_question_error"
	//friend_apply_no_one
	FriendApplyNoOne = "friend_apply_no_one"
	//friend_apply_error
	FriendApplyError = "friend_apply_error"
	//friend_apply_are_friend
	FriendApplyAreFriend = "friend_apply_are_friend"

	//group_create_not_allow_role
	GroupCreateNotAllowRole = "group_create_not_allow_role"
	//group_create_joined_max_count
	GroupCreateJoinedMaxCount = "group_create_joined_max_count"
	//group_create_required_name
	GroupCreateRequiredName = "group_create_required_name"
	//group_create_required_avatar
	GroupCreateRequiredAvatar = "group_create_required_avatar"
)

func marshalToString(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func NewInvalidDataError(msg string) *ResponseHeader {
	logx.Errorf("invalid data error: %s", msg)
	data := &ToastActionData{
		Level:   ToastActionData_ERROR,
		Message: ParamError,
	}
	return &ResponseHeader{
		Code:       ResponseCode_INVALID_DATA,
		ActionType: ResponseActionType_TOAST_ACTION,
		ActionData: marshalToString(data),
	}
}

func NewInvalidMethodError() *ResponseHeader {
	return &ResponseHeader{
		Code: ResponseCode_INVALID_METHOD,
	}
}

func NewServerError(env string, err error) *ResponseHeader {
	data := &ToastActionData{
		Level:   ToastActionData_ERROR,
		Message: ServerError,
	}
	if env != "pro" && err != nil {
		data = &ToastActionData{
			Level:   ToastActionData_ERROR,
			Message: err.Error(),
		}
	}
	return &ResponseHeader{
		Code:       ResponseCode_SERVER_ERROR,
		ActionType: ResponseActionType_TOAST_ACTION,
		ActionData: marshalToString(data),
	}
}

func NewOkHeader() *ResponseHeader {
	return &ResponseHeader{
		Code:       ResponseCode_SUCCESS,
		ActionType: ResponseActionType_NONE_ACTION,
		ActionData: "",
	}
}

// AuthErrorExtra

type AuthErrorExtra struct {
	Type    AuthErrorType `json:"type"`
	Message string        `json:"message"`
}

func NewAuthError(typ AuthErrorType, message string) *ResponseHeader {
	extra := &AuthErrorExtra{
		Type:    typ,
		Message: message,
	}
	buf, _ := json.Marshal(extra)
	return &ResponseHeader{
		Code:  ResponseCode_UNAUTHORIZED,
		Extra: string(buf),
	}
}

func NewForbiddenError() *ResponseHeader {
	return &ResponseHeader{
		Code: ResponseCode_FORBIDDEN,
	}
}

func NewToastHeader(level ToastActionData_Level, message string) *ResponseHeader {
	data := &ToastActionData{
		Level:   level,
		Message: message,
	}
	return &ResponseHeader{
		Code:       ResponseCode_INVALID_DATA,
		ActionType: ResponseActionType_TOAST_ACTION,
		ActionData: marshalToString(data),
		Extra:      "",
	}
}
