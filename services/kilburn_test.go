package services

import (
	"testing"
)

func TestGetLabList(t *testing.T) {

	correctLabList := []string{
		"G23",
		"LF5",
		"LF6",
		"LF17",
		"LF31",
		"Toot0",
		"Toot1",
		"Collab1",
		"Collab2",
		"Quiet",
		"MSc",
	}

	returnedLabList := GetLabList()

	if len(returnedLabList) != len(correctLabList) {
		t.Error("List sizes don't match!")
		return
	}

	for i, lab := range returnedLabList {
		if lab != correctLabList[i] {
			t.Error("Wrong list is returned!")
			return
		}
	}
}

// Test GetLabInfo given an existing lab (LF31)
func TestGetLabInfo(t *testing.T) {

	const realLabName = "LF31"

	labInfo, err := GetLabInfo(realLabName)
	if err != nil {
		t.Errorf("GetLabInfo(\"%s\") returns an error", realLabName)
		return
	}

	if labInfo.Name != realLabName {
		t.Errorf("GetLabInfo(\"%s\") returns the wrong name", realLabName)
		return
	}

	if labInfo.Description == "" {
		t.Errorf("GetLabInfo(\"%s\") has no description", realLabName)
		return
	}
}

//Test GetLabInfo given a wrong lab name
func TestGetLabInfoWrongLabName(t *testing.T) {

	_, err := GetLabInfo("FooBar")
	if err == nil {
		t.Error("GetLabInfo(\"FooBar\") should return error")
		return
	}
}
