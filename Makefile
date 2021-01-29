all:
	go build -o example cmd/example.go
	go build -o testrun cmd/test.go
run:
	sudo ./example
test:
	sudo ./testrun
clean:
	rm example

ctest:
	gcc -I/usr/include/glib-2.0 -I/usr/lib64/glib-2.0/include \
	-lblockdev -lglib-2.0 \
	-o test ./cmd/test.c && ./test

gtest:
	GODEBUG=cgocheck=0 go run ./cmd/main.go
