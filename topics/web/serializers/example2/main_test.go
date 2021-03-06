package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_EncodeUser_Blank(t *testing.T) {
	bb := &bytes.Buffer{}
	err := EncodeUser(bb, User{})
	if err != nil {
		t.Fatal(err)
	}

	act := strings.TrimSpace(bb.String())
	exp := `{"Admin":false,"Bio":null,"CreatedAt":"0001-01-01T00:00:00Z","first_name":""}`

	if act != exp {
		t.Fatalf("expected %s to equal '%s'", exp, act)
	}
}

func Test_EncodeUser_WithData(t *testing.T) {
	bb := &bytes.Buffer{}
	err := EncodeUser(bb, User{FirstName: "Mary", LastName: "Jane"})
	if err != nil {
		t.Fatal(err)
	}

	act := strings.TrimSpace(bb.String())

	exp := `{"Admin":false,"Bio":null,"CreatedAt":"0001-01-01T00:00:00Z","LastName":"Jane","first_name":"Mary"}`

	if act != exp {
		t.Fatalf("expected %s to equal '%s'", exp, act)
	}
}
