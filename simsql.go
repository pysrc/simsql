package simsql

import (
	"fmt"
	"strings"
)

func GetKeysValues(m map[string]interface{}) ([]string, []interface{}) {
	var keys = make([]string, len(m))
	var values = make([]interface{}, len(m))
	var it = 0
	for k, v := range m {
		keys[it] = k
		values[it] = v
		it++
	}
	return keys, values
}

type SqlNode struct {
	Query string
	Args  []interface{}
}

func Insert(table string, data map[string]interface{}) SqlNode {
	keys, values := GetKeysValues(data)
	pre := make([]string, len(keys))
	for i := 0; i < len(pre); i++ {
		pre[i] = "?"
	}
	return SqlNode{fmt.Sprintf(`insert into %s(%s) values(%s)`, table, strings.Join(keys, ","), strings.Join(pre, ",")), values}
}

// 生成插入语句，多条回调
func Inserts(table string, data []map[string]interface{}) []SqlNode {
	var res = make([]SqlNode, len(data))
	for i, v := range data {
		res[i] = Insert(table, v)
	}
	return res
}

func Query(op string, schema []string, table string, where map[string]interface{}, per, page int64) SqlNode {
	keys, values := GetKeysValues(where)
	pre := make([]string, len(keys))
	for i := 0; i < len(pre); i++ {
		if i == 0 {
			pre[i] = "where " + keys[i] + "=?"
			continue
		}
		pre[i] = keys[i] + "=?"
	}
	values = append(values, per*(page-1))
	values = append(values, per)
	return SqlNode{fmt.Sprintf("select %s from %s %s limit ?, ?", strings.Join(schema, ","), table, strings.Join(pre, " "+op+" ")), values}
}

// 修改
func Update(keyop string, table string, data map[string]interface{}, where map[string]interface{}) SqlNode {
	k, v := GetKeysValues(data)
	pre1 := make([]string, len(k))
	for i := 0; i < len(pre1); i++ {
		pre1[i] = k[i] + "=?"
	}
	k2, v2 := GetKeysValues(where)
	pre2 := make([]string, len(k2))
	for i := 0; i < len(pre2); i++ {
		pre2[i] = k2[i] + "=?"
		v = append(v, v2[i])
	}
	return SqlNode{fmt.Sprintf("update %s set %s where %s", table, strings.Join(pre1, ","), strings.Join(pre2, " "+keyop+" ")), v}
}

// 删除
func Delete(keyop string, table string, where map[string]interface{}) SqlNode {
	k, v := GetKeysValues(where)
	pre := make([]string, len(k))
	for i := 0; i < len(pre); i++ {
		pre[i] = k[i] + "=?"
	}
	return SqlNode{fmt.Sprintf("delete from %s where %s", table, strings.Join(pre, " "+keyop+" ")), v}
}
