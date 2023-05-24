package ip2location

import (
	"github.com/fimreal/goutils/ezap"
	mfile "github.com/fimreal/goutils/file"
	mzip "github.com/fimreal/goutils/zip"
	"github.com/fimreal/rack/pkg/utils"
	"github.com/spf13/viper"
)

func init() {
	if viper.GetBool("ip2location") {
		go GetDB()
	}
}

var DB_FILENAME string

// 预下载数据库，并解压
func GetDB() {
	WORKDIR := viper.GetString("workdir")
	DB_LEVEL := viper.GetString("ip2location.db")
	TOKEN := viper.GetString("ip2location.token")

	DB_CODE := DB_LEVEL + "LITEBIN"
	DB_URL := "https://www.ip2location.com/download/?token=" + TOKEN + "&file=" + DB_CODE
	DB_ZIPFILE := WORKDIR + DB_CODE + ".zip"
	DB_FILENAME = WORKDIR + "IP2LOCATION-LITE-" + DB_LEVEL + ".BIN"

	// 发现旧文件，则跳过下载
	if mfile.PathExists(DB_FILENAME) {
		ezap.Warn("发现数据库文件[", DB_FILENAME, "], 跳过下载，如需重新下载数据库请手动移除旧文件")
		ezap.Infof("配置使用数据文件[%s]", DB_FILENAME)
		return
	}

	ezap.Infof("开始下载数据库[%s]]，速度取决于您的网络连接速度", DB_ZIPFILE)
	ezap.Debug("Download DB Url: ", DB_URL)

	err := utils.DownloadNotls(DB_URL, DB_ZIPFILE)
	if err != nil {
		ezap.Fatalf("下载数据库[%s]出错: %s", DB_ZIPFILE, err)
	}
	ezap.Debugf("完成数据库[%s]下载", DB_ZIPFILE)

	err = mzip.Unzip(DB_ZIPFILE, WORKDIR)
	if err != nil {
		ezap.Fatalf("解压数据库[%s]出错: %s", DB_ZIPFILE, err)
	}
	ezap.Debugf("完成数据库[%s]解压缩", DB_ZIPFILE)

	ezap.Infof("配置使用数据文件[%s]", DB_FILENAME)
}
