package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	DataDir   = "data"
	MapFile   = "data/wikimap.json"
	GraphFile = "data/wikigraph.json"
)

type MapData struct {
	IdToTitle map[int]string `json:"id_to_title"`
	TitleToID map[string]int `json:"title_to_id"`
}

func init() {
	if err := os.MkdirAll(DataDir, 0755); err != nil {
		panic(err)
	}
}

func SaveMap(idToTitle map[int]string, titleToID map[string]int) error {
	data := MapData{
		IdToTitle: idToTitle,
		TitleToID: titleToID,
	}

	file, err := os.Create(MapFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func LoadMap() (map[int]string, map[string]int, error) {
	file, err := os.Open(MapFile)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var data MapData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, nil, err
	}

	return data.IdToTitle, data.TitleToID, nil
}

func SaveGraph(graph map[int][]int) error {
	file, err := os.Create(GraphFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(graph)
}

func SaveGraphStreaming(graph map[int][]int) error {

	fmt.Println("Saving graph...")

	file, err := os.Create(GraphFile)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("{\n")
	first := true
	for k, v := range graph {
		if !first {
			file.WriteString(",\n")
		}
		first = false
		valJson, _ := json.Marshal(v)
		file.WriteString(fmt.Sprintf("\"%d\":%s", k, valJson))
	}
	file.WriteString("\n}\n")
	return nil
}

func LoadGraph() (map[int][]int, error) {
	file, err := os.Open(GraphFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var graph map[int][]int
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&graph); err != nil {
		return nil, err
	}

	return graph, nil
}
