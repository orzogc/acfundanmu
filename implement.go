package acfundanmu

// GetSendTime 获取弹幕发送时间
func (d *Comment) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Comment) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *Like) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Like) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *EnterRoom) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *EnterRoom) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *FollowAuthor) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *FollowAuthor) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *ThrowBanana) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *ThrowBanana) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *Gift) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Gift) GetUserInfo() UserInfo {
	return d.UserInfo
}

// GetSendTime 获取弹幕发送时间
func (d *RichText) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息，返回第一个RichTextUserInfo的UserInfo，否则返回空的UserInfo
func (d *RichText) GetUserInfo() UserInfo {
	for _, segment := range d.Segments {
		if u, ok := segment.(*RichTextUserInfo); ok {
			return u.UserInfo
		}
	}
	return UserInfo{}
}

// RichTextType 返回RichText的类型，也就是 "RichTextUserInfo"
func (*RichTextUserInfo) RichTextType() string {
	return "RichTextUserInfo"
}

// RichTextType 返回RichText的类型，也就是 "RichTextPlain"
func (*RichTextPlain) RichTextType() string {
	return "RichTextPlain"
}

// RichTextType 返回RichText的类型，也就是 "RichTextImage"
func (*RichTextImage) RichTextType() string {
	return "RichTextImage"
}
