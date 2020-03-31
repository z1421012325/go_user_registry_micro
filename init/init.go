package init

import (
	"godotenv"
	"io/ioutil"
	"log"
	"os"
	"strings"
)




func init() {
	pwd,_ := os.Getwd()
	ReadDir(pwd)
}

func ReadDir(dirName string){
	fileInfoList,err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for i := range fileInfoList {
		if strings.HasSuffix(fileInfoList[i].Name(),".env") {
			err := godotenv.Load(fileInfoList[i].Name())
			if err != nil {
				panic(err)
			}
			break
		}

		if fileInfoList[i].IsDir(){
			ReadDir(fileInfoList[i].Name())
		}
	}
}

