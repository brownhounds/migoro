module github.com/brownhounds/migoro/dispatcher

go 1.22.2

replace github.com/brownhounds/migoro/utils => ../utils

replace github.com/brownhounds/migoro/adapters => ../adapters

require (
	github.com/brownhounds/migoro/adapters v0.0.0-00010101000000-000000000000
	github.com/brownhounds/migoro/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/logrusorgru/aurora/v4 v4.0.0 // indirect
)
