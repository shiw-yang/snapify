package main

import (
	"embed"

	app "github.com/shiwyang/snapify/backend/application"
	constant "github.com/shiwyang/snapify/backend/constant"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := app.App()
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "snapify",
		Width:     constant.LockWidth,
		Height:    constant.LockHeight,
		MinWidth:  constant.LockWidth,
		MinHeight: constant.LockHeight,
		MaxWidth:  constant.LockWidth,
		MaxHeight: constant.LockHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: options.NewRGBA(constant.BackgroundColour[0], constant.BackgroundColour[1], constant.BackgroundColour[2], constant.BackgroundColour[3]),
		OnStartup:        app.OnStartup,
		Bind:             app.Bind(),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
