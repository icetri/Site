package types

type Blog struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Anous    string `json:"anous"`
	FullText string `json:"full_text"`
	Now      string `json:"time"`
	Username string `json:"username"`
}

type User struct {
	Name      string
	Age       uint8
	Rating    float32
	Happiness float32
}
