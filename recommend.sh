#!/bin/sh
bookstore recommend create deep 2 \
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

bookstore recommend delete 1 1 

bookstore recommend add 8 --book-id 21

bookstore recommend add 8 --bundle-id 5

bookstore recommend create test1
