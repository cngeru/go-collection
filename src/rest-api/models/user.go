package models

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (user User) Validate() *HttpError {
	if user.FirstName == "" {
		return BadRequestError("invalid first name")
	}
	if user.LastName == "" {
		return BadRequestError("invalid last name")
	}
	if user.Email == "" {
		return BadRequestError("invalid email address")
	}
	return nil
}
