package games

import (
	"math/rand"
	"strconv"
)

func SimpleEquation() GameOutput {

	min := 2
	max := 10

	result := rand.Intn(max-min) + min
	firstNumber := rand.Intn(result-1) + 1
	secondNumber := result - firstNumber

	task := createTask(firstNumber, secondNumber, result)
	resultSet := createResultSet(min, max, result)

	var output = GameOutput{
		Task: task,
		Result: Element{
			Type:  "string",
			Value: strconv.Itoa(result),
		},
		ResultSet: resultSet,
	}

	return output
}

func createTask(firstNumber int, secondNumber int, result int) ElementsSet {
	var task ElementsSet
	task.Type = "equation"

	var leftSide ElementsSet
	leftSide.Type = "expression"

	n := 1
	for n < firstNumber+1 {
		var element = Element{
			Type:  "picture",
			Value: "1",
		}
		leftSide.Elements = append(leftSide.Elements, element)
		n *= 2
	}

	leftSide.Elements = append(leftSide.Elements, Element{
		Type:  "string",
		Value: "+",
	})

	m := 1
	for m < secondNumber+1 {
		var element = Element{
			Type:  "picture",
			Value: "1",
		}
		leftSide.Elements = append(leftSide.Elements, element)
		m *= 2
	}

	var rightSide Element
	rightSide.Type = "string"
	rightSide.Value = "?"

	task.Elements = append(task.Elements, leftSide, rightSide)

	return task
}

func createResultSet(min int, max int, result int) ElementsSet {
	var resultSet ElementsSet
	resultSet.Type = "set"

	indexOfResult := rand.Intn(4)

	var gameRange []int
	n := 0
	for n < max-min {
		gameRange = append(gameRange, n+1)
	}

	for i := range gameRange {
		j := rand.Intn(i + 1)
		gameRange[i], gameRange[j] = gameRange[j], gameRange[i]
	}

	k := 0
	for k < 5 {
		if k == indexOfResult {
			var element = Element{
				Type:  "string",
				Value: strconv.Itoa(result),
			}
			resultSet.Elements = append(resultSet.Elements, element)
		} else {
			var element = Element{
				Type:  "string",
				Value: strconv.Itoa(gameRange[k] + min),
			}
			resultSet.Elements = append(resultSet.Elements, element)
		}
		k += 1
	}
	return resultSet
}
