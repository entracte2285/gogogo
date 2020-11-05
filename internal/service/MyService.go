package service

type MyService struct {
	Name  string
	Place string
	Job   string
}

func (m *MyService) WhoAmI() string {
	return m.Name
}
func (m *MyService) WhereAmI() string {
	return m.Place
}
func (m *MyService) WhatAmIDoing() string {
	return m.Job
}
