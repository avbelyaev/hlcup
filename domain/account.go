package domain

import (
	"errors"
	v "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

const (
	PHONE_MAX_LEN = 16
	MAX_LEN_100   = 100
	MAX_LEN_50    = 50
)

var BIRTH_MIN_VALUE = time.Date(1950, 1, 1,
	0, 0, 0, 0, time.UTC).Unix()
var BIRTH_MAX_VALUE = time.Date(2005, 1, 1,
	0, 0, 0, 0, time.UTC).Unix()

var JOINED_MIN_VALUE = time.Date(2011, 1, 1,
	0, 0, 0, 0, time.UTC).Unix()
var JOINED_MAX_VALUE = time.Date(2018, 1, 1,
	0, 0, 0, 0, time.UTC).Unix()

var PREMIUM_MIN_VALUE = time.Date(2018, 1, 1,
	0, 0, 0, 0, time.UTC).Unix()

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID      int    `json:"id,omitempty"`
	Email   string `json:"email,omitempty"`
	Fname   string `json:"fname,omitempty"` // optional
	Sname   string `json:"sname,omitempty"` // optional
	Phone   string `json:"phone,omitempty"` // optional
	Sex     string `json:"sex,omitempty"`
	Birth   int64  `json:"birth,omitempty"`
	Country string `json:"country,omitempty"` // optional
	City    string `json:"city,omitempty"`    // optional

	// dating-specific
	Joined    int64           `json:"joined,omitempty"`
	Status    string          `json:"status,omitempty"`
	Interests []string        `json:"interests"`         // can be empty
	Premium   *AccountPremium `json:"premium,omitempty"` // optional
	Likes     []AccountLike   `json:"likes"`             // can be empty
}

type AccountPremium struct {
	Start  int64 `json:"start,omitempty"`
	Finish int64 `json:"finish,omitempty"`
}

type AccountLike struct {
	Ts int64 `json:"ts,omitempty"`
	ID int   `json:"id,omitempty"`
}

//func (accs *Accounts) Validate() error {
//	for _, acc := range accs.Accounts {
//		// fail on first error
//		var err = acc.Validate()
//		if nil != err {
//			return errors.New("invalid account: " + err.Error())
//		}
//	}
//	return nil
//}

func (a *Account) Validate() error {
	return v.ValidateStruct(a,
		v.Field(&a.ID, v.Required),
		v.Field(&a.Email, v.Required, is.Email, v.Length(0, MAX_LEN_100)),
		v.Field(&a.Fname, v.Length(0, MAX_LEN_50)),
		v.Field(&a.Sname, v.Length(0, MAX_LEN_50)),
		v.Field(&a.Phone, v.Length(0, PHONE_MAX_LEN)),
		v.Field(&a.Sex, v.Required, v.In("f", "m")),
		v.Field(&a.Birth, v.Required, v.Max(BIRTH_MAX_VALUE), v.Min(BIRTH_MIN_VALUE)),
		v.Field(&a.Country, v.Length(0, MAX_LEN_50)),
		v.Field(&a.City, v.Length(0, MAX_LEN_50)),

		// dating-specific
		v.Field(&a.Joined, v.Required, v.Max(JOINED_MAX_VALUE), v.Min(JOINED_MIN_VALUE)),
		v.Field(&a.Status, v.Required, v.In("свободны", "заняты", "всё сложно")),
		v.Field(&a.Interests, v.By(validateInterests)),
		v.Field(&a.Premium, v.By(validatePremium)),
		v.Field(&a.Likes, v.By(validateLikes)),
	)
}

func validateInterests(value interface{}) error {
	var val, _ = value.([]string)
	for _, interest := range val {
		if len(interest) > MAX_LEN_100 {
			return errors.New("invalid interest len > 100")
		}
	}
	return nil
}

func validatePremium(value interface{}) error {
	var prem = value.(*AccountPremium)
	if nil != prem {
		return v.ValidateStruct(prem,
			v.Field(&prem.Start, v.Required, v.Min(PREMIUM_MIN_VALUE)),
			v.Field(&prem.Finish, v.Required, v.Min(PREMIUM_MIN_VALUE)),
		)
	}
	// account.premium can be null
	return nil
}

func validateLikes(value interface{}) error {
	var likes, ok = value.([]AccountLike)
	if ok {
		var now = time.Now().UTC().Unix()
		for _, like := range likes {
			var err = v.ValidateStruct(&like,
				v.Field(&like.ID, v.Required),
				v.Field(&like.Ts, v.Max(now)),
			)
			if nil != err {
				return errors.New("invalid account.likes: " + err.Error())
			}
		}
		return nil
	}
	return errors.New("account.likes is null")
}
