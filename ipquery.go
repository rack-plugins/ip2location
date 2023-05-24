package ip2location

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/ip2location/ip2location-go/v9"
)

func IpQuery(ctx *gin.Context) {
	ip := ctx.Param("ip")
	res, err := Query(ip)
	if err != nil {
		ezap.Error("ip 查询失败: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func Query(ip string) (map[string]interface{}, error) {
	t := ip2location.OpenTools()

	if !t.IsIPv4(ip) && !t.IsIPv6(ip) {
		return nil, fmt.Errorf("传入 ip[%s] 格式错误", ip)
	}

	db, err := ip2location.OpenDB(DB_FILENAME)
	if err != nil {
		ezap.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	results, err := db.Get_all(ip)
	resMap := mapStruct(results)
	resMap["IP"] = ip

	return resMap, err
}

// 过滤数据库不包含的字段
func mapStruct(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	m := make(map[string]interface{})
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		value := v.Field(k).Interface()
		if value != "This parameter is unavailable for selected data file. Please upgrade the data file." {
			m[name] = value
		}
	}
	return m
}
