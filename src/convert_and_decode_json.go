/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Manipulação de JSON
 */

package src

import (
	"encoding/json"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Winners int    `json:"-"` // Omite a informação no JSON
}

func ToJson() {
	user := User{"John Doe", 20, 0}

	js, ok := json.Marshal(user)

	if ok != nil {
		panic(ok)
	}

	println(string(js))
}

func ToObject() {
	ctx := `{"name": "John Doe", "age": 20}`

	var user User

	err := json.Unmarshal([]byte(ctx), &user)

	if err != nil {
		panic(err)
	}

	println(user.Name)
	println(user.Age)
}

func ToJsonSendStdout() {
	user := User{"John Doe", 20, 0}

	err := json.NewEncoder(os.Stdout).Encode(user)

	if err != nil {
		panic(err)
	}
}

func MapToJson() {
	userMap := map[string]interface{}{
		"name": "John Doe",
		"age":  20,
	}

	js, ok := json.Marshal(userMap)

	if ok != nil {
		panic(ok)
	}

	println(string(js))
}

func JsonToMap() {
	ctx := `{"name": "John Doe", "age": 20}`

	var userMap map[string]interface{}

	err := json.Unmarshal([]byte(ctx), &userMap)

	if err != nil {
		panic(err)
	}

	println(userMap["name"])
	println(userMap["age"])
}

func JsonToListMap() {
	ctx := `[{"name": "John Doe", "age": 20}, {"name": "John Doe", "age": 20}]`

	var userMap []map[string]interface{}

	err := json.Unmarshal([]byte(ctx), &userMap)

	if err != nil {
		panic(err)
	}

	println(userMap[0]["name"])
	println(userMap[0]["age"])
}

func JsonToListObject() {
	ctx := `[{"name": "John Doe", "age": 20}, {"name": "Margareth Menezes", "age": 21}]`

	var users []*User

	err := json.Unmarshal([]byte(ctx), &users)

	if err != nil {
		panic(err)
	}

	for _, u := range users {
		println(u.Name, u.Age)
	}
}
