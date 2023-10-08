package types

type User struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}
type Post struct {
	Title    string `json:"Title"`
	Category string `json:"Category"`
	Contacts string `json:"Contacts"`
	Details  string `json:"Details"`
	Goals    string `json:"Goals"`
}
