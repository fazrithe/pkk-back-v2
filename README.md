## Setting configuration file
Create a .env file in the root of the project and set the following parameters

```
db_name = database_name       # Name of database
db_user = user                # Database username
db_pass = secret              # Database password
db_type = mysql               # MySQL driver
db_host = localhost           # Database host
db_port = 3306                # Database port
charset = utf8                # Database charset
parse_time = True             # Database parse time
web_port = 8085               # Port to serve api
prefix = /api/v1              # API route sub route prefix

# jwt config
access_token_expire   =       # in minutes, default is 15 minutes if no value is passed

# this is not used yet as I haven't implemented refresh tokens yet # TODO
refresh_token_expire  = 72    # in hours, default is 48 minutes if no value is passed
```

## Running the project

`go run *.go`

## Database Table Creation Statement
Use the following DDL (Data Definition Language) to create the users table.

``` SQL
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
```

## API Endpoints & Usage

To be able to login, you need to use the create new user endpoint to set up a user account.

* POST    `api/v1/login` login with username and password in your db to get token back
* POST    `api/v1/register` create a new account with a username, password and name
* GET     `api/v1/users` retrieve all users
* GET     `api/v1/users/1` retrieve user with id = 1
* POST    `api/v1/users` create a new user
* PUT     `api/v1/users/1` update the record with id = 1
* DELETE  `api/v1/users/1` delete the user with id = 1

### To create a new user

1. POST `api/v1/register`

```
{
	"Name": "Joe Bloke",
	"Username": "joe.bloke@fake-domain.com",
    "Password": "secret"
}
```

*** Output ***

Note that the current implementation still returns the encrypted password, this needs to be removed from the response.

```
{
    "message": "success",
    "status": true,
    "user": {
        "ID": 1,
        "CreatedAt": "2019-05-06T00:54:22.09382+01:00",
        "UpdatedAt": "2019-05-06T00:54:22.09382+01:00",
        "DeletedAt": null,
        "Name": "Joe Bloke",
        "Username": "joe.bloke@fake-domain.com"
    }
}
```

2. POST `api/v1/login`



*** Input ***
```
{
  "Username": "joe.bloke@fake-domain.com"
  "Password": "secret"
}
```

*** Output ***

```
{
    "data": {
        "access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkwMDAwMDAwMDAwMCwicGFzc3dvcmQiOiJzZWNyZXQiLCJ1c2VybmFtZSI6ImFtYXZpQHh5ei5jb20ifQ.WJ5VMnH5ijHQOZhUlrrnrh7NCYfFpww3jBz26EkRsHQ"
    },
    "message": "success",
    "status": true
}
```#   m a s t e r - g o l a n g - a u t h  
 #   p k k - b a c k - v 2  
 #   p k k - b a c k - v 2  
 #   r e q u e s t _ p r o d u c t s  
 