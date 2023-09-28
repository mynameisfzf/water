package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/strutil"
	"github.com/nfnt/resize"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"image"
	"image/draw"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"
)

// App struct
type App struct {
	ctx        context.Context
	backFiles  map[string]string
	waterFiles map[string]string
}

type SetImage struct {
	WaterFile   string
	WaterWidth  int
	WaterHeight int

	BackFile   string
	BackWidth  int
	BackHeight int
}

func NewApp() *App {
	return &App{
		backFiles:  map[string]string{},
		waterFiles: map[string]string{},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SelectBackFiles() {
	files := SelectImages(a.ctx)
	for _, file := range files {
		data, err := GetImageBase64(file)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			msg(a.ctx, "出错了", err.Error())
			return
		}
		a.backFiles[file] = data
	}

}

func (a *App) GetBackFiles() map[string]string {
	return a.backFiles
}

func (a *App) SelectWaterFiles() {
	files := SelectImages(a.ctx)
	for _, file := range files {
		data, err := GetImageBase64(file)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
			msg(a.ctx, "出错了", err.Error())
			return
		}
		a.waterFiles[file] = data
	}
	return
}

func (a *App) GetWaterFiles() map[string]string {
	return a.waterFiles
}

func (a *App) Delimg(name string, typ int) map[string]string {
	if typ > 0 {
		//水印
		delete(a.waterFiles, name)
		return a.waterFiles
	}
	delete(a.backFiles, name)
	return a.backFiles
}
func (a *App) GetSetImage() (ret SetImage) {
	err := a.getBackFileOne(&ret)
	if err != nil {
		msg(a.ctx, "出错了", err.Error())
		return
	}
	err = a.getWaterFileOne(&ret)
	if err != nil {
		msg(a.ctx, "出错了", err.Error())

	}
	return
}

func (a *App) getBackFileOne(conf *SetImage) error {
	for _, data := range a.backFiles {
		d, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return err
		}
		config, _, err := image.DecodeConfig(bytes.NewReader(d))
		if err != nil {
			return err
		}
		conf.BackFile = data
		conf.BackWidth = config.Width
		conf.BackHeight = config.Height
		return nil
	}
	return errors.New("没有设置背景图片")
}

func (a *App) getWaterFileOne(conf *SetImage) error {
	for _, data := range a.waterFiles {
		d, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return err
		}
		config, _, err := image.DecodeConfig(bytes.NewReader(d))
		if err != nil {
			return err
		}
		conf.WaterFile = data
		conf.WaterWidth = config.Width
		conf.WaterHeight = config.Height
		return nil
	}
	return errors.New("没有设置背景图片")
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
	total := len(a.backFiles) * len(a.waterFiles)
	index := 0
	for bfName, backf := range a.backFiles {
		for wfName, waterf := range a.waterFiles {
			index++
			nfile := outdir + "/" + getFileName(bfName) + "_" + getFileName(wfName) + ".png"
			err = generate(backf, waterf, nfile, realTop, realLeft, realWidth, realHeight)
			if err != nil {
				runtime.LogError(a.ctx, err.Error())
				msg(a.ctx, "出错了", err.Error())
				return
			}
			rate := 100 * index / total
			runtime.EventsEmit(a.ctx, "starting", rate)
		}

	}
	msg(a.ctx, "完成了", "图片已全部生成")
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
	src, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(src), nil
}

func getImageByData(data string) (image.Image, error) {
	src, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(src)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	if img == nil {
		return nil, errors.New("未知图片格式")
	}
	return img, nil
}

func loghander(message string, err error) {
	if err != nil {
		str := fmt.Sprintf("%s %s", message, err)
		runtime.LogError(context.Background(), str)
		msg(context.Background(), "出错了", str)
	}
}

func generate(backFile, waterFile, savefile string, top, left, width, height int) error {
	bImg, err := getImageByData(backFile)
	if err != nil {
		return err
	}
	wImg, err := getImageByData(waterFile)
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
	return png.Encode(imgDist, m)

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
