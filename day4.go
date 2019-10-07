package main

import (
	"fmt"
	"io"
	"os"
)

func open_file(dstname, srcfile string) ( written int64, err error){
	src, err := os.Open(srcfile)
	//defer src.Close()  // if file is not present then use this
	defer func() {
		fmt.Println("Recovered from Panic")
		recover()
	}()
	if err != nil {
		//return fmt.Println("Error: while running file") // give type error
		//return fmt.Errorf( "Error : while running file")
		panic("Error in file open")
	}

	dst, err := os.Create(dstname)
	if err != nil{
		return
	}

	written, err = io.Copy(dst,src)
	//src.Close()
	dst.Close()
	return
}
