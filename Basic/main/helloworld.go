package main

import (
	"fmt"
	"log"
	"os"
	"time"
)
func init()  {
log.SetOutput(os.Stdout)
}
func main(){
	var abc string = "Runoob"
	fmt.Println("hello world"+abc)
	time.Sleep(time.Second)
}