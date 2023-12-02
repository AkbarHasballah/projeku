package projeku

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
type SignupPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
