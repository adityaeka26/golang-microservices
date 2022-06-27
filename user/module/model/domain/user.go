package domain

type InsertUser struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
}
type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
}
