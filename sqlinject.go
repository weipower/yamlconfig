package yamlconfig

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"
)

func CheckSQLInject(data interface{}) (err error) {
	dataValue := reflect.ValueOf(data)

	t := reflect.TypeOf(data)

	if dataValue.Kind() != reflect.Struct {
		fmt.Println(dataValue.Kind())
		return errors.New("Input is not a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		k := dataValue.Type().Field(i)
		value := dataValue.Field(i).Interface()
		var valueStr string
		switch value.(type) {
		case string:
			valueStr = fmt.Sprintf(`%v`, value)
			if MatchSQLInject(valueStr) {
				return errors.New(k.Name + " 参数错误(" + valueStr + ")")
			}
		default:
			valueStr = fmt.Sprintf("%v", value)
		}
	}

	return nil
}

func CheckAllFieldsAreSet(data interface{}) (err error) {
	dataValue := reflect.ValueOf(data)

	t := reflect.TypeOf(data)

	if dataValue.Kind() != reflect.Struct {
		fmt.Println(dataValue.Kind())
		return errors.New("Input is not a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		k := dataValue.Type().Field(i)
		value := dataValue.Field(i).Interface()
		var valueStr string

		switch value.(type) {
		case string:
			valueStr = fmt.Sprintf(`%v`, value)
			if valueStr == "" {
				return errors.New(k.Name + " Empty (" + valueStr + ")")
			}
		case int:

			if value.(int) == 0 {
				return errors.New(k.Name + " Empty (0)")
			}
		case float32:
			if value.(float32) == 0 {
				return errors.New(k.Name + " Empty (0)")
			}
		case float64:
			if value.(float64) == 0 {
				return errors.New(k.Name + " Empty (0)")
			}

		case time.Time:
			valueStr = fmt.Sprintf(`%s`, value.(time.Time).Format("2006-01-02 15:04:05"))
			if valueStr == "0001-01-01 00:00:00" {
				return errors.New(k.Name + " Empty (0001-01-01T00:00:00Z0)")
			}
		case interface{}:
			err := CheckAllFieldsAreSet(value)
			if err != nil {
				return err
			}
		default:
			valueStr = fmt.Sprintf("DEFAULT:%v", value)
			return errors.New("Type is not checked:" + k.Name + ":" + valueStr)

		}
	}

	return nil
}

func MatchSQLInject(tomathstr string) bool {
	//过滤 ‘
	//ORACLE 注解 --  /**/
	//关键字过滤 update ,delete
	// 正则的字符串, 不能用 " " 因为" "里面的内容会转义
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		return false
	}
	return re.MatchString(tomathstr)
}
