package main

// JSONAdminUsers to keep all admin users for auth JSON
type JSONAdminUsers struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Admin    bool   `json:"admin"`
}

// OsqueryTable to show tables to query
type OsqueryTable struct {
	Name      string   `json:"name"`
	URL       string   `json:"url"`
	Platforms []string `json:"platforms"`
	Filter    string
}

// JWTData to return all the fields from a JWT token
type JWTData struct {
	Subject  string
	Email    string
	Display  string
	Username string
}
