package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFetchAll(t *testing.T) {
	buffer1 := &bytes.Buffer{}
	buffer2 := &bytes.Buffer{}
	buffer3 := &bytes.Buffer{}
	fetchAll([]string{"https://www.ayataka.jp/"}, buffer1)
	fetchAll([]string{"https://www.ayataka.jp/"}, buffer2)
	fetchAll([]string{"https://www.ayataka.jp/"}, buffer3)
	fmt.Println(buffer1)
	fmt.Println(buffer2)
	fmt.Println(buffer3)
}
