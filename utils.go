package main

type SettingsDB struct {
	//urlExample := "postgres://suername:password@localhost:5432/database_name"
	Username string 'json:"username"'
	Password string 'json:"password"'
	DatabasName string 'json:"database_Name"'
}
