package cameraalarm

import (
	"encoding/json"
	"encoding/xml"
	"httpserver-temp/model/cameraalarm"
	"io/ioutil"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CameraAlarm 摄像头报警事件处理函数
func CameraAlarm(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorln("error: " + err.Error())
		return
	}
	v := cameraalarm.EventNotificationAlert{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		log.Errorln("error: " + err.Error())
		return
	}

	byteValue, err := ioutil.ReadFile("./dev.json")
	if err != nil {
		time.Sleep(10 * time.Millisecond)
		byteValue, err = ioutil.ReadFile("./dev.json")
		if err != nil {
			log.Errorln("error: " + err.Error())
			return
		}
	}
	pudatas := cameraalarm.PubDatas{}
	err = json.Unmarshal(byteValue, &pudatas)
	if err != nil {
		log.Errorln("error: " + err.Error())
		return
	}
	log.Infoln(v.Extensions.SerialNumber)
	devs := pudatas.Devs
	for index, dev := range devs {
		if dev.SubSN == v.Extensions.SerialNumber {
			pudatas.Devs[index].DateTime = v.DateTime
			pudatas.Devs[index].EventType = v.EventType
			pudatas.Devs[index].EventState = v.EventState
			pudatas.Devs[index].SerialNumber = v.Extensions.SerialNumber
			pudatas.Devs[index].IPAddress = v.IPAddress
			pudatas.Devs[index].MacAddress = v.MacAddress
			pudatas.Devs[index].CameraName = dev.SubName
		}
	}

	content, err := json.Marshal(pudatas)
	if err != nil {
		log.Errorln("error: " + err.Error())
	}

	// filename := "data.json"
	filename := "dev.json"
	writefile(filename, content)

}

func writefile(filename string, content []byte) {
	//开始写文件
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	if err != nil {
		time.Sleep(10 * time.Millisecond)
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
		if err != nil {
			log.Errorln("error: " + err.Error())
		}
	}
	defer f.Close()

	n, err := f.Write(content)
	if err != nil {
		time.Sleep(10 * time.Second)
		n, err = f.Write(content)
		if err != nil {
			log.Errorln("error: " + err.Error())
		}
	}
	if 0 < n {
		log.Infof("file write sucess. n = %v", n)
	}
}
