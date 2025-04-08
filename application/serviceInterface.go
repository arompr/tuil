package service

type Service interface {
	Increase(percentage float64)
	Decrease(percentage float64)
}
