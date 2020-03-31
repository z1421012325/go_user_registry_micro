package utils

import (
	"log"
	"fmt"
	"time"

	RD "project/db/redis"
)

var (
	DefaultRegistryTime = time.Minute*6
)
const (
	registryfield = "_Registry"
)

// string
func SaveRegistryStr(k,v string){
	key := k + registryfield
	err := RD.RD.Set(key,v,DefaultRegistryTime).Err()
	if err != nil {
		log.Printf("save to redis faild [key]: %s  [value] : %s  [error] : %s ",key,v,err)
	}
}
func GetRegistryStr(k string,v int32) bool {
	key := k + registryfield
	code ,err := RD.RD.Get(key).Result()
	if err != nil {
		return false
	}
	ValueStr := fmt.Sprintln(v)
	if code != ValueStr {
		return false
	}
	return true
}

func SaveCustomStr(k,v string,ex time.Duration){
	err := RD.RD.Set(k,v,ex).Err()
	if err != nil {
		log.Printf("save to redis faild [key]: %s  [value] : %s  [error] : %s ",k,v,err)
	}
}
func GetCustomStr(k string,v int32) bool {

	key := k + registryfield
	code ,err := RD.RD.Get(key).Result()
	if err != nil {
		return false
	}
	ValueStr := fmt.Sprintln(v)
	if code != ValueStr {
		return false
	}
	return true
}


// list
func SaveRegistryList(){}
func GetRegistryList(){}

func SaveCustomList() {}
func GetCustomList() bool {
	return false
}




// hash
func SaveRegistryHash(){}
func GetRegistryHash(){}

func SaveCustomHash() {}
func GetCustomHash() bool {
	return false
}


// set
func SaveRegistrySet(){}
func GetRegistrySet(){}

func SaveCustomSet() {}
func GetCustomSet() bool {
	return false
}






// zset
func SaveRegistryZset(){}
func GetRegistryZset(){}

func SaveCustomZset() {}
func GetCustomLZset() bool {
	return false
}

