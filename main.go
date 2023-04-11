package main

import (
	//"fmt"
	"os"
	"log"
	//"io"
	//"strings"
	"bufio"
	"github.com/HermanKE/minyr/yr"
)
func main() {
	src, err := os.Open("table.csv")
	//src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()


	scanner := bufio.NewScanner(bufio.NewReader(src))
	writer := bufio.NewWriter(dest)
	
	for scanner.Scan() {
		line := scanner.Text()
		newLine, err := yr.CelsiusToFahrenheitLine(line)
		_, err = writer.WriteString(newLine + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
