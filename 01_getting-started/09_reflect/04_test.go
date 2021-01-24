/*
* @Time    : 2021-01-23 11:35
* @Author  : CoderCharm
* @File    : 04_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

利用反射校验参数

源项目参考地址
https://github.com/flipped-aurora/gin-vue-admin

尝试添加一项 正则校验

**/
package _9_reflect

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// 定义一些简单等参数校验函数
// 小于 <
func Lt(mark string) string {
	return "lt=" + mark
}

// 小于等于 <=
func Le(mark string) string {
	return "le=" + mark
}

// 等等于 ==
func Eq(mark string) string {
	return "eq=" + mark
}

// 大于等于 >=
func Ge(mark string) string {
	return "ge=" + mark
}

// 大于 >
func Gt(mark string) string {
	return "gt=" + mark
}

// ...其他就不列举了

// 非空
func NotEmpty() string {
	return "notEmpty"
}

// 正则判断
func Regex(regex string) string {
	return "regex=" + regex
}

// 存放验证规则的地方
type Rules map[string][]string

func verify(st interface{}, roleMap Rules) (err error) {

	// 限定 比较返回值为 以下几个
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)

	// 判断待验证参数 是否是结构体 不是直接返回错误
	if val.Kind() != reflect.Struct {
		return errors.New("expect struct")
	}

	// 遍历结构体的所有字段
	for i := 0; i < val.NumField(); i++ {

		// 获取反射后的具体字段
		tagVal := typ.Field(i)
		val := val.Field(i)

		// 判断此字段是否有校验规则 >0 则说明有
		if len(roleMap[tagVal.Name]) > 0 {

			// 循环此字段的校验规则(一个字段可以存在多个校验规则)  规则为 各个判断类型函数 返回值
			for _, v := range roleMap[tagVal.Name] {
				switch {
				// 非空判断
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				// 正则校验
				case strings.Split(v, "=")[0] == "regex":
					if !isRegexMatch(val.String(), v) {
						return errors.New(tagVal.Name + "正则校验不合法" + v)
					}

				// 比较符判断 分割返回值里面的 = 符号  compareMap 确保输入的函数正确
				case compareMap[strings.Split(v, "=")[0]]:
					// 比较值    val 为反射字段后的值  v为 lt=1等校验值
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "长度或值不在合法范围," + v)
					}
				default:
					fmt.Println("检查 Rules 校验函数输入是否正确: " + v)
				}
			}
		}
	}
	return nil
}

// 判断是否为空
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// 正则校验
func isRegexMatch(value string, VerifyStr string) bool {
	regexStr := strings.Split(VerifyStr, "=")[1]
	match, _ := regexp.MatchString(regexStr, value)
	return match

}

func compareVerify(value reflect.Value, VerifyStr string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		return compare(value.Len(), VerifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), VerifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), VerifyStr)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), VerifyStr)
	default:
		return false
	}
}

func compare(value interface{}, VerifyStr string) bool {
	VerifyStrArr := strings.Split(VerifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// 分割后的string转换为整数类型
		VInt, VErr := strconv.ParseInt(VerifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Int() < VInt
		case VerifyStrArr[0] == "le":
			return val.Int() <= VInt
		case VerifyStrArr[0] == "eq":
			return val.Int() == VInt
		case VerifyStrArr[0] == "ne":
			return val.Int() != VInt
		case VerifyStrArr[0] == "ge":
			return val.Int() >= VInt
		case VerifyStrArr[0] == "gt":
			return val.Int() > VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		VInt, VErr := strconv.Atoi(VerifyStrArr[1])
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Uint() < uint64(VInt)
		case VerifyStrArr[0] == "le":
			return val.Uint() <= uint64(VInt)
		case VerifyStrArr[0] == "eq":
			return val.Uint() == uint64(VInt)
		case VerifyStrArr[0] == "ne":
			return val.Uint() != uint64(VInt)
		case VerifyStrArr[0] == "ge":
			return val.Uint() >= uint64(VInt)
		case VerifyStrArr[0] == "gt":
			return val.Uint() > uint64(VInt)
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		VFloat, VErr := strconv.ParseFloat(VerifyStrArr[1], 64)
		if VErr != nil {
			return false
		}
		switch {
		case VerifyStrArr[0] == "lt":
			return val.Float() < VFloat
		case VerifyStrArr[0] == "le":
			return val.Float() <= VFloat
		case VerifyStrArr[0] == "eq":
			return val.Float() == VFloat
		case VerifyStrArr[0] == "ne":
			return val.Float() != VFloat
		case VerifyStrArr[0] == "ge":
			return val.Float() >= VFloat
		case VerifyStrArr[0] == "gt":
			return val.Float() > VFloat
		default:
			return false
		}
	default:
		return false
	}
}

// 任意一个函数
func AnyFunc() string {
	return "任意函数的返回值"
}

var (
	// AnyFunc() 为非法的校验函数 避免人为写错
	UserInfoVerify = Rules{"Page": {Ge("1"), AnyFunc()}, "PageSize": {Le("50")}, "Name": {Regex(`^\d{3}$`), NotEmpty()}}
)

type UserInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
}

func TestVerify(t *testing.T) {

	// 参数信息
	u := UserInfo{Page: 1, PageSize: 30, Name: "234"} // 合法
	//u := UserInfo{Page:1, PageSize:30, Name: "1234"} // Name非法

	// 验证参数是否合法
	if err := verify(u, UserInfoVerify); err != nil {
		t.Log(fmt.Sprintf("验证失败 %s \n", err))
	} else {
		t.Log("success")
	}
}
