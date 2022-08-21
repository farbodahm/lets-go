package main

func Hack(start string) []string {
	finalResult := []string{}

	tempResult := Retrieve(start)
	finalResult = append(finalResult, tempResult...)

	for _, v := range tempResult {
		finalResult = append(finalResult, Hack(v)...)
	}

	return finalResult
}
