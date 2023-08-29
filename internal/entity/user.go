package entity

type User struct {
	Balance string       `json:"balance"`
	Tags    []string     `json:"tags"`
	Friends []UserFriend `json:"friends"`
}

type UserFriend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Wrapper struct {
	Name   string
	Result []User
	Err    error
}
