## Homeapp: app for my localhost
###### Idea: sometimes my family needs to collect some files in one place, but always connect phones to pc for sharing data isn't comfortable, it waste a lot of time. This simple API is a simple and good solution for us.
___
#### Try to add some files by curl
```
curl -X POST http://127.0.0.1:8080/file/uploads -F "files=@/Users/0.mp4" -F "files=@/Users/1.jpg" -F "files=@/Users/2.jpg" -F "files=@/Users/3.jpg" -F "files=@/Users/4.jpg" -F "files=@/Users/5.mp4" -H "Content-Type: multipart/form-data"
```
