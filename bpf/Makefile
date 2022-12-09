TARGET := xdp_main.o

.PHONY: all
all: build

build: $(TARGET)

$(TARGET): xdp_main.c
		clang -I../includes -O2 -target bpf -c $^ -o $@

clean:
	rm $(TARGET)