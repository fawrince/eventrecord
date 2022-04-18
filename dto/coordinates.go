package dto

import "encoding/json"

type Coordinates struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Client string  `json:"client"`

	encoded []byte
}

func (coord Coordinates) ensureEncoded() {
	if coord.encoded == nil {
		coord.encoded, _ = json.Marshal(coord)
	}
}

func (coord Coordinates) Length() int {
	coord.ensureEncoded()
	return len(coord.encoded)
}

func (coord Coordinates) Encode() ([]byte, error) {
	coord.ensureEncoded()
	return coord.encoded, nil
}

func (coord *Coordinates) Decode(data []byte) {
	json.Unmarshal(data, coord)
}
