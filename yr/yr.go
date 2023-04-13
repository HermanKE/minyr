package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"io"
	"os"
	"bufio"
	"log"
	"math"
	"github.com/HermanKE/funtemps/conv"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFarhenheit(celsiusFloat) //gramatikkfeil pga. funtemps
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}


func CelsiusToFahrenheitLine(line string) (string, error) {
	dividedString := strings.Split(line, ";")
	var err error

	if (len(dividedString) == 4) {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil

	//return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
}


func CountLines(filename string)(int, error) {
	//Åpner filen:
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	//Lager scanner som leser linjene:
	scanner := bufio.NewScanner(file)

	//Loop som går igjennom filen og teller linjer:
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %v", err)
	}
	return lineCount, nil
}


func EditLastLine(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    buf := make([]byte, 3)
    _, err = file.Seek(-3, io.SeekEnd)
    if err != nil {
        return err
    }
    _, err = file.Read(buf)
    if err != nil {
        return err
    }
    if string(buf) != ";;;" {
        return errors.New("last line doesn't end with ';;;'")
    }

    _, err = file.Seek(-2, io.SeekEnd)
    if err != nil {
        return err
    }
    _, err = file.Write([]byte("endringen er gjort av Herman Erlingsen"))
    if err != nil {
        return err
    }

	return nil
}


func CalculateAverageFourthElement(filePath string) (float64, error) {
	// Åpner CSV-filen:
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	//Lager scanner som leser linjene:
	scanner := bufio.NewScanner(file)

	// Variabler for summen og antallet av temperaturene:
	sum := 0.0
	count := 0

	//Loop som går igjennom linjene i filen:
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if lineNumber == 1 || lineNumber == 16756 {
			continue
		}

		// Deler linjene inn i deler, som sjekker om de har 4 elementer:
		line := scanner.Text()
		fields := strings.Split(line, ";")
		if len(fields) < 4 {
			return 0, fmt.Errorf("line %d has less than 4 fields", lineNumber)
		}

		// Konverterer det fjerde elementet til float og legger det til i summen:
		num, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return 0, fmt.Errorf("error converting field %d in line %d to float: %v", 3, lineNumber, err)
		}
		sum += num
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	// Regner ut gjennomsnittet av alle temperaturene:
	if count == 0 {
		return 0, fmt.Errorf("no valid lines found")
	}
	average := sum / float64(count)
	average = math.Round(average * 100) / 100

	return average, nil
}


func ConvertCelsiusFileToFahrenheitFile() {
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	lineNumber := 0
	scanner := bufio.NewScanner(bufio.NewReader(src))
	writer := bufio.NewWriter(dest)
	
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if lineNumber == 1 {
			_, err = writer.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		if lineNumber == 27 {
			_, err = writer.WriteString(line)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		newLine, err := CelsiusToFahrenheitLine(line)
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


func ReadLastLine(filePath string)(string, error) {
	// Åpner CSV-filen:
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	//Lager scanner som leser linjene:
	scanner := bufio.NewScanner(file)

	// Variabel som lagrer den siste linjen:
	var lastLine string

	// Loop som går igjennom linjen i filen og lagrer den siste linejn den leser:
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	// Sjekker om det ble en error under scanningen:
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Sjekker at den siste linjen inneholder forventet string:
	expectedString := "Data er gyldig per 20.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Herman Erlingsen"
	if strings.Contains(lastLine, expectedString) {
		return lastLine, nil
	}

	// Returnerer en error om siste linje ikke inneholder det den skal:
	return "", fmt.Errorf("last line does not contain the expected string")
}