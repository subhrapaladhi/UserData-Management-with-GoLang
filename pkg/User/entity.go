package user

type User struct {
	email    string `bson:"email,omitempty"`
	name     string `bson:"name,omitempty"`
	phone    string `bson:"phone,omitempty"`
	password string `bson:"password,omitempty"`
}
