package main

import (
	"os"
	"log"
	"bufio"
	//"io"
	//"strings"
	//"github.com/HermanKE/funtemps/conv"
)
func main() {
	src, err := os.Open("table.csv")
	//src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	scanner := bufio.NewScanner(bufio.NewReader(src))
	writer := bufio.NewWriter(dest)

	for scanner.Scan() {
		line := scanner.Text()
		line = line + "\n"
		writer.Write([]byte(line))
	}
	writer.Flush()

/*
	var buffer []byte
	var linebuf []byte //nil
	buffer = make([]byte, 1)
	bytesCount := 0

	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		//log.Println("%c", buffer[:n])
		if buffer[0] == 0x0A {
			log.Println(string(linebuf))
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				fahr := conv.CelsiusToFahrenheit(celsius)
				log.Println(elementArray[3])
			}
			linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}
		//log.Println(string(linebuf))
		if err == io.EOF {
			break
		}
	}
*/
}
