package f64

import (
	"encoding/json"
	"strconv"
)

type F64 float64

func (f F64) F() float64 {
	return float64(f)
}

func (f *F64) UnmarshalJSON(data []byte) (err error) {
	var raw string

	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var v float64
	if v, err = strconv.ParseFloat(raw, 64); err != nil {
		return err
	}

	*f = F64(v)

	return nil
}
