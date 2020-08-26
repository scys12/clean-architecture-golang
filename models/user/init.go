package user

type User struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
}
