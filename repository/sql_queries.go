package repository

const (
	callAddUser = "CALL p_add_user($1::JSONB)"
	callGetUser = "CALL p_get_user($1::JSONB,$2::JSONB)"
)
