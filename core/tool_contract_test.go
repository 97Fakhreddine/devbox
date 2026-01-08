package core

import "testing"

type FakeTool struct{}

func (f *FakeTool) ID() string                { return "fake" }
func (f *FakeTool) Name() string              { return "Fake Tool" }
func (f *FakeTool) Description() string       { return "Fake" }
func (f *FakeTool) IsInstalled(*Context) bool { return false }
func (f *FakeTool) Install(*Context) error    { return nil }

func TestToolContract(t *testing.T) {
	var _ Tool = (*FakeTool)(nil)
}
