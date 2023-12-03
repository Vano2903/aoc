package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readFile(path string) []byte {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return c
}

func traverseArray(data []interface{}, spacing string) int {
	sum := 0
	for _, v := range data {
		switch vv := v.(type) {
		case string:
			fmt.Println(spacing, "is string", vv)
		case float64:
			fmt.Println(spacing, "is int", vv)
			sum += int(vv)
		case map[string]interface{}:
			fmt.Println(spacing, "is an object:")
			sum += traverse(vv, spacing+"  ")

		case []interface{}:
			fmt.Println(spacing, "is an array:")
			sum += traverseArray(vv, spacing+"  ")
		default:
			fmt.Println(spacing, "is of a type I don't know how to handle")
			// fmt.Printf(spacing+"%T\n", vv)

		}
	}
	return sum
}

func traverse(data map[string]interface{}, spacing string) int {
	sum := 0
	for k, v := range data {
		// fmt.Println(k, v)
		switch vv := v.(type) {
		case string:
			fmt.Println(spacing, k, "is string", vv)
		case float64:
			fmt.Println(spacing, k, "is int", vv)
			sum += int(vv)
		case map[string]interface{}:
			fmt.Println(spacing, k, "is an object:")
			sum += traverse(vv, spacing+"  ")

		case []interface{}:
			fmt.Println(spacing, k, "is an array:")
			sum += traverseArray(vv, spacing+"  ")
		default:
			fmt.Println(spacing, k, "is of a type I don't know how to handle")
			// fmt.Println(spacing, vv)
		}
	}
	return sum
}

func solve1(c []byte) {
	data := make(map[string]interface{})
	json.Unmarshal(c, &data)
	fmt.Println(data)
	sum := traverse(data, "")
	fmt.Println("sum", sum)

}

func traverseArray2(data []interface{}, spacing string, isInObject bool) int {
	sum := 0
	for _, v := range data {
		switch vv := v.(type) {
		case string:
			fmt.Println(spacing, "is string", vv)
		case float64:
			fmt.Println(spacing, "is int", vv)
			sum += int(vv)
		case map[string]interface{}:
			fmt.Println(spacing, "is an object:")
			sum += traverse2(vv, spacing+"  ", true)

		case []interface{}:
			fmt.Println(spacing, "is an array:")
			sum += traverseArray2(vv, spacing+"  ", false)
		default:
			fmt.Println(spacing, "is of a type I don't know how to handle")
			// fmt.Printf(spacing+"%T\n", vv)

		}
	}
	return sum
}

func traverse2(data map[string]interface{}, spacing string, isInObject bool) int {
	sum := 0
	for k, v := range data {
		// fmt.Println(k, v)
		switch vv := v.(type) {
		case string:
			fmt.Println(spacing, k, "is string", vv)
			if vv == "red" && isInObject {
				return 0
			}
		case float64:
			fmt.Println(spacing, k, "is int", vv)
			sum += int(vv)
		case map[string]interface{}:
			fmt.Println(spacing, k, "is an object:")
			sum += traverse2(vv, spacing+"  ", true)

		case []interface{}:
			fmt.Println(spacing, k, "is an array:")
			sum += traverseArray2(vv, spacing+"  ", false)
		default:
			fmt.Println(spacing, k, "is of a type I don't know how to handle")
			// fmt.Println(spacing, vv)
		}
	}
	return sum
}
func solve2(c []byte) {
	data := make(map[string]interface{})
	json.Unmarshal(c, &data)
	fmt.Println(data)
	sum := traverse2(data, "", true)
	fmt.Println("sum", sum)

}

func main() {
	c := readFile("input.json")
	solve1(c)
	solve2(c)
}
