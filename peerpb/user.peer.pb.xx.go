package peerpb

//UserRegister

func (x *UserRegisterReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *UserRegisterResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

// UserToken

func (x *UserTokenReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *UserTokenResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

// CreateRobot

func (x *CreateRobotReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *CreateRobotResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

// RefreshUserToken

func (x *RefreshUserTokenReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *RefreshUserTokenResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}

// RevokeUserToken

func (x *RevokeUserTokenReq) SetHeader(header *RequestHeader) {
	x.Header = header
}

func (x *RevokeUserTokenResp) SetHeader(header *ResponseHeader) {
	x.Header = header
}
