package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "bufio"
)

const (
	SEEK_SET int = 0 // seek relative to the origin of the file
	SEEK_CUR int = 1 // seek relative to the current offset
	SEEK_END int = 2 // seek relative to the end
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func showLast(fileName string, linesToShow int) {
    fmt.Printf("Show last %d lines from %s\n", linesToShow, fileName)
    f, err := os.Open(fileName)
    check(err)
    defer f.Close()
    f.Seek(0, SEEK_END)
    lines := 0
    b := make([]byte, 1)
    var newLine byte = '\n'
    for lines <= linesToShow {
	    _, err = f.Seek(-1, SEEK_CUR)
	    if err != nil {
	        break
	    }
	    _, err = f.Read(b)
	    check(err)
	    if b[0] == newLine {
		lines += 1
	    }
	    _, err = f.Seek(-1, SEEK_CUR)
	    if err != nil {
	        break
	    }
    }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
	fmt.Println(scanner.Text())
    }
}

func main() {

    last := flag.Bool("n", false, "number of lines to show from the end of the file")
    continuous := flag.Bool("f", false, "watch continuously for changes")
    flag.Parse()

    var fileName string
    if *continuous {
	if len(os.Args) < 3 {
	    fmt.Println("Missing file name")
	    os.Exit(1)
        }
	fileName = os.Args[2]
    	fmt.Printf("Show last lines from %s continuously\n", fileName)
    } 
    if *last {
	if len(os.Args) < 4 {
	    fmt.Println("Missing file name")
	    os.Exit(1)
        }
    	lines, err := strconv.Atoi(os.Args[2])
	if err != nil {
	    fmt.Println("The '-n' option requires an integer number of lines")
	}
    	fileName = os.Args[3]
	showLast(fileName, lines)
    }
}
