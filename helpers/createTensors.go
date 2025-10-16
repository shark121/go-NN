package helpers

import (
	mat "gonum.org/v1/gonum/mat"
)

func CreateTensors(tensorData []byte, offset int) []mat.Dense {

	// because it seemed skewed

	offsetArr := make([]byte, offset)

	tensorData = append(offsetArr, tensorData...)

	tensor := []mat.Dense{}

	for i := 0; i < len(tensorData)-784-offset; i += 784 {
		toFloat := []float64{}

		for _, b := range tensorData[i : i+784] {
			toFloat = append(toFloat, float64(b))
		}

		tensor = append(tensor, *mat.NewDense(28, 28, toFloat))

	}

	return tensor
}
