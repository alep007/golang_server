# golang_server
little practice to run a server with go

## Getting Started

Follow the instruction to copy and run the project.

### Prerequisites

* Install Go from its page at https://golang.org/.

### Installing

* Clone the project to your src Golang project folder, on windows its on C:\Users\go\src. 
* In the folder open a terminal an write the following:
```
go get -u github.com/golang/dep/cmd/dep
```
dep init -v
```
dep ensure -v 
```
* Execute the http server:
```
go run main.go
```

* Open an Internet Browser and go to http://localhost:9000/documents

### Output

* The program will return all MD5 checksum, names and sizes of the files located in the "files" folder in JSON format.

Example:
Making a GET request to http://localhost:9000/documents
The server will respond with a dictionary of all files located in "files" folder in JSON format:

```
[
    {
        "ID": "0f57603976bc07ab482ce4945a56e7b8",
        "Name": "DIRECCIONES SAE BOLIVIA.docx",
        "Size": 15718
    },
    {
        "ID": "96455886211813030e62dbe50797df35",
        "Name": "Icono Sae-01.png",
        "Size": 23561
    }
]
```

### To Do
* will add some test 
* file main.go will be splited in various go files
