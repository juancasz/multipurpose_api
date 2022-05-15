package repository

const (
	callAddUser        = "CALL p_add_user($1::JSONB)"
	callGetUser        = "CALL p_get_user($1::JSONB,$2::JSONB)"
	callEditUser       = "CALL p_edit_user($1::JSONB)"
	callDeleteUser     = "CALL p_delete_user($1::JSONB)"
	getUserCredentials = "SELECT * FROM f_get_user_credentials($1::VARCHAR)"
)
