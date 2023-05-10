package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	backFiles   []string
	waterFiles  []string
	waterLeft   int
	waterTop    int
	waterWidth  int
	waterHeight int
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SelectBackFiles() {
	a.backFiles = SelectImages(a.ctx)
}

func (a *App) GetBackFiles() (ret []string) {
	for _, file := range a.backFiles {
		src, err := ioutil.ReadFile(file)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
		}
		ret = append(ret, base64.StdEncoding.EncodeToString(src))
	}
	return
}

func (a *App) SelectWaterFiles() {
	a.waterFiles = SelectImages(a.ctx)
}

func (a *App) GetWaterFiles() (ret []string) {
	for _, file := range a.waterFiles {
		src, err := ioutil.ReadFile(file)
		if err != nil {
			runtime.LogError(a.ctx, err.Error())
		}
		ret = append(ret, base64.StdEncoding.EncodeToString(src))
	}
	return
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
