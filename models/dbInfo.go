package models

import "fmt"

// DBInfo data need to connect db
type DBInfo struct {
	Username string
	Password string
	DB       string
	Host     string
	Port     int
}

// GetConnectString return string connection from DBInfo
func (info DBInfo) GetConnectString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d",
		info.Username, info.Password, info.DB, info.Host, info.Port)
}
