package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func FileHandler(c *gin.Context) {
	filename := c.Param("filename")
	filePath := "./public/" + filename

	fmt.Println(filePath)
	//打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusNotFound, "File not found")
		return
	}

	//关闭打开的文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)

	//传递文件
	c.File(filePath)
}
