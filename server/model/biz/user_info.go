package biz

type UserInfo struct {
	Uid        string `json:"uid"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Username   string `json:"username"`
	CreateDate string `json:"createDate"`
	UpdateDate string `json:"updateDate"`
	Token      string `json:"token"`
}
