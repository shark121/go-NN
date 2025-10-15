package modules

type Network struct {
	Layers []Layer
}

func CreateNetwork(neuronsPerLayer []int) Network {

	layers := []Layer{}

	for idx, v := range neuronsPerLayer {
		layers = append(layers, CreateLayer(int8(idx+1), int64(v)))
	}

	return Network{Layers: layers}
}
