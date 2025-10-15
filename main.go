package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	h "github.com/shark121/go-NN/helpers"
	m "github.com/shark121/go-NN/modules"
	mat "gonum.org/v1/gonum/mat"
)

func main() {
	network := m.CreateNetwork([]int{28 * 28, 16, 16})

	for idx, v := range network.Layers {
		fmt.Println("Neuron", idx+1, len(v.Neurons), v.Level)
	}

	_, paths := h.UbyteReader("./MNISTdataset")

	base := "C:/Users/HP/Desktop/nuclear-launch-codes/go-NN/"

	dataMap := map[int][]byte{}

	for idx, path := range paths {

		fmt.Println(path)

		data, err := os.ReadFile(filepath.Join(base, path))

		if err != nil {
			log.Fatal(err)

			continue
		}

		dataMap[idx] = data
	}

	offset := 10 // because it seemed skewed

	offsetArr := make([]byte, offset)

	dataMap[0] = append(offsetArr, dataMap[0]...)

	tensor := []mat.Dense{}

	for i := 0; i < len(dataMap[0])-784-offset; i += 784 {
		toFloat := []float64{}

		for _, b := range dataMap[0][i : i+784] {
			toFloat = append(toFloat, float64(b))
		}

		tensor = append(tensor, *mat.NewDense(28, 28, toFloat))

	}

	print(len(tensor))
	fmt.Println(reflect.TypeOf(tensor[0]))
	fmt.Println(tensor[20].Dims())
	h.DrawImage(28, 28, tensor[0])
}
