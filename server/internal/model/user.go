package model

// User 对应数据库user表
type User struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`                 // 用户名字
	Password    string `gorm:"column:password" json:"-"`                // 密码json化时要忽略避免泄露，用不到时sql中不要查询该字段
	Mobile      string `gorm:"column:mobile" json:"mobile"`             // 用于手机号
	Email       string `gorm:"column:email" json:"email"`               // 用于邮箱
	Sex         uint   `gorm:"column:sex" json:"sex"`                   // 用于展示用户性别
	Avatar      string `gorm:"column:avatar" json:"avatar"`             // 用于展示用户头像
	Online      bool   `gorm:"column:online" json:"online"`             // 用于标记用户是否在线
	EmailValid  bool   `gorm:"column:email_valid" json:"email_valid"`   // 用于校验邮箱是否为真实邮箱
	MobileValid bool   `gorm:"column:mobile_valid" json:"mobile_valid"` // 用于校验手机号是否为真实手机号
	OpenID      string `gorm:"column:open_id" json:"open_id"`           // 微信用户唯一标识
	UnionID     string `gorm:"column:union_id" json:"union_id"`         // 微信开放平台唯一标识
	Nickname    string `gorm:"column:nickname" json:"nickname"`         // 微信用户昵称
}

func (User) TableName() string {
	return "user"
}

// RegisterReq 注册用户请求，密码介于 6-32 之间
type RegisterReq struct {
	Name     string `json:"name" validate:"required,gte=3,lte=16" label:"用户名"`
	Password string `json:"password" validate:"required,gte=6,lte=32" label:"密码"`
	Mobile   string `json:"mobile" validate:"required,mobile" label:"手机号"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Sex      uint   `json:"sex" label:"性别"`
	Avatar   string `json:"avatar" label:"头像"`
}

func (rq RegisterReq) ToUserModel(pwd string) User {
	return User{
		Name:     rq.Name,
		Password: pwd,
		Mobile:   rq.Mobile,
		Email:    rq.Email,
		Sex:      rq.Sex,
		Avatar:   rq.Avatar,
		Online:   false,
	}
}

// LoginReq 登录请求，登录标识ID需要为邮件或者手机号码，密码介于6-32之间
type LoginReq struct {
	// 手机号
	// 密码
	Mobile   string `json:"mobile" validate:"required,mobile" label:"手机号"`
	Password string `json:"password" validate:"required,gte=6,lte=32" label:"密码"`
}
