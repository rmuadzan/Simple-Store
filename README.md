# Simple-Store
This website is built using golang with echo framework and mysql database with gorm.


## Installation

Install the project dependencies

```bash
  go get ./...
```

After that, make a .env file that contain this variable :
- APPLICATION_NAME=Simple Store
- MONGO_URI
- MYSQL_HOST
- MYSQL_USER
- MYSQL_PASSWORD
- MYSQL_DB
- MYSQL_PORT
- JWT_SIGNATURE_KEY
- COOKIE_HASH_KEY
- COOKIE_BLOCK_KEY (16/24/32 bytes length)

Then run following command to migrate models to your database:
```bash
  go run main.go db:migrate
```

And then run following command to store dummy product in your database:
```bash
  go run main.go db:populate
```
    
## Running App

In the project directory, you can run:

```bash
  go run main.go
```
Runs the app in the development mode.
Open http://localhost:8080 to view it in your browser.




