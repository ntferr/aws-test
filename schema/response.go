package schema

type Response struct {
	User	User	`json:"user" binding:"required"`
	Comment	string	`json:"comment" binding:"required"`
}
