package initialize

import (
	"encoding/csv"
	"os"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/consts"
	"go.uber.org/zap"
)

func InitCSV() {
	fileName := consts.CSVFilePath
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			zap.S().Panic("创建文件失败: " + err.Error())
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		err = writer.Write([]string{"uid", "vid", "isLike", "isCollect", "timestamp"})
		if err != nil {
			zap.S().Panic("写入文件失败: " + err.Error())
		}
	}
	return
}
