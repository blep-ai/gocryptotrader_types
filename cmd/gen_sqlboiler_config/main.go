package main

type driverConfig struct {
	DBName    string   `json:"dbname,omitempty"`
	Host      string   `json:"host,omitempty"`
	Port      uint16   `json:"port,omitempty"`
	User      string   `json:"user,omitempty"`
	Pass      string   `json:"pass,omitempty"`
	Schema    string   `json:"schema,omitempty"`
	SSLMode   string   `json:"sslmode,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
}
