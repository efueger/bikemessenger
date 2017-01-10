package main

import "testing"

func TestSetup(t *testing.T) {
	srv := Setup()
	if srv == nil {
		t.Fail()
	}
}

// func TestMain(m *testing.M) {
// 	m.Run()
// }
