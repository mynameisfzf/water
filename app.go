package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/strutil"
	"github.com/nfnt/resize"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	backFiles  []string
	waterFiles []string
	// outDir     string
}

type SetImage struct {
	WaterFile   string
	WaterWidth  int
	WaterHeight int

	BackFile   string
	BackWidth  int
	BackHeight int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

func (a *App) SelectBackFiles() {
	a.backFiles = SelectImages(a.ctx)
}

func (a *App) GetBackFiles() (ret []string) {
	for _, file := range a.backFiles {
		data, err := GetImageBase64(file)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			return
		}
		ret = append(ret, data)
	}
	return
}

func (a *App) SelectWaterFiles() {
	a.waterFiles = SelectImages(a.ctx)
}

func (a *App) GetWaterFiles() (ret []string) {
	for _, file := range a.waterFiles {
		data, err := GetImageBase64(file)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			return
		}
		ret = append(ret, data)
	}
	return
}

func (a *App) GetSetImage() (ret SetImage) {

	if len(a.backFiles) > 0 {
		data, err := GetImageBase64(a.backFiles[0])
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			return
		}

		ret.BackWidth, ret.BackHeight = GetImageWH(a.backFiles[0])
		ret.BackFile = data

	}

	if len(a.waterFiles) > 0 {
		data, err := GetImageBase64(a.waterFiles[0])
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			return
		}
		ret.WaterFile = data
		ret.WaterWidth, ret.WaterHeight = GetImageWH(a.waterFiles[0])
	}

	// fmt.Println(ret)
	return
}

func (a *App) SetOutDir() string {
	dir, err := SelectDir(a.ctx)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
	}
	return dir
}

func (a *App) Start(outdir string, top, left, width, height int, resizeRate float64) {
	err := createDir(outdir)
	if err != nil {
		msg(a.ctx, "错误", err.Error())
		return
	}
	if len(a.backFiles) == 0 {
		msg(a.ctx, "错误", "至少要有一张背景图片")
		return
	}
	if len(a.waterFiles) == 0 {
		msg(a.ctx, "错误", "至少要有一张水印图片")
		return
	}
	realTop := int(float64(top) * resizeRate)
	realLeft := int(float64(left) * resizeRate)
	realWidth := int(float64(width) * resizeRate)
	realHeight := int(float64(height) * resizeRate)
	for _, backf := range a.backFiles {
		for _, waterf := range a.waterFiles {
			nfile := outdir + "/" + getFileName(backf) + "_" + getFileName(waterf) + ".png"
			err = generate(backf, waterf, nfile, realTop, realLeft, realWidth, realHeight)
			fmt.Println(err)
		}
	}
}

func SelectImages(ctx context.Context) []string {
	filter := runtime.FileFilter{
		DisplayName: "图片文件",
		Pattern:     "*.jpg;*.jpeg;*.png",
	}
	files, err := runtime.OpenMultipleFilesDialog(ctx, runtime.OpenDialogOptions{
		Title:   "选择文件",
		Filters: []runtime.FileFilter{filter},
	})
	if err != nil {
		runtime.LogError(ctx, err.Error())
		return []string{}
	}
	return files
}

func SelectDir(ctx context.Context) (string, error) {
	return runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{
		Title: "保存",
	})
}

func GetImageBase64(file string) (string, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(src), nil
}

func GetImageWH(file string) (int, int) {
	handle, err := os.Open(file)
	Loghander("打开文件失败", err)
	if err != nil {

		return 0, 0
	}
	defer handle.Close()
	img, _, err := image.DecodeConfig(handle)
	Loghander("打开图片失败", err)
	// log.Println(s)
	return img.Width, img.Height
}

func GetImage(file *os.File) (image.Image, error) {
	file.Seek(0, 0)
	img, str, err := image.Decode(file)
	fmt.Println(str)
	return img, err
	// _, format, err := image.DecodeConfig(file)
	// if (format != "jpeg" && format != "gif" && format != "png") || err != nil {
	// 	Loghander("image config", err)
	// 	return nil, err
	// }

	// if format == "jpge" {
	// 	return jpeg.Decode(file)
	// } else if format == "png" {
	// 	return png.Decode(file)
	// } else if format == "gif" {
	// 	return gif.Decode(file)
	// }
	// return nil, errors.New("未知类型")

	// img, err := png.Decode(file)
	// if err == nil {

	// 	return &img, nil
	// }
	// Loghander("png", err)
	// img, err = jpeg.Decode(file)
	// if err == nil {

	// 	return &img, nil
	// }
	// Loghander("jpg", err)
	// img, err = gif.Decode(file)
	// if err == nil {

	// 	return &img, nil
	// }
	// Loghander("gif", err)

}

func Loghander(message string, err error) {
	if err != nil {
		log.Printf("%s %s", message, err)
	}
}

func generate(backFile, waterFile, savefile string, top, left, width, height int) error {
	back, err := os.Open(backFile)
	if err != nil {
		return err
	}
	defer back.Close()
	water, err := os.Open(waterFile)
	if err != nil {
		return err
	}
	defer water.Close()

	bImg, err := GetImage(back)
	if err != nil {
		return err
	}
	wImg, err := GetImage(water)
	if err != nil {
		return err
	}

	wImg = resize.Resize(uint(width), uint(height), wImg, resize.Lanczos3)

	bimgBounds := (bImg).Bounds()
	m := image.NewRGBA(bimgBounds)
	draw.Draw(m, bimgBounds, bImg, image.Point{0, 0}, draw.Src)
	draw.Draw(m, wImg.Bounds().Add(image.Pt(int(left), int(top))), wImg, image.Point{0, 0}, draw.Src)
	imgDist, err := os.Create(savefile)
	if err != nil {
		return err
	}
	defer imgDist.Close()
	png.Encode(imgDist, m)

	return nil
}

func createDir(path string) error {
	f, err := os.Stat(path)
	if err == nil {
		if f.IsDir() {
			return nil
		}
		return errors.New(path + "已经存在且不是一个目录")
	}
	return os.MkdirAll(path, 0666)
}

func getFileName(path string) string {

	return strutil.BeforeLast(filepath.Base(path), ".")

}

func msg(ctx context.Context, title, message string) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         title,
		Message:       message,
		DefaultButton: "Ok",
	})
}
