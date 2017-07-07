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
	"math/rand"
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

//生成随机数
func GenerateRandomNumber(start int, end int) int {
	if end > start {
		//随机数生成器，加入时间戳保证每次生成的随机数不一样
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//生成随机数
		num := r.Intn((end - start)) + start

		return num
	}else if start == end {
		 return start
	}else {
		return 0
	}
}
