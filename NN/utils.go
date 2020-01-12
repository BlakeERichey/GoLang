package network

//ValidationSplit returns input, targets, validInputs, validTargets
//splitting with perc% of the values being present in the validation arrays
func ValidationSplit(inputs, targets [][]float64, perc float64) (in, tar, vI, vT [][]float64) {
	validStart := int((1 - perc) * float64(len(inputs)))
	return inputs[:validStart], targets[:validStart], inputs[validStart:], targets[validStart:]
}

//minMax returns the min and max of an array
func minMax(array ...float64) (float64, float64) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func sumArr(array ...float64) (total float64) {
	for _, val := range array {
		total += val
	}
	return total
}

func Argmax(arr ...float64) (index int) {
	index = 0
	maxVal := arr[0]
	for i, val := range arr {
		if val > maxVal {
			index = i
			maxVal = val
		}
	}
	return index
}
