package service

import "sort"

func NewCalculator() *Calculator {
	return &Calculator{}
}

type Calculator struct{}

func (c *Calculator) OrderArray(input []int) []int {

	output := make([]int, len(input))
	copy(output, input)

	//sort array
	sort.Ints(output)

	//move repeated elements to the end of the array
	read := 0
	write := 0

	for read < len(output) {

		// Swap the values pointed at by read and write.
		pointerWrite := output[write]
		output[write], output[read] = output[read], output[write]

		/*
			Advance the read pointer forward to the next unique value.  Since we
			moved the unique value to the write location, we compare values
			against input[write] instead of input[read].
		*/
		for read < len(output) && (output[read] == output[write] || output[read] == pointerWrite) {
			read++
		}

		write++
	}

	return output
}
