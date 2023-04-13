package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"io"
	"os"
	"bufio"
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

func StudentString(string)(string, error) {
	var err error

	tekst := ("Endringen er gjort av Herman Erlingsen")
	return tekst, err
}

func StudentLine(line string)(string, error) {
	dividedString := strings.Split(line, ";")
	var err error

	if (len(dividedString) == 4) {
		dividedString[3], err = StudentString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke fovantet format")
	}
	return strings.Join(dividedString, ";"), nil
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
		if lineNumber == 1 || lineNumber == 27 {
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

	return average, nil
}