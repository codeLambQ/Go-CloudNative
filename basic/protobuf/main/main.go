package main

import (
	"encoding/json"
	"fmt"
	protobuf "portobuf"

	"google.golang.org/protobuf/proto"
)

func main() {
	lesson := &protobuf.Lession{}
	lesson.Id = 1
	lesson.Name = "go"
	lesson.Rating = 4.5
	bytes, _ := proto.Marshal(lesson)
	fmt.Println(len(bytes))
	jsonBytes, _ := json.Marshal(lesson)
	fmt.Println(len(jsonBytes))
}
