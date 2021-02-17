package schema

type User struct {
	UserId	string	`json:"user_id" binding:"required"`
	Name 	string	`json:"name" binding:"required"`
	Age		int		`json:"age" binding:"required"`
}
