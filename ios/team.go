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
var distance = 0.01
type ScreenshotRes struct {
	Value     string `json:"value"`
	SessionID string `json:"sessionId"`
	Status    int    `json:"status"`
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
	subimage := rgbimage.SubImage(image.Rect(160,215,556,931)).(*image.RGBA)
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
	subimage := rgbimage.SubImage(image.Rect(90,470,260,679)).(*image.RGBA)
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
	subimage := rgbimage.SubImage(image.Rect(432,615,547+32,889)).(*image.RGBA)
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
	subimage := rgbimage.SubImage(image.Rect(458,440,525,695)).(*image.RGBA)
	if err != nil {
		log.Fatal("图片解码失败，请参考 https://github.com/faceair/youjumpijump/issues/41")
	}
	return res, subimage
}
func xuanshang_touch(){
	res3, pic := xuanshang_screenshot(ip)
	jump.SavePNG("jump_team.png", pic)
	cos3,_err:= imgo.CosineSimilarity("jump_team.png","SE_xuanshang_team.png")
	cos3_1w:=cos3*10000
	if _err!=nil{
		println(_err.Error())
	}
	if(cos3_1w>9500.0){
		log.Println(cos3_1w)
		_, _, err := r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/touchAndHold", ip, res3.SessionID), map[string]interface{}{
			"x":        jump.Random(164, 194),
			"y":        jump.Random(750, 886),
			"duration": distance * 2 / 10000,
		})
		if err != nil {
			log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
		}
	}
}
func hunshi_touch(){
	res, pic := shengli_screenshot(ip)
	jump.SavePNG("jump_team.png", pic)
	cos1,_ := imgo.CosineSimilarity("jump_team.png","SE_shengli_team.png")
	cos1_1w:=cos1*10000
	if(cos1_1w>9800.0){
		log.Println(cos1_1w)
		timer.Reset(time.Second*10)
		pic = nil
		_, _, err := r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/touchAndHold", ip, res.SessionID), map[string]interface{}{
			"x":        jump.Random(432, 579),
			"y":        jump.Random(615, 889),
			"duration": distance / 10000,
		})
		_, _, err = r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/touchAndHold", ip, res.SessionID), map[string]interface{}{
			"x":        jump.Random(432, 579),
			"y":        jump.Random(615, 889),
			"duration": distance / 10000,
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
	jump.SavePNG("jump_team.png", pic)
	cos2,_ := imgo.CosineSimilarity("jump_team.png","SE_jiesuan.png")
	cos2_1w:=cos2*10000
	if(cos2_1w>9900.0){
		count++;
		log.Println(cos2_1w)
		log.Println(count)
		timer.Reset(time.Second*20)
		_, _, err := r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/touchAndHold", ip, res2.SessionID), map[string]interface{}{
			"x":        jump.Random(165, 380),
			"y":        jump.Random(950, 1100),
			"duration": distance / 10000,
		})
		if err != nil {
			log.Fatal("WebDriverAgentRunner 连接失败，请参考 https://github.com/faceair/youjumpijump/issues/71")
		}
		time.Sleep(time.Millisecond *  300)
		_, _, err = r.PostJSON(fmt.Sprintf("http://%s/session/%s/wda/touchAndHold", ip, res2.SessionID), map[string]interface{}{
			"x":        jump.Random(165, 380),
			"y":        jump.Random(950, 1100),
			"duration": distance / 10000,
		})
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

	for {
		jump.Debugger()


		if touch_flag == true{
			hunshi_touch()
		}

		if touch_flag == false {
			jiesuan_touch()
		}
	}
}
