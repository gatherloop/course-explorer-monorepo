module course-explorer-monorepo

go 1.18

replace github.com/gatherloop/course-explorer-monorepo => ./libs/api/__generated__/contract

require (
	github.com/gatherloop/course-explorer-monorepo v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.5.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
