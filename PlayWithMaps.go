package main

import (
	"fmt"
	"strconv"
)

func wordFrquencyCounter(word []string) map[string]int {
	mapper := make(map[string]int)
	for _, val := range word {
		mapper[val] = mapper[val] + 1

	}
	return mapper
}

func findStudent(scores map[string]int, name string) string {
	value, exists := scores[name]
	if exists {
		return name + " scored " + strconv.Itoa(value)
	} else {
		return name + " not found"
	}
}
func main() {
	fmt.Println("welcome to the maps demo..!")

	//words frequency counter
	words := []string{"go", "go", "c", "c", "java"}
	fmt.Println(words)
	mapRes := wordFrquencyCounter(words)
	fmt.Println(mapRes)

	fmt.Println("-------------------------------")
	studentsDatabase := make(map[string]int)
	studentsDatabase["vishal"] = 100
	studentsDatabase["delson"] = 90
	studentsDatabase["rohan"] = 70
	studentsDatabase["narayan"] = 85
	studentsDatabase["subbu"] = 35

	fmt.Println(studentsDatabase)
	fmt.Println(findStudent(studentsDatabase, "vishal"))
	fmt.Println(findStudent(studentsDatabase, "amogh"))
	fmt.Println(findStudent(studentsDatabase, "subbu"))
	fmt.Println(findStudent(studentsDatabase, "panav"))

}
