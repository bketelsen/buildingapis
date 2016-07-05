package main

import "testing"

func TestIDOrListWithID(t *testing.T) {
	const expect = "1"
	id, e := idOrList(courseBase, "/api/courses/1")
	if e != nil {
		t.Error(e)
	}
	if id != expect {
		t.Errorf("Expected %s, got %s", expect, id)
	}
}

func TestIDOrListWithIDAndTrailingSlash(t *testing.T) {
	const expect = "1"
	id, e := idOrList(courseBase, "/api/courses/1/")
	if e != nil {
		t.Error(e)
	}
	if id != expect {
		t.Errorf("Expected %s, got %s", expect, id)
	}
}

func TestIDOrListWithNoIDAndTrailingSlash(t *testing.T) {
	id, e := idOrList(courseBase, "/api/courses/")
	if e != nil {
		t.Error(e)
	}
	if id != "" {
		t.Errorf("Expected empty string  got %s", id)
	}
}

func TestIDOrListWithNoIDNoTrailingSlash(t *testing.T) {
	id, e := idOrList(courseBase, "/api/courses")
	if e != nil {
		t.Error(e)
	}
	if id != "" {
		t.Errorf("Expected empty string  got %s", id)
	}
}

func TestIDOrListWithExtras(t *testing.T) {
	id, e := idOrList(courseBase, "/api/courses/1/blue/red")
	if e == nil {
		t.Error("Expected error with garbage path")
	}
	if id != "" {
		t.Errorf("Expected no id got %s", id)
	}
}
