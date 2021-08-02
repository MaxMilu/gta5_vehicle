package users

import "time"

type Code struct {
	ID         uint       `bson:"_id"`
	CreatedAt  time.Time  `bson:"created_at"`
	UpdatedAt  *time.Time `bson:"updated_at"`
	CodeHeader string     `bson:"code_header"`
	CodeFooter string     `bson:"code_footer"`
	Status     bool       `bson:"status"`
}

func (Code) TableName() string {
	return "user_codes"
}
