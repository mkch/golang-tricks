package main

/*

#include <stdlib.h>
#include <stdio.h>

char array[] = {1,2,3};

int buf_len = 2;
char* alloc_buf(int size) {
	return (char*)malloc(size);
}
void print_buf(char* buf, int size) {
	printf("print_buf: ");
	for(int i = 0; i < size; i++) {
		printf("%c ", buf[i]);
	}
	printf("\n");
}

*/
import "C"
import (
	"fmt"
	"io"
	"math"
	"os"
	"unsafe"
)

func main() {
	var bufSize = 2
	var buf = C.alloc_buf(C.int(bufSize))
	defer C.free(unsafe.Pointer(buf))
	// Wraps the C buf in a go slice.
	// CAUTION: this slice can't be used after buf is freed.
	var sliceOfCBuf = (*[math.MaxUint32]byte)(unsafe.Pointer(buf))[:bufSize]

	sliceOfCBuf[0] = 'a'
	sliceOfCBuf[1] = 'b'
	fmt.Println(sliceOfCBuf)
	C.print_buf(buf, C.int(bufSize))

	copy(sliceOfCBuf, "cd")
	C.print_buf(buf, C.int(bufSize))

	fmt.Println("Please enter something:")
	_, err :=io.ReadFull(os.Stdin, sliceOfCBuf)
	if err != nil {
		panic(err)
	}
	C.print_buf(buf, C.int(bufSize))
}
