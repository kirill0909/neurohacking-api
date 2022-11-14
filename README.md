# neurohacking-api

## sign-up

### Request

* `curl -k -X POST https://46.101.253.18:8001/user/auth/sign-up -d '{"name":"UserName", "email":"user@gmail.com", "password":"UserPass"}'`

### Responses

* `{"ok":1}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`

## sign-in

### Request

* `curl -k -X POST https://46.101.253.18:8001/user/auth/sign-in -d '{"email":"user@gmail.com", "password":"UserPass"}'`

### Responses

* `{"token":"token"}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`

## category

### Request

* `curl -k -X POST https://46.101.253.18:8001/category/ -H "Authorization: Bearer token" -d '{"name":"CategoryName"}'`

### Responses

* `{"id":1}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X GET https://46.101.253.18:8001/resource/ -H "Authorization: Bearer token"`

### Responses

* `{"category":[{"id":1,"user_id":1,"name":"CategoryName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}]}`
* `{"message":"signature is invalid"}`
* `{"message":"user is not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X GET https://46.101.253.18:8001/category/1 -H "Authorization: Bearer token"`

### Responses

* `{"category":{"id":1,"user_id":1,"name":"CategoryName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X PUT https://46.101.253.18:8001/category/1 -H "Authorization: Bearer token" -d '{"name":"NewCategoryName"}'`

### Responses

* `{"id":1,"user_id":1,"name":"NewCategoryName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"change_to_time_of_last_update"}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request 

* `curl -k -X DELETE https://46.101.253.18:8001/category/1 -H "Authorization: Bearer token"`

### Responses

* `{"category":{"id":1,"user_id":1,"name":"CategoryName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

## word

### Request

* `curl -k -X POST https://46.101.253.18:8001/category/word/ -H "Authorization: Bearer token" -d '{"name":"WordName"}'`

### Responses

* `{"id":1}`
* `{"message":"invalid input body"}`
* `{"message":"invalid input value"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X GET https://46.101.253.18:8001/resource/word/ -H "Authorization: Bearer token"`

### Responses

* `{"word":[{"id":1,"user_id":1,"category_id":1,"name":"WordName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}]}`
* `{"message":"signature is invalid"}`
* `{"message":"user is not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X GET https://46.101.253.18:8001/category/1/word/1 -H "Authorization: Bearer token"`

### Responses

* `{"word":{"id":1,"user_id":1,"category_id":1,"name":"WordName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request

* `curl -k -X PUT https://46.101.253.18:8001/category/1/word/1 -H "Authorization: Bearer token" -d '{"name":"NewWordName"}'`

### Responses

* `{"word":{"id":1,"user_id":1,"category_id":1,"name":"NewWordName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`

### Request 

* `curl -k -X DELETE https://46.101.253.18:8001/category/1/word/1 -H "Authorization: Bearer token"`

### Responses

* `{"word":{"id":1,"user_id":1,"category_id":1,"name":"WordName","date_creation":"2022-11-04T10:48:30.762652Z","last_update":"2022-11-04T10:48:30.762652Z"}}`
* `{"message":"invalid id param"}`
* `{"message":"resource id not found"}`
* `{"message":"user id not found"}`
* `{"message":"user id is of invalid type"}`
