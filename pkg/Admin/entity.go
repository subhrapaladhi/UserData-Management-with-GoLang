package admins

type Admin struct {
	Id       string `bson:"_id,omitempty"`
	Email    string `bson:"email,omitempty"`
	Name     string `bson:"name,omitempty"`
	Phone    string `bson:"phone,omitempty"`
	Password string `bson:"password,omitempty"`
}
