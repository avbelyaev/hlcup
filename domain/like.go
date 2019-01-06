package domain

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
	"github.com/pkg/errors"
)

type Likes struct {
	Likes []Like `json:"likes"`
}

type Like struct {
	Likee int   `json:"likee,omitempty"`
	Ts    int64 `json:"ts,omitempty"`
	Liker int   `json:"liker,omitempty"`
}

func (likes *Likes) Validate() error {
	var now = time.Now().UTC().Unix()

	for _, like := range likes.Likes {
		// fail on first error
		var err = like.validate(now)
		if nil != err {
			return errors.New("invalid like: " + err.Error())
		}
	}
	return nil
}

func (like *Like) validate(tsMax int64) error {
	return validation.ValidateStruct(like,
		validation.Field(&like.Ts, validation.Required, validation.Max(tsMax)),
		validation.Field(&like.Likee, validation.Required),
		validation.Field(&like.Liker, validation.Required),
	)
}

//func validateTimestamp(ranger ...int64) func(value interface{}) error {
//	return func(value interface{}) error {
//		var v = value.(int64)
//		if 0 < len(ranger) && v < ranger[0] {
//			return errors.New("too small")
//		}
//		if 1 < len(ranger) && v > ranger[1] {
//			return errors.New("")
//		}
//		return nil
//	}
//}