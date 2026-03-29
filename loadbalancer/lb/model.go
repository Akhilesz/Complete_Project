package main

type Server struct {
	Url           string `yaml:"url"`
	IsAlive       bool   `yaml:"isAlive"`
	RequestsCount uint64 `yaml:"requestsCount"`
}

type ServerPool struct {
	Servers []*Server `yaml:"servers"`
}
