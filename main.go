package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	h "github.com/shark121/go-NN/helpers"
	m "github.com/shark121/go-NN/modules"
)

var pathsMap map[string]string = map[string]string{
	"images":       "MNISTdataset/t10k-images.idx3-ubyte",
	"image_labels": "t10k-labels.idx1-ubyte",
	"train":        "train-images.idx3-ubyte",
	"train_labels": "train-labels.idx1-ubyte",
}

func main() {
	network := m.CreateNetwork([]int{28 * 28, 16, 16})

	for idx, v := range network.Layers {
		fmt.Println("Neuron", idx+1, len(v.Neurons), v.Level)
	}

	mainPath := "./MNISTdataset"

	basePath := "C:/Users/HP/Desktop/nuclear-launch-codes/go-NN/"

	_, paths := h.UbyteReader(mainPath)

	dataMap := map[string][]byte{}

	for _, path := range paths {

		fmt.Println(path)

		data, err := os.ReadFile(filepath.Join(basePath, path))

		if err != nil {
			log.Fatal(err)

			continue
		}

		dataMap[filepath.ToSlash(path)] = data
	}

	for k, _ := range dataMap {
		fmt.Println(reflect.TypeOf(k))
	}
	// fmt.Println(dataMap)
	// fmt.Println(dataMap[pathsMap["images"]])

	image_tensor := h.CreateTensors(dataMap[pathsMap["images"]], 10)

	print(len(image_tensor))
	// fmt.Println(reflect.TypeOf(tensor[0]))
	// fmt.Println(tensor[20].Dims())
	h.DrawImage(28, 28, image_tensor[100])
}
