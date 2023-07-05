package main

import (
	"encoding/json"
	"os"
)

type std struct {
	LastName, FirstName, MiddleName, Birthday, Address, Phone string
	Rating                                                    []int
}

type rtg struct {
	ID       int
	Number   string
	Year     int
	Students []std
}

type avg struct {
	Average float64 `json:"Average"`
}

func avgrating(r *rtg) {
	var val float64
	for _, s := range r.Students {
		val += float64(len(s.Rating))
	}
	var a = map[string]float64{
		"Average": val / float64(len(r.Students)),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(&a)
}

func main() {
	var r rtg
	dec := json.NewDecoder(os.Stdin)
	dec.Decode(&r)
	avgrating(&r)
}
