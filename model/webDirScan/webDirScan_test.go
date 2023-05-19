package webDirScan

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func Test_putColorTest(t *testing.T) {

	clenit := &http.Client{}
	resp, _ := clenit.Get("http://www.baidu.com")

	for _, v := range resp.Header {
		log.Println(v)
	}
	//fmt.Printf("\033[32;40m%s\033[0m\n", "红色文字，黑色底哒")
	//fmt.Printf("\x1b[32m%s 32: 绿 \x1b[0m", "test")
	//
	//tmpurl := "https://test.com//123"
	//oneurl := tmpurl[7:]
	//twourl := tmpurl[:7]
	//
	//if strings.Contains(tmpurl[7:], "//") {
	//	oneurl = strings.Replace(oneurl, "//", "/", -1)
	//}
	//
	//Glog.Println(twourl + oneurl)
}

//func Test_timer1(t *testing.T) {
//	ticker := time.NewTicker(5 * time.Second)
//
//	for {
//		select {
//		case <- ticker.C:
//			Glog.Println("ticker .")
//		default:
//
//		}
//	}
//}

func Test_putcharTest(t *testing.T) {
	messagxx := make(chan string)
	go func() {
		for {
			messagxx <- "123"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg := <-messagxx:
			log.Println("received message", msg)
		default:
			log.Println("no message received")
		}
	}

}
