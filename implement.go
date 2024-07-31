package acfundanmu

// 验证是否实现了 DanmuMessage 接口
var _ DanmuMessage = (*Comment)(nil)
var _ DanmuMessage = (*Like)(nil)
var _ DanmuMessage = (*EnterRoom)(nil)
var _ DanmuMessage = (*FollowAuthor)(nil)
var _ DanmuMessage = (*ThrowBanana)(nil)
var _ DanmuMessage = (*Gift)(nil)
var _ DanmuMessage = (*RichText)(nil)
var _ DanmuMessage = (*JoinClub)(nil)
var _ DanmuMessage = (*ShareLive)(nil)

// GetSendTime 获取弹幕发送时间
func (d *Comment) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Comment) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *Like) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Like) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *EnterRoom) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *EnterRoom) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *FollowAuthor) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *FollowAuthor) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *ThrowBanana) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *ThrowBanana) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *Gift) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *Gift) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *RichText) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息，返回第一个 RichTextUserInfo 的 UserInfo，否则返回 nil
func (d *RichText) GetUserInfo() *UserInfo {
	for _, segment := range d.Segments {
		if u, ok := segment.(*RichTextUserInfo); ok {
			info := u.UserInfo
			return &info
		}
	}
	return nil
}

// GetSendTime 获取弹幕发送时间，实际上返回的是用户加入守护团的时间
func (d *JoinClub) GetSendTime() int64 {
	return d.JoinTime
}

// GetUserInfo 获取弹幕的用户信息，实际上返回的是加入守护团的用户的信息
func (d *JoinClub) GetUserInfo() *UserInfo {
	info := d.FansInfo
	return &info
}

// GetSendTime 获取弹幕发送时间
func (d *ShareLive) GetSendTime() int64 {
	return d.SendTime
}

// GetUserInfo 获取弹幕的用户信息
func (d *ShareLive) GetUserInfo() *UserInfo {
	info := d.UserInfo
	return &info
}

// 验证是否实现了 RichTextSegment 接口
var _ RichTextSegment = (*RichTextUserInfo)(nil)
var _ RichTextSegment = (*RichTextPlain)(nil)
var _ RichTextSegment = (*RichTextImage)(nil)

// RichTextType 返回 RichText 的类型，也就是 "RichTextUserInfo"
func (*RichTextUserInfo) RichTextType() string {
	return "RichTextUserInfo"
}

// RichTextType 返回 RichText 的类型，也就是 "RichTextPlain"
func (*RichTextPlain) RichTextType() string {
	return "RichTextPlain"
}

// RichTextType 返回 RichText 的类型，也就是 "RichTextImage"
func (*RichTextImage) RichTextType() string {
	return "RichTextImage"
}
