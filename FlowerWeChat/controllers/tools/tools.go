package tools

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"reflect"
	"net"
	"errors"
	"golang.org/x/crypto/scrypt"
	"math"
)

var currentPath string = getCurrentPath()

//检测是否存在该路径
func IsDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}

//错误检查
func CheckError(err error)  {
	if err != nil {
		fmt.Println(err)
		return
	}
}



//***
//[]byte转string
//**
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}


//***
//form匹配struct
//**
//用map填充结构
func FillStruct(data map[string][]string, obj interface{}) error {
	for k, v := range data {
		fmt.Println(v[0])
		err := SetField(obj, k, v[0])
		if err != nil {
			return err
		}
	}
	return nil
}
//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值
	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}
	structFieldValue.Set(val)
	return nil
}
//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	//else if .......增加其他一些类型的转换
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

//获取当前ip地址
func GetIPAddress() string {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		CheckError(err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//获取当前项目路径
func getCurrentPath() string {
	file, _ := os.Getwd()
	return file
}

//加密字符串
func EncryptionPassWord(pwd, salt string) string  {
	password, _ := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x",password)
}

//获取当前时间
func TimeNow() time.Time  {
	return time.Now().UTC().Add(time.Hour*8)
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	//var firstpage int //前一页地址
	//var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		//firstpage = page - 1
		//lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		//firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		//firstpage = page - 1
		//lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		//firstpage = int(math.Max(float64(1), float64(page-1)))
		//lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = 1
	paginatorMap["lastpage"] = totalpages
	paginatorMap["currpage"] = page
	return paginatorMap
}
