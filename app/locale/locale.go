package locale

import (
	"encoding/json"
	"fmt"
	"gohub-api/pkg/app"
	"os"

	"github.com/gin-gonic/gin"
)

func (c gin.Context) Translate(key string) string {
	appPath := app.GetAppPath()
	fileName := appPath + "\\locale\\" + "zh-CN.json"

	jsonFile, err := os.Open(fileName)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	messageMap := make(map[string]string)
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&messageMap)
	//message := helpers.JSONToMap(contents)
	fmt.Println(messageMap)
	return fileName

}
