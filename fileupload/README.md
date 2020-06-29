### File Upload in Golang

Examples for file upload can be found at; [upload-file](https://github.com/gin-gonic/examples/tree/master/upload-file)

To upload files in `golang, gin framework` there are basically two methods;

- `c.FormFile("file")` to get the file from post request body.

- `c.SaveUploadedFile(file, destination)` moves the file in post request to actual directory or storage. 

- `form, err := c.MultipartForm()` to get all the data from form data. And then `form.File` is used to find the data passed to multipart form.