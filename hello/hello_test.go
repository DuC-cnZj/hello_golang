package hello

import (
	"testing"
)

func TestSay(t *testing.T) {
	say := Say()
	if say != "hello go" {
		t.Errorf("error print :%s", say)
	}
}

func TestRun(t *testing.T) {
	s := run()
	if s != "run" {
		t.Errorf("error print :%s", s)
	}
}

func BenchmarkSay(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N ;i++  {
		Say()
	}
}
