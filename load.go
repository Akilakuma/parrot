package parrot

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

// ReadEnv Read config file
func ReadEnv(v interface{}) (interface{}, error) {

	// 指定只能處理struct
	var (
		tag   = "config"
		value = reflect.ValueOf(v)
		vKind = value.Kind()
	)

	newConfig := reflect.New(value.Type())
	if vKind == reflect.Struct {
		// 定義型態： main.Config
		tValue := value.Type()

		for i := 0; i < value.NumField(); i++ {
			// 逐一處理欄位
			sf := tValue.Field(i)

			// 忽略 - 的field tag
			if sf.Tag.Get(tag) == "-" {
				continue
			}

			// 取得tag name
			tagName := sf.Tag.Get(tag)
			// 用tag name 取的環境變數
			envValue := os.Getenv(tagName)

			if envValue == "" {
				return newConfig, errors.New(tagName + " is empty")
			}

			// 取得struct欄位
			f := reflect.Indirect(newConfig).FieldByName(fmt.Sprintf(sf.Name))
			// 對欄位set資料
			f.Set(reflect.ValueOf(envValue))
		}

	}
	return newConfig, nil
}
