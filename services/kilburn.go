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

// GetLabList ...
func GetLabList() []string {
	return []string{
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
}

// GetLabInfo ...
func GetLabInfo(lab string) (LabInfo, error) {
	switch strings.ToLower(lab) {
	case
		"g23",
		"lf5",
		"lf6",
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

/*
{
	"name": "LF31",
	"location":"Lower First",
	"description": "Has HTMI and MicroUSB cables",
	"count": 3,
	"used": 1,
	"computers": [
		{
			"name": "e-0001",
			"used": true
		},
		{
			"name": "e-0002",
			"used": true
		},
		{
			"name": "e-0003",
			"used": false
		}
	]
}
*/
