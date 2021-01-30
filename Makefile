ctest:
	gcc `pkg-config --cflags glib-2.0` \
	./cmd/test.c -lblockdev -lglib-2.0 \
	-o test && ./test

gtest:
	GODEBUG=cgocheck=0 go run ./cmd/main.go
