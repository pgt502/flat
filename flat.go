package flat

import (
	"log"
	"reflect"
)

type Options struct {
	Delimiter string
	Safe      bool
	Object    bool
	MaxDepth  int
}

func Flatten(nested map[string]interface{}, opts Options) (m map[string]interface{}, err error) {
	m, err = flatten("", nested, opts)
	return
}

func flatten(prefix string, nested interface{}, opts Options) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	switch nested := nested.(type) {
	case map[string]interface{}:
		// map
		for k, v := range nested {
			// create new key
			newKey := k
			if prefix != "" {
				newKey = prefix + "." + k
			}
			switch v.(type) {
			case map[string]interface{}, []interface{}:
				temp, fErr := flatten(newKey, v, opts)
				if fErr != nil {
					err = fErr
					return
				}
				// empty map {}
				if reflect.DeepEqual(temp, map[string]interface{}{}) {
					m[newKey] = temp
				} else {
					// merge temp to m
					for kt, vt := range temp {
						m[kt] = vt
					}
				}
			default:
				m[newKey] = v
			}
		}
	case []interface{}:
		log.Println("slice")
		// slice
	default:
		log.Println("error")
	}
	return
}
