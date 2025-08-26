package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
)

func readFromReader(r io.Reader) {
	buf:= make([]byte,1024)
	n,err:= r.Read(buf)
	if err!=nil{
		log.Fatal("🔴ERROR reading from READER:",err)
	}
	fmt.Println("Data:",string(buf[:n])) // not return empty buffers
}

func writeToWriter(w io.Writer, data string){
	_,err:= w.Write([]byte(data)) // We don't need the data (Optnl.)
	if err!=nil{
		log.Fatal("🔴ERROR writing to WRITER:",err)
	}
}

//⚠️ close the resource/file - if we create/open/read
func closeResource(c io.Closer){
	err:=c.Close()
	if err!=nil{
		log.Fatal("🔴ERROR closing the resource:",err)
	}
}

func bufferExample(){
	var buf bytes.Buffer // buff is a value-type (memory on stack)
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample(){
	r1:= strings.NewReader("Hey ")
	r2:= strings.NewReader("World")
	mr:= io.MultiReader(r1,r2)
	
	// create buffer
	buf := new(bytes.Buffer) // buff is a *ptr-type (memory in heap)
	_,err:=buf.ReadFrom(mr)
	if err!=nil{
		log.Fatal("🔴ERROR reading from READER:",err)
	}
	fmt.Println("Data/Bufs:",buf.String())
}

// pipes
// Connecting 2 peices of code so that once can travel from the other like a tunnel
// r - from where it will be read
// w - to where data will be written

func pipeExample(){
	pr,pw:=io.Pipe()
	// goroutine - extracting the fx out of the main thread and executed seperately, anonymously
	go func(){
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	// create buffer
	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string){}


func main() {

	fmt.Println("====== READ FROM READER =======")
	readFromReader(strings.NewReader("random str. to be read"))

	fmt.Println("\n====== WRITE TO WRITER =======")
	// create writer first
	var writer bytes.Buffer
	writeToWriter(&writer, "random str. to be written to WRITER")
	fmt.Println(writer.String())

	fmt.Println("\n====== BUFFER EXAMPLE =======")
	bufferExample()

	fmt.Println("\n====== MULTI-READER EXAMPLE =======")
	multiReaderExample()

	fmt.Println("\n====== READING/WRITING PIPE =======")
	pipeExample()



}