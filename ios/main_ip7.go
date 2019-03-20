package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"runtime/debug"


	jump "github.com/faceair/youjumpijump"
	imgo "github.com/Comdex/imgo"
	"time"
)

var r = jump.NewRequest()
var timer =time.AfterFunc(time.Second*60,xuanshang_touch)
var ip string
var touch_count = 0;
var touch_flag = true;
var count = 0;
var fubenModel = ""
var distance = 0.01
type ScreenshotRes struct {
	Value     string `json:"value"`
	SessionID string `json:"sessionId"`
	Status    int    `json:"status"`
}
func fubenSelect(model string)(string){
	if model == "1"{
		return  "hunshi"
	}else if model == "2" {
		return  "yuling"
	}else {
		return "unknown"
	}
}
func screenshot(ip string) (*ScreenshotRes, image.Image) {
	_, body, err := r.Get(fmt.Sprintf("http://%s/screenshot", ip))
	if err != nil {
		log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
	}

	res := new(ScreenshotRes)
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Fatal("WebDriverAgentRunner 响应数据异常，请检查 WebDriverAgentRunner 运行状态")
	}

	pngValue, err := base64.StdEncoding.DecodeString(res.Value)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}

	pic, err := png.Decode(bytes.NewReader(pngValue))
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}
	return res, pic
}
func hunshi_screenshot(ip string) (*ScreenshotRes, image.Image) {
	_, body, err := r.Get(fmt.Sprintf("http://%s/screenshot", ip))
	if err != nil {
		log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
	}

	res := new(ScreenshotRes)
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Fatal("WebDriverAgentRunner 响应数据异常，请检查 WebDriverAgentRunner 运行状态")
	}

	pngValue, err := base64.StdEncoding.DecodeString(res.Value)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}

	pic, err := png.Decode(bytes.NewReader(pngValue))
	rgbimage := pic.(*image.RGBA)
	subimage := rgbimage.SubImage(image.Rect(620,510,653,800)).(*image.RGBA)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}
	return res, subimage
}
func jiesuan_screenshot(ip string) (*ScreenshotRes, image.Image) {
	_, body, err := r.Get(fmt.Sprintf("http://%s/screenshot", ip))
	if err != nil {
		log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
	}

	res := new(ScreenshotRes)
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Fatal("WebDriverAgentRunner 响应数据异常，请检查 WebDriverAgentRunner 运行状态")
	}

	pngValue, err := base64.StdEncoding.DecodeString(res.Value)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}

	pic, err := png.Decode(bytes.NewReader(pngValue))
	rgbimage := pic.(*image.RGBA)
	subimage := rgbimage.SubImage(image.Rect(115,566,310,770)).(*image.RGBA)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}
	return res, subimage
}
func shengli_screenshot(ip string) (*ScreenshotRes, image.Image) {
	_, body, err := r.Get(fmt.Sprintf("http://%s/screenshot", ip))
	if err != nil {
		log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
	}

	res := new(ScreenshotRes)
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Fatal("WebDriverAgentRunner 响应数据异常，请检查 WebDriverAgentRunner 运行状态")
	}

	pngValue, err := base64.StdEncoding.DecodeString(res.Value)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}

	pic, err := png.Decode(bytes.NewReader(pngValue))
	rgbimage := pic.(*image.RGBA)
	subimage := rgbimage.SubImage(image.Rect(400,615,547,889)).(*image.RGBA)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}
	return res, subimage
}
func xuanshang_screenshot(ip string) (*ScreenshotRes, image.Image) {
	_, body, err := r.Get(fmt.Sprintf("http://%s/screenshot", ip))
	if err != nil {
		log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
	}

	res := new(ScreenshotRes)
	err = json.Unmarshal(body, res)
	if err != nil {
		log.Fatal("WebDriverAgentRunner 响应数据异常，请检查 WebDriverAgentRunner 运行状态")
	}

	pngValue, err := base64.StdEncoding.DecodeString(res.Value)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}

	pic, err := png.Decode(bytes.NewReader(pngValue))
	rgbimage := pic.(*image.RGBA)
	subimage := rgbimage.SubImage(image.Rect(530,514,619,817)).(*image.RGBA)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}
	return res, subimage
}
func xuanshang_touch(){
	res3, pic := xuanshang_screenshot(ip)
	jump.SavePNG("IP7_jump.png", pic)
	cos3,_err:= imgo.CosineSimilarity("IP7_jump.png","IP7_xuanshang.png")
	cos3_1w:=cos3*10000
	if _err!=nil{
		println(_err.Error())
	}
	if(cos3_1w>9500.0){
		log.Println(cos3_1w)
		_, _, err := r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/tap/0", ip, res3.SessionID), map[string]interface{}{
			"y":        jump.Random(260, 280),
			"x":        jump.Random(439, 500),
		})
		if err != nil {
			log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
		}
	}
}
func fuben_touch(){
	res, pic := hunshi_screenshot(ip)
	jump.SavePNG("IP7_jump.png", pic)
	cos1,_ := imgo.CosineSimilarity("IP7_jump.png","IP7_"+fubenModel+".png")
	cos1_1w:=cos1*10000
	if(cos1_1w>9998.0){
		log.Println(cos1_1w)
		timer.Reset(time.Second*60)
		pic = nil
		_, _, err := r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/tap/0", ip, res.SessionID), map[string]interface{}{
			"y":        jump.Random(244, 269),
			"x":        jump.Random(470, 527),
		})
		if err != nil {
			log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
		}
		touch_flag = false
		touch_count++
		if touch_count >=5 {
			log.Fatal("按下次数过多")
		}
	}
}
func jiesuan_touch(){
	res2, pic := jiesuan_screenshot(ip)
	jump.SavePNG("IP7_jump.png", pic)
	cos2,_ := imgo.CosineSimilarity("IP7_jump.png","IP7_jiesuan.png")
	cos2_1w:=cos2*10000
	if(cos2_1w>9900.0){
		count++;
		log.Println(cos2_1w)
		log.Println(count)
		timer.Reset(time.Second*10)
		_, _, err := r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/doubleTap", ip, res2.SessionID), map[string]interface{}{
			"y":        jump.Random(80, 230),
			"x":        jump.Random(30, 90),
		})
		if err != nil {
			log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
		}
		time.Sleep(time.Millisecond *  1000)
		touch_flag = true
		touch_count = 0
	}
}
func main() {
	defer func() {
		jump.Debugger()
		if e := recover(); e != nil {
			log.Printf("%s: %s", e, debug.Stack())
			fmt.Print("程序已崩溃，按任意键退出")
			var c string
			fmt.Scanln(&c)
		}
	}()

	flag.StringVar(&ip, "ip", "", "WebDriverAgentRunner 监听的 IP 和端口 (例如 192.168.9.94:8100)")
	flag.Parse()

	ip="localhost:8100"

	if ip == "" {
		fmt.Print("请输入 WebDriverAgentRunner 监听的 IP 和端口 (例如 192.168.9.94:8100):")
		_, err := fmt.Scanln(&ip)
		if err != nil {
			log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
		}
	}
	var fubenInput =""
	fmt.Print("1.魂十业原火\n2.御灵\n请输入数字选择副本:")
	fmt.Scanln(&fubenInput)
	if fubenInput == "" {
		log.Fatal("未选择副本，退出")
	}
	fubenModel = fubenSelect(fubenInput)
	if fubenModel == "unknown"{
		log.Fatal("未知副本，退出")
	}
	touch_flag = true
	for {
		jump.Debugger()

		//if time.Now().Unix() >= 1530416429{
		//	log.Fatal("时间到")
		//}

		if touch_flag == true{
			fuben_touch()
		}

		if touch_flag == false {
			jiesuan_touch()
		}
		//xuanshang_touch()
	}
}
