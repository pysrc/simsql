package main

import (
	"fmt"
	"github.com/pysrc/simsql"
)

func main() {
	fmt.Println(simsql.Insert("student", map[string]interface{}{"Name": "xiaoming", "Pwd": "password"}))
	fmt.Println(simsql.Inserts("student", []map[string]interface{}{map[string]interface{}{"Name": "xianom", "Pwd": "12345"}, map[string]interface{}{"Name": "xianom", "Pwd": "12345"}}))
	fmt.Println(simsql.Query("or", []string{"field_1", "field_2"}, "student", map[string]interface{}{"Id": 56, "Name": "xiaoming"}, 3, 1))
	fmt.Println(simsql.Update("and", "student", map[string]interface{}{"Id": 446, "Name": "小米"}, map[string]interface{}{"Id": 56}))
	fmt.Println(simsql.Delete("and", "student", map[string]interface{}{"Id": 446, "Name": "小米"}))
}
