package request

type SmsCode struct {
	Phone string `json:"phone" binding:"required,len=11"`
}