package model

type UserAccount struct {
	Name  string `json:"user" form:"user"`
	Email string `json:"email" form:"email" binding:"email"`
	Pwd   string `json:"password" form:"password"`
	Pwd2  string `json:"confirm-password" form:"confirm-password" validate:"eqfield=Pwd"`
}

func (ua UserAccount) GetName() string {
	return ua.Name
}

func (ua UserAccount) GetPwd() string {
	return ua.Pwd
}

type User struct {
	Name string `json:"name,omitempty" form:"name"`
	Id   int64  `json:"id,omitempty" form:"name"`
	Age  int8   `json:"age,omitempty" form:"name"`
}
