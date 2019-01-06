package domain

import "github.com/go-ozzo/ozzo-validation"

type Likes struct {
	Likes []Like `json:"likes"`
}

type Like struct {
	Likee int `json:"likee"`
	Ts    int `json:"ts"`
	Liker int `json:"liker"`
}

// TODO not needed
func (lk *Like) Validate() error {
	return validation.ValidateStruct(lk,
		validation.Field(&lk.Likee, validation.Required), )
}
