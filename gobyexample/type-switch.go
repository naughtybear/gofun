package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	MA := []byte(`{"name": "Massachusetts", "area": 27336, "water": 25.7, "senators":["John Kerry", "Scott Brown"]}`)

	fmt.Println(string(MA))

	var object interface{}
	if err := json.Unmarshal(MA, &object); err != nil {
		log.Fatal(err)
	} else {
		jsonObject := object.(map[string]interface{})
		// fmt.Printf("%T", jsonObject)
		// fmt.Println(jsonObject)
		fmt.Println(jsonObjectAsString(jsonObject))
	}
}

func jsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	comma := ""
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) {
		case nil:
			fmt.Fprintf(&buffer, "%q: null", key)
		case bool:
			fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			fmt.Fprintf(&buffer, "%q: %f", key, value)
		case string:
			fmt.Fprintf(&buffer, "%q: %s", key, value)
		case []interface{}:
			fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, s := range value {
				if s, ok := s.(string); ok {
					fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ","
				}
			}
			buffer.WriteString("]")
		}
		comma = ","
	}
	buffer.WriteString("}")
	return buffer.String()
}

type State struct {
	Name     string
	Senators []string
}
