package tool

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"go/types"
	"reflect"
	"strconv"
	"strings"
)
func Map_to_json(map_str map[string]interface{}) (string,error){
	b, err := json.Marshal(map_str)
	if err != nil {
		//fmt.Println("json.Marshal failed:", err)
		return "",err
	}
	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(map_str)
	return  string(b),nil
}
func Get_value_type(i interface{})string {
	ty:=""
	switch i.(type) {
	case bool:
		ty="bool"
	case string:
		ty="string"
	case types.Array:
		ty="Array"
	case int:
		ty="int"
	case int8:
		ty="int8"
	case int32:
		ty="int32"
	case int64:
		ty="int64"
	case float32:
		ty="float32"
	case float64:
		ty="float64"
	case types.Map:
		ty="Map"
	case map[string]int:
		ty="map[string]int"
	case map[string]string:
		ty="map[string]string"
	case map[string]interface{}:
		ty="map[string]interface{}"
	case []map[string]interface{}:
		ty="[]map[string]interface{}"
	case []interface{}:
		ty="[]interface{}"
	}

	return ty
}
func MapGroup_to_jsonGroup(map_str []map[string]interface{}) (string,error){
	b, err := json.Marshal(map_str)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return "",err
	}
	return  string(b),nil
}
func Json_to_map(json_string string) (map[string]interface{},error){
	//json str 转map
	var map_data map[string]interface{}
	if err := json.Unmarshal([]byte(json_string), map_data); err != nil {
		return nil ,err
	}
	return map_data,nil
}
func Json_to_struct(json_string string,config *interface{}) (*interface{},error){
	if err := json.Unmarshal([]byte(json_string), config); err != nil {
		return nil ,err
	}
	return  config,nil
}
func Struct_to_json(config *interface{}) (string,error){
	json_string,err := json.Marshal( config)
	if  err != nil {
		return "" ,err
	}
	return  string(json_string),nil
}
func Struct_to_map(obj interface{})map[string]interface{}  {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
func String_to_int(value string) (int,error) {
	i,err:=strconv.Atoi(value)
	if  err != nil {
		return 0,err
	}
	return i,nil
}
func String_to_int64(value string) (int64,error) {
	i,err:=strconv.ParseInt(value, 10, 64)
	if  err != nil {
		return 0,err
	}
	return i,nil
}
func String_to_float64(value string)(float64,error){
	v2, err := strconv.ParseFloat(value, 64)
	return v2,err
}
func int_to_int64(i int)(int64,error){
	s := strconv.Itoa(i)
	s64,err := strconv.ParseInt(s,10,64)
	return s64,err
}
func int64_to_int(i int64)(int){
	s:=int(i)
	return s
}
func Map_to_struct(Map interface{}, pointer interface{}) {
	// reflect.Ptr类型 *main.Person
	pointertype := reflect.TypeOf(pointer)
	// reflect.Value类型
	pointervalue := reflect.ValueOf(pointer)
	// struct类型 main.Person
	structType := pointertype.Elem()
	// 将interface{}类型的map转换为 map[string]interface{}
	m := Map.(map[string]interface{})
	// 遍历结构体字段
	for i := 0; i < structType.NumField(); i++ {
		// 获取指定字段的反射值
		f := pointervalue.Elem().Field(i)
		//获取struct的指定字段
		stf := structType.Field(i)
		// 获取tag
		name := strings.Split(stf.Tag.Get("json"), ",")[0]
		// 判断是否为忽略字段
		if name == "-" { continue }
		// 判断是否为空，若为空则使用字段本身的名称获取value值
		if name == "" {
			name = stf.Name
		}
		//获取value值
		v, ok := m[name]
		if !ok { continue }
		//获取指定字段的类型
		kind := pointervalue.Elem().Field(i).Kind()
		// 若字段为指针类型
		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = f.Type().Elem().Kind()
		}
		// 设置对应字段的值
		switch kind {
			case reflect.Int:
				res, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)
				pointervalue.Elem().Field(i).SetInt(res)
			case reflect.String:
				pointervalue.Elem().Field(i).SetString(fmt.Sprint(v))
				}
		}
	}
//replace_num为替换次数 为-1时全部替换
func String_repalce(value string,old string,new string ) string{
	return strings.Replace(value, old, new, -1)
}
//replace_num为替换次数
func String_repalce_num(value string,old string,new string,replace_num  int ) string{
	return strings.Replace(value, old, new, replace_num)

}

func String2IntSlice(value string) []int {
	s := strings.Split(value,",")
	ids := make([]int,0)
	for _,v := range s {
		id,_ := strconv.Atoi(v)
		ids = append(ids, id)
	}
	return ids
}

// int64转string
func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

// string转int64
func StringToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err
}
func TrimEmailSuffix(email string)string{
	s:=strings.Split(email,"@")
	if len(s)>0{
		return s[0]
	}
	return email
}
func AppendEmailSuffix(username string)string{
	if strings.HasSuffix(username,"@shopee.com"){
		return username
	}
	return username+"@shopee.com"
}

func SetPatchVersionName(versionName string) string {
	indexRight := strings.LastIndex(versionName, "sp")
	indexLeft := strings.LastIndex(versionName, ".")
	c := []byte(versionName)

	var result string
	for i:=0; i<indexLeft; i++ {
		result += string(c[i])
	}
	for i:=indexRight; i<len(c); i++ {
		result += string(c[i])
	}
	return result
}

func ConvertNewLineToSpace(s string) string {
	strs := strings.Fields(s)
	var str string
	for _, v := range strs {
		str += " " + v
	}
	str = strings.Trim(str, " ")
	return str
}

//beego orm用原装SQL时对于in中的参数只支持传数组，因此数组中有多少个值则对应的in子句中就要有多少个问号，此处为对问号的拼接
func GenerateParamOfIn(ids []string) string {
	num := len(ids)
	sqlWhere := "?"
	for i:=1; i<num; i++ {
		sqlWhere += ",?"
	}
	sqlWhere = strings.Trim(sqlWhere, ",")
	return sqlWhere
}

/**
去除[]orm.Params中某字符串中重复的字符
 */
func GetNonDuplicate(params []orm.Params, col string)  {
	for i:=0; i<len(params); i++ {
		//将对菜单的操作分隔开
		roles := strings.Split(params[i][col].(string),",")
		var temp string
		for _,v := range roles {
			//将不包含的操作添加到temp中
			if !strings.Contains(temp, v) {
				temp += "," + v
			}
			//去除temp中首尾出现的逗号
			temp = strings.Trim(temp, ",")
		}
		params[i][col] = temp
	}
}