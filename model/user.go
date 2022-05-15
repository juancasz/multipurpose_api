package model

type (
	User struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		CountryID     int    `json:"country_id"`
		UniversityID  int    `json:"university_id"`
		Username      string `json:"username"`
		HashPassword  string `json:"hash_password"`
		UserIDCreator string `json:"user_id_creator"`
	}

	UserInfo struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Country    string `json:"country"`
		University string `json:"university"`
	}
)
