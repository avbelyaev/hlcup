package domain

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Fname   string `json:"fname"`
	Sname   string `json:"sname"`
	Phone   string `json:"phone"`
	Sex     string `json:"sex"`
	Birth   int    `json:"birth"`
	Country string `json:"country"`
	City    string `json:"city"`

	// dating-specific
	Joined    int      `json:"joined"`
	Status    string   `json:"status"`
	Interests []string `json:"interests"`
	Premium struct {
		Start  int `json:"start"`
		Finish int `json:"finish"`
	} `json:"premium"`
	Likes []struct {
		Ts int `json:"ts"`
		ID int `json:"id"`
	} `json:"likes"`
}
