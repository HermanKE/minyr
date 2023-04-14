package main

import (
	"fmt"
	"os"
	//"log"
	//"io"
	//"strings"
	"bufio"
	"errors"
	"github.com/HermanKE/minyr/yr"
)
func main() {
	fmt.Println("Minyr er åpnet. Velg mellom 'exit/q' for å stoppe, 'convert' for å konvertere, 'average' for å finne gjennomsnittet:")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
	
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")

			if _, err := os.Stat("./kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
				fmt.Println("Filen eksisterer i systemet")
				fmt.Println("Velg mellom: 'q/exit' for å gå ut, 'j' for å generere og konvertere ny fil, 'n' for å konvertere eksisterende fil")
				var inputConv string
				scannerConv := bufio.NewScanner(os.Stdin)
				for scannerConv.Scan() {
					inputConv = scannerConv.Text()

					if inputConv == "q" || inputConv == "exit" {
							fmt.Println("exit")
							os.Exit(0)
					} else if inputConv == "j" {
						os.Remove("kjevik-temp-fahr-20220318-20230318.csv")
						yr.ConvertCelsiusFileToFahrenheitFile()
						yr.EditLastLine("kjevik-temp-fahr-20220318-20230318.csv")
					} else if inputConv == "n" {
						yr.ConvertCelsiusFileToFahrenheitFile()
						yr.EditLastLine("kjevik-temp-fahr-20220318-20230318.csv")
					} else {
						fmt.Println("Venligst velg mellom 'j' for å genere en ny fil eller 'n' for å beholde eksisterende")
					}
				}
			} else if errors.Is(err, os.ErrNotExist) {
				yr.ConvertCelsiusFileToFahrenheitFile()
				yr.EditLastLine("kjevik-temp-fahr-20220318-20230318.csv")
				fmt.Println("Nå har du konvertert fra celsius til fahrenheit")
			}
		} else if input == "average" {
			fmt.Println("Velg mellom: 'q/exit' for å gå ut, 'c' for å få gjennomsnitt i celsius eller 'f' for å få gjennomsnitt i fahrenheit")
			var inputAvg string
			scannerAvg := bufio.NewScanner(os.Stdin)
			for scannerAvg.Scan() {
				inputAvg = scannerAvg.Text()
				
				if inputAvg == "q" || inputAvg == "exit" {
					fmt.Println("exit")
					os.Exit(0)
				} else if inputAvg == "c" {
					avg, err := yr.CalculateAverageFourthElement("kjevik-temp-celsius-20220318-20230318.csv")
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						return
					}
					fmt.Printf("Average of fourth elements: %v\n", avg)

				} else if inputAvg == "f" {
					avg, err := yr.CalculateAverageFourthElement("kjevik-temp-fahr-20220318-20230318.csv")
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						return
					}
					fmt.Printf("Average of fourth elements: %v\n", avg)
				} else if inputAvg != "c" && inputAvg != "f"{
					fmt.Println("Venligst velg mellom 'c' eller 'f'")
				}
			}
		} else {
			fmt.Println("Venligst velg convert, average eller exit:")
		}
	}
}
