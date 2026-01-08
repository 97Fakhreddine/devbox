package testutil

type FakeLogger struct{}

func (f *FakeLogger) Info(string)    {}
func (f *FakeLogger) Success(string) {}
func (f *FakeLogger) Error(string)   {}
