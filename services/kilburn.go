package services

import (
	"errors"
	"strings"
)

// LabInfo ...
type LabInfo struct {
	Name        string     `json:"name"`
	Location    string     `json:"location"`
	Description string     `json:"description"`
	Count       int        `json:"total"`
	Used        int        `json:"used"`
	Computers   []Computer `json:"computers"`
}

// Computer ...
type Computer struct {
	Name string `json:"name"`
	Used bool   `json:"used"`
}

// labList is the list of labs in the Kilburn building for which I'll implement an API
var labList = []string{
	"G23",
	"LF17",
	"LF31",
	"Toot0",
	"Toot1",
	"Collab1",
	"Collab2",
	"Quiet",
	"MSc",
}

// GetLabList returns the list of labs
func GetLabList() []string {
	return labList
}

// GetLabInfo ...
func GetLabInfo(lab string) (LabInfo, error) {
	switch strings.ToLower(lab) {
	case
		"g23",
		"lf17",
		"lf31",
		"toot0",
		"toot1",
		"collab1",
		"collab2",
		"quiet",
		"msc":
		return LabInfo{
			Name:        "LF31",
			Location:    "Lower first",
			Description: "Has HDMI and MicroUSB cables",
			Count:       3,
			Used:        2,
			Computers: []Computer{
				{
					Name: "e-0001",
					Used: true,
				},
				{
					Name: "e-0002",
					Used: true,
				},
				{
					Name: "e-0003",
					Used: false,
				},
			},
		}, nil
	default:
		return LabInfo{}, errors.New("Lab does not exist")
	}

}
