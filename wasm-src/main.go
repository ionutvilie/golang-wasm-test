package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	yaml "gopkg.in/yaml.v2"
)

func decode(request []byte, inType string) (*map[string]interface{}, error) {

	var data map[string]interface{}

	switch inType {
	case "json":
		err := json.Unmarshal(request, &data) // handle error
		if err != nil {
			return nil, err
		}
		return &data, nil
	case "yaml":
		fmt.Println("yaml")
		return nil, nil
	case "xml":
		fmt.Println("xml")
		return nil, nil
	case "toml":
		fmt.Println("toml")
		return nil, nil
	default:
		return nil, nil
	}
}

func encode(data *map[string]interface{}, inType string) ([]byte, error) {

	switch inType {
	case "json":
		out, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return out, nil
	case "yaml":
		out, err := yaml.Marshal(data)
		if err != nil {
			return nil, err
		}
		return out, nil
	case "xml":
		fmt.Println("toml")
		return nil, nil
	case "toml":
		fmt.Println("toml")
		return nil, nil
	default:
		return nil, nil
	}
}

func convert(args []js.Value) {

	// requestType := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value").String()

	reqType := args[0].String()
	resType := args[1].String()
	fmt.Println("Req Type: " + reqType)
	fmt.Println("Res Type: " + resType)
	request := js.Global().Get("document").Call("getElementById", "request").Get("value").String()

	// fmt.Println(request)

	data, err := decode([]byte(request), reqType)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fmt.Sprintf("%#v", *data))

	response, err := encode(data, resType)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(fmt.Sprintf("%#v", response))

	js.Global().Get("document").Call("getElementById", "response").Set("value", string(response))
}

func registerCallbacks() {
	js.Global().Set("convert", js.NewCallback(convert))
}

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly Initialized")
	registerCallbacks()

	<-c
}
