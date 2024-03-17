# Simple File Service [WIP]

This is a simple file upload service written in Go. It allows users to upload files to a server and store them for later retrieval.

The service provides a straightforward API for handling file uploads. Users can send a POST request to the designated endpoint with the file they want to upload, and the server will handle the rest. The uploaded files are stored on the server's file system, making them accessible for future use.

If you want to try out:

```sh
go run cmd/main.go
```

cURL

```sh
curl --location 'http://localhost:8080/upload' \
--form 'myFile=/path/to/your/file'
```
