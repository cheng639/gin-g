package request

type Login struct {
	Phone string `json:"phone" binding:"required,len=11"`
	Code string `json:"code" binding:"required"`
}
