package main

import "testing"

func TestServer(t *testing.T) {
	r := Router()
	srv := Server(r)
	if srv == nil {
		t.Fail()
	}
}

// func TestMain(m *testing.M) {
// 	m.Run()
// }
