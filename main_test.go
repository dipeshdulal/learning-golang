package main

import "testing"

func TestHello(t *testing.T) {
	if emptyResult := hello(""); emptyResult != "Hello Dude!" {
		t.Errorf(`hello("") failed, expected %v, got %v`, "Hello Dude!", emptyResult)
	}

	if mikeResult := hello("Mike"); mikeResult != "Hello Mike!" {
		t.Error("Error for mike input")
	}

}
