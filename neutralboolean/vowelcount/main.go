package main 

import (
    //"bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
	var filename string
	fmt.Print("Input filename: ")
	fmt.Scanln(&filename)

	fmt.Printf("echo: %q\n", filename)
	count := 0

	if filename == "" {
		var input string
		fmt.Print("Input string whose vowels are to be counted: ")
		fmt.Scanln(&input)
		lowered := strings.ToLower(input)

		for i := 0; i < len(input); i++ {
			if isVowel(lowered[i]) {
				count++
			}
		}

		fmt.Printf("There are %d vowels in %q.\n", count, input)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			panic("failed to open file")
		}

		bufferlen := 8
		excerptCount := 0

		fmt.Println("\"")
		
		var buffer []byte
		for err != io.EOF {
			buffer, err = trimmedRead(file, bufferlen)
			if excerptCount <= 250 {
				fmt.Printf("%s", buffer)
				if excerptCount + len(buffer) >= 250 {
					fmt.Println("...\n\"")
				}
				
				excerptCount += len(buffer)
			}

			lowered := strings.ToLower(fmt.Sprintf("%s", buffer))
			for i := 0; i < len(lowered); i++ {
				if isVowel(lowered[i]) {
					count++
				}
			}
		}

		fmt.Printf("There are %d vowels in the file[%q].\n", count, filename)
	}
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

//returns a byte buffer trimmed to the number of elements in the array, up to a max size given as a parameter
func trimmedRead(f *os.File, maxBufferSize int) (b []byte, err error) {
	b = make([]byte, maxBufferSize)
	n, err := f.Read(b)
	
	return b[:n], err
}