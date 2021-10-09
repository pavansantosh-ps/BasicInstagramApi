# BasicInstagramApi


* install all dep `go install`
* run the server `go run main.go`
* run unit tests `go test`

## After running the server use the following commands to test the apis manually

* Create user 
```
curl -X POST \
  'http://127.0.0.1:8080/users' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)' \
  -H 'Content-Type: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "pavan",
    "email": "pavan@gmail.com",
    "password": "sbbjefakjbke"
}'
```

* Get user by id
```
curl -X GET \
  'http://127.0.0.1:8080/users/616167f159cfa187f431f77c' \
  -H 'Accept: */*'
```

* Create post
```
curl -X POST \
  'http://127.0.0.1:8080/posts/' \
  -H 'Accept: */*' \
  -H 'User-Agent: Thunder Client (https://www.thunderclient.io)' \
  -H 'Content-Type: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "userid": "616167f159cfa187f431f77c",
    "caption": "santosh",
    "imageurl": "http://google.com"
}'
```

* Get post by id
```
curl -X GET \
  'http://127.0.0.1:8080/posts/616175b639d5b491849d8897' \
  -H 'Accept: */*' 
```

* Get all posts by userid
```
curl -X GET \
  'http://127.0.0.1:8080/posts/users/616167f159cfa187f431f77c' \
  -H 'Accept: */*'
```

* you can also add pagination
```
curl -X GET \
  'http://127.0.0.1:8080/posts/users/616167f159cfa187f431f77c?page=1&size=3' \
  -H 'Accept: */*'
```

### Note

* Before running the app please make sure Mongodb is installed 
* The host and port of mongodb server are hardcoded as `mongodb://localhost:27017`. If required it can be changed in helper.go