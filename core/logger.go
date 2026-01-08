package core

type Logger interface {
	Info(string)
	Success(string)
	Error(string)
}
