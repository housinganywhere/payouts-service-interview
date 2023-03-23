# Payout service

Payout service is responsible for creating payout information in the database.
Payouts are created after booking is made and tenant paid the amount to advertiser.
Several payouts for the same booking could be created, for example the booking price, deposit and utilities.

## Purpose

This is a simple repository, that reads the data from the file and write it to the database.
You will be asked to do some modifications for this code during the interview.

## Language Preference

The code is written on Go. If you prefer other language, you can rewrite it on the language of your own.
The repository utilizes [SQLite](https://sqlite.org/), if you decide to create your own implementation, use it as well.
You can use the SQL queries from `repo.go`.

## How to run

* Clone repository
* Install [Golang](https://go.dev/doc/install) on your machine
* Run `go get ./...` to load the dependencies
* Run `go run .`

