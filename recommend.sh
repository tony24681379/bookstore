#!/bin/sh
go run main.go recommend create test2 1 \
--capacity 10 \
--day-high 10 \
--day-low 2 \
--week-high 35 \
--week-low 10 \
--month-high 110 \
--month-low 40 \
--stack-high 50 \
--stack-low 10 \
--stock-high 200 \
--stock-low 40

go run main.go recommend delete 1 1 

go run main.go recommend add 4 --book-id 7

go run main.go recommend add 4 --bundle-id 6