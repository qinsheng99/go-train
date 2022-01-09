package ceshi

type CeShiService interface {
	Ceshi(string) string
}

type CC struct{}

func NewCeshi() CeShiService {
	return CC{}
}

func (c CC) Ceshi(name string) string {
	return name
}
