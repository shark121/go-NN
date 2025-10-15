package modules

type Neuron struct {
	Value  int64
	Number int64
	Layer  int16
}

type Layer struct {
	Neurons []Neuron
	Level   int8
}

func CreateLayer(level int8, numberOfNeurons int64) Layer {

	ns := []Neuron{}

	for i := range numberOfNeurons {

		ns = append(ns, CreateNeuron(0, int64(i), int16(level)))

	}

	return Layer{ns, level}
}

func CreateNeuron(value int64, number int64, layer int16) Neuron {
	return Neuron{value, number, layer}
}
