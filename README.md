## Homeapp: app for my lan
###### Idea: sometimes my family needs to collect some files in one place, but always connect phones to pc for sharing data isn't comfortable, it waste a lot of time. This simple API is a fast and good solution for us.
___
<p align="center"><img src="https://github.com/aeonva1ues/homeapp/blob/main/preview.png" width="200px" /></p>

### Installation
Clone repository
```
git clone https://github.com/aeonva1ues/homeapp.git
```
Create env file
```
cp configs/.env.example configs/.env
```
Set variables in new .env file

Example:
```
HOST=127.0.0.1:8080  # web server's host
# database data
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=12345
DATABASE_NAME=homeapp
SSL_MODE=disable
# ===============
UPLOADS=F://uploads/  # on this path files will be loaded
```
### Run web server
Using make:
```
make compose-up
```
or using docker-compose:
```
docker-compose --env-file ./configs/.env up --build -d && docker-compose logs -f
```
### Usage
Try to add some files by curl:
```
curl -X POST http://127.0.0.1:8080/file/uploads -F "files=<PATH_TO_FILE>" -F "files=<PATH_TO_FILE>" -H "Content-Type: multipart/form-data"
```
or by form on file/ endpoint http://127.0.0.1:8080/file/
