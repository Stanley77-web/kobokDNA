package ValidatorInput

import (
	"errors"
	"fmt"
	"regexp"
)

func DNAValidator(DNApengguna string) error {
	// errorMessage := ""
	// errorCount := 0
	// reg1, _ := regexp.Compile("[a-z]")
	// if reg1.MatchString(DNApengguna) {
	// 	errorMessage += "not capital letter"
	// 	errorCount++
	// }

	// reg2, _ := regexp.Compile("[1234567890]+")

	// if reg2.MatchString(DNApengguna) {
	// 	if errorCount > 0 {
	// 		errorMessage += ", "
	// 	}
	// 	errorCount++
	// 	errorMessage += "contains number"
	// 	// c.Error(err)
	// }

	// reg3, _ := regexp.Compile("[^A-Z]")

	// if reg3.MatchString(DNApengguna) {
	// 	if errorCount > 0 {
	// 		errorMessage += ", "
	// 	}
	// 	errorCount++
	// 	errorMessage += "contains illegal character"
	// 	// c.Error(err)
	// }

	reg1, _ := regexp.Compile("[^ACGT]")
	if reg1.MatchString(DNApengguna) {
		return errors.New("DNA sequence contains illegal character")
		// if errorCount > 0 {
		// 	errorMessage += ", "
		// }
		// errorCount++
		// errorMessage += "not in (A, C, G, T)"

	}
	return nil
	// if errorCount == 0 {
	// 	return errorMessage, false
	// } else {
	// 	return "DNA not valid because " + errorMessage, true
	// }
}

func DateValidator(date string) (int, error) {
	bulan := "(Januari|Februari|Maret|April|Mei|Juni|Juli|Agustus|September|Oktober|November|Desember)"

	regexp1, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	if regexp1.MatchString(date) {
		return 1, nil
	}

	regexp2, _ := regexp.Compile("(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])/((19|20)\\d\\d)")

	if regexp2.MatchString(date) {
		return 2, nil
	}

	regexp3, _ := regexp.Compile("((19|20)\\d\\d)/(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])")

	if regexp3.MatchString(date) {
		return 3, nil
	}

	regexp4, _ := regexp.Compile("((19|20)\\d\\d)/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])")

	if regexp4.MatchString(date) {
		return 4, nil
	}

	regexp5, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")

	if regexp5.MatchString(date) {
		return 5, nil
	}

	regexp6, _ := regexp.Compile("(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])-((19|20)\\d\\d)")

	if regexp6.MatchString(date) {
		return 6, nil
	}

	regexp7, _ := regexp.Compile("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])")

	if regexp7.MatchString(date) {
		return 7, nil
	}

	regexp8, _ := regexp.Compile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")

	if regexp8.MatchString(date) {
		return 8, nil
	}

	regexp9, _ := regexp.Compile("(0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012]) ((19|20)\\d\\d)")

	if regexp9.MatchString(date) {
		return 9, nil
	}

	regexp10, _ := regexp.Compile("(0?[1-9]|1[012]) (0?[1-9]|[12][0-9]|3[01]) ((19|20)\\d\\d)")

	if regexp10.MatchString(date) {
		return 10, nil
	}

	regexp11, _ := regexp.Compile("((19|20)\\d\\d) (0?[1-9]|[12][0-9]|3[01]) (0?[1-9]|1[012])")

	if regexp11.MatchString(date) {
		return 11, nil
	}

	regexp12, _ := regexp.Compile("((19|20)\\d\\d) (0?[1-9]|1[012]) (0?[1-9]|[12][0-9]|3[01])")

	if regexp12.MatchString(date) {
		return 12, nil
	}

	regexp13, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])/%s/((19|20)\\d\\d)", bulan))

	if regexp13.MatchString(date) {
		return 13, nil
	}

	regexp14, _ := regexp.Compile(fmt.Sprintf("%s/(0?[1-9]|[12][0-9]|3[01])/((19|20)\\d\\d)", bulan))

	if regexp14.MatchString(date) {
		return 14, nil
	}

	regexp15, _ := regexp.Compile(fmt.Sprintf("((19|20)\\d\\d)/(0?[1-9]|[12][0-9]|3[01])/%s", bulan))

	if regexp15.MatchString(date) {
		return 15, nil
	}

	regexp16, _ := regexp.Compile(fmt.Sprintf("((19|20)\\d\\d)/%s/(0?[1-9]|[12][0-9]|3[01])", bulan))

	if regexp16.MatchString(date) {
		return 16, nil
	}

	regexp17, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01])-%s-((19|20)\\d\\d)", bulan))

	if regexp17.MatchString(date) {
		return 17, nil
	}

	regexp18, _ := regexp.Compile(fmt.Sprintf("%s-(0?[1-9]|[12][0-9]|3[01])-((19|20)\\d\\d)", bulan))

	if regexp18.MatchString(date) {
		return 18, nil
	}

	regexp19, _ := regexp.Compile(fmt.Sprintf("((19|20)\\d\\d)-(0?[1-9]|[12][0-9]|3[01])-%s", bulan))

	if regexp19.MatchString(date) {
		return 19, nil
	}

	regexp20, _ := regexp.Compile(fmt.Sprintf("((19|20)\\d\\d)-%s-(0?[1-9]|[12][0-9]|3[01])", bulan))

	if regexp20.MatchString(date) {
		return 20, nil
	}

	regexp21, _ := regexp.Compile(fmt.Sprintf("(0?[1-9]|[12][0-9]|3[01]) %s ((19|20)\\d\\d)", bulan))

	if regexp21.MatchString(date) {
		return 21, nil
	}

	regexp22, _ := regexp.Compile(fmt.Sprintf("%s (0?[1-9]|[12][0-9]|3[01]) ((19|20)\\d\\d)", bulan))

	if regexp22.MatchString(date) {
		return 22, nil
	}

	regexp23, _ := regexp.Compile(fmt.Sprintf("((19|20)\\d\\d) (0?[1-9]|[12][0-9]|3[01]) %s", bulan))

	if regexp23.MatchString(date) {
		return 23, nil
	}

	regexp24, _ := regexp.Compile(fmt.Sprintf("((19|20)\\d\\d) %s (0?[1-9]|[12][0-9]|3[01])", bulan))

	if regexp24.MatchString(date) {
		return 24, nil
	}

	return 0, errors.New("Date format is not valid")
}
