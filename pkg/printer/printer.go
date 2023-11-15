package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Params struct {
	V    interface{}
	Path string
}

func ToJSON(params Params) {
	// Convert string array to JSON with pretty print
	jsonData, err := json.MarshalIndent(params.V, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	path := params.Path
	if path == "" {
		path = "output.json"
	}

	createDirIfNotExist(path)

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data written to", path)
}

func createDirIfNotExist(path string) {
	// remove filename from path
	path = path[:len(path)-len(path[strings.LastIndex(path, "/"):])]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0775)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	}
}

func Print(s interface{}) {
	currentType := reflect.TypeOf(s)

	fmt.Print("\n")
	if currentType.String() == "[]string" {
		printString(s)
	} else {
		printStruct(s)
	}
	fmt.Print("\n")
}

func printString(s interface{}) {
	val := reflect.ValueOf(s)
	for i := 0; i < val.Len(); i++ {
		fmt.Println(val.Index(i).Interface())
	}
}

func printStruct(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()

		fmt.Printf("%s: ", fieldName)

		kind := reflect.TypeOf(fieldValue).Kind()
		switch kind {
		case reflect.String:
			s := reflect.ValueOf(fieldValue)
			if s.Len() == 0 {
				fmt.Printf("%v\n", "empty")
			} else {
				fmt.Printf("%v\n", fieldValue)
			}

		case reflect.Slice:
			s := reflect.ValueOf(fieldValue)
			length := reflect.ValueOf(fieldValue).Len()
			if length == 0 {
				fmt.Print("[]\n")
			} else {
				fmt.Println()
				for i := 0; i < s.Len(); i++ {
					value := s.Index(i).Interface()
					printStructItem(value)
				}
			}

		case reflect.Map:
			length := reflect.ValueOf(fieldValue).Len()
			if length == 0 {
				fmt.Print("[]\n")
			} else {
				fmt.Println()
				m := reflect.ValueOf(fieldValue)
				for _, key := range m.MapKeys() {
					fmt.Println("  Key:", key, "Value:", m.MapIndex(key))
				}
			}

		default:
			fmt.Print(kind)
		}
	}
}

func printStructItem(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	list := []string{}
	length := val.NumField()
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface().(string)

		val := fieldName + ": " + fieldValue
		if i != length-1 {
			val = val + ","
		}
		list = append(list, val)
	}
	fmt.Println(list)
}
