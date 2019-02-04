package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os/exec"
	"runtime/debug"
	"strconv"
	//"time"

	jump "github.com/faceair/youjumpijump"
	"github.com/Comdex/imgo"
	"time"
)
var inputRatio float64
var touch_count = 0;
var count = 0;
var touch_flag = true;
var err error
var timer =time.AfterFunc(time.Second*60,xuanshang_touch)
func screenshot() image.Image {
	out, err := exec.Command("adb","shell"," /system/bin/screencap", "-p").Output()
	if err != nil {
		println(err.Error())
	}
	pic, err := png.Decode(bytes.NewReader(out))
	if err != nil {
		log.Println("png 截图解码失败。")
	}
	return pic
}
func jiesuan_screenshot() image.Image {
	out, err := exec.Command("adb","shell"," /system/bin/screencap", "-p").Output()
	if err != nil {
		println(err.Error())
	}
	pic, err := png.Decode(bytes.NewReader(out))
	rgbimage := pic.(*image.NRGBA)
	subimage := rgbimage.SubImage(image.Rect(626,462,843,616)).(*image.NRGBA)
	if err != nil {
		log.Println("png 截图解码失败。")
	}
	return subimage
}
func hun10_screenshot() image.Image {
	out, err := exec.Command("adb","shell"," /system/bin/screencap", "-p").Output()
	if err != nil {
		println(err.Error())
	}
	pic, err := png.Decode(bytes.NewReader(out))
	rgbimage := pic.(*image.NRGBA)
	subimage := rgbimage.SubImage(image.Rect(580,85,880,125)).(*image.NRGBA)
	if err != nil {
		log.Println("png 截图解码失败。")
	}
	return subimage
}
func xuanshang_screenshot() image.Image {
	out, err := exec.Command("adb","shell"," /system/bin/screencap", "-p").Output()
	if err != nil {
		println(err.Error())
	}
	pic, err := png.Decode(bytes.NewReader(out))
	rgbimage := pic.(*image.NRGBA)
	subimage := rgbimage.SubImage(image.Rect(600,135,876,205)).(*image.NRGBA)
	if err != nil {
		log.Println("png 截图解码失败。")
	}
	return subimage
}
func hun10_touch(){
	pic := hun10_screenshot()
	jump.SavePNG("jumpAndroid.png", pic)
	cos1, _err := imgo.CosineSimilarity("jumpAndroid.png", "samsung_hun10.png")
	cos1_1w := cos1 * 10000
	if _err != nil {
		println(_err.Error())
	}
	if (cos1_1w > 9900.0) {
		log.Println(cos1_1w)
		//go jump.SavePNG("samsung_hun10.png",hun10_screenshot())
		touchX, touchY := strconv.Itoa(jump.Random(1000, 1100)), strconv.Itoa(jump.Random(480, 510))
		_, err = exec.Command("adb", "shell", "/system/bin/sh", "/system/bin/input", "swipe", touchX, touchY, touchX, touchY,
			strconv.Itoa(int(0.01*inputRatio))).Output()
		if err != nil {
			log.Fatal("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
		}
		touch_flag = false
		touch_count++
		timer.Reset(time.Second*60)
		if touch_count >= 5 {
			log.Fatal("按下次数过多")
		}
	}
}
func jiesuan_touch(){
	pic := jiesuan_screenshot()
	jump.SavePNG("jumpAndroid.png",pic)
	cos2,_err:= imgo.CosineSimilarity("jumpAndroid.png","samsung_jiesuan.png")
	cos2_1w:=cos2*10000
	if _err!=nil{
		println(_err.Error())
	}
	if(cos2_1w>9900.0) {
		count++;
		//go jump.SavePNG("samsung_jiesuan.png",jiesuan_screenshot())
		touchX, touchY := strconv.Itoa(jump.Random(1232, 1425)), strconv.Itoa(jump.Random(443, 525))
		touch_res, err2 := exec.Command("adb", "shell", "/system/bin/sh", "/system/bin/input", "swipe", touchX, touchY, touchX, touchY,
			strconv.Itoa(int(0.01*inputRatio))).Output()
		if err2 != nil {
			log.Fatal("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
		}
		time.Sleep(time.Millisecond * 500)
		touchX, touchY = strconv.Itoa(jump.Random(1232, 1425)), strconv.Itoa(jump.Random(443, 525))
		touch_res, err2 = exec.Command("adb", "shell", "/system/bin/sh", "/system/bin/input", "swipe", touchX, touchY, touchX, touchY,
			strconv.Itoa(int(0.01*inputRatio))).Output()
		if err2 != nil {
			log.Fatal("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
		}
		if touch_res !=nil {
			log.Println(cos2_1w)
			log.Println(count)
		}
		timer.Reset(time.Second*10)
		touch_flag = true
		touch_count = 0
	}
}
func xuanshang_touch(){
	pic := xuanshang_screenshot()
	jump.SavePNG("jumpAndroid.png",pic)
	cos3,_err:= imgo.CosineSimilarity("jumpAndroid.png","samsung_xuanshang.png")
	cos3_1w:=cos3*10000
	if _err!=nil{
		println(_err.Error())
	}
	if(cos3_1w>9800.0) {
		count++;
		log.Println(cos3_1w)
		touchX, touchY := strconv.Itoa(jump.Random(942, 1066)), strconv.Itoa(jump.Random(500, 525))
		_, err = exec.Command("adb", "shell", "/system/bin/sh", "/system/bin/input", "swipe", touchX, touchY, touchX, touchY,
			strconv.Itoa(int(0.01*inputRatio))).Output()
		if err != nil {
			log.Fatal("模拟触摸失败，请检查开发者选项中的 USB 调试安全设置是否打开。")
		}
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

	for {
		jump.Debugger()

		if touch_flag == true {
			hun10_touch()
		}
		if touch_flag == false{
			jiesuan_touch()
		}
		// time.Sleep(time.Millisecond * 500)
		// xuanshang_touch()
	}
}
