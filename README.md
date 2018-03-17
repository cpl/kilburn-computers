# Kilburn Computers
Go web application that shows computer availability in the Kilburn Building at the University of Manchester

## Installing the project
Open a Terminal window and type
```
go get github.com/colinx05/kilburn-computers
```

## Running the project
Open a Terminal window and type `kilburn-computers`

## Using the API
`/api/list` returns a list of labs for which information is available:
```
{
    "lab_list": ["G23", "LF5", "LF6", "LF17", "LF31", "Toot0", "Toot1", "Collab1", "Collab2", "Quiet", "MSc"]
}
```

`/api/{lab}` returns a JSON with info about a lab. 
Example:
```
{
	"name": "LF31",
	"location":"Lower First",
	"description": "Has HTMI and MicroUSB cables",
	"count": 3,
	"used": 2,
	"computers": [
		{ "name": "c-0001", "used": true },
		{ "name": "c-0002", "used": true },
		{ "name": "c-0003", "used": false }
	]
}
```