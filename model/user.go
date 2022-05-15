package model

type (
	User struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		CountryID     int    `json:"country_id"`
		UniversityID  int    `json:"university_id"`
		Username      string `json:"username"`
		Password      string `json:"password"`
		HashToken     string `json:"hash_token"`
		UserIDCreator string `json:"user_id_creator"`
	}

	UserInfo struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Country    string `json:"country"`
		University string `json:"university"`
	}

	UserCredentials struct {
		Id        string `json:"id"`
		HashToken string `json:"hash_token"`
	}

	UserLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
