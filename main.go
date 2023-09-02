package main

import (
	"embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func openFile(_ *menu.CallbackData) {
	//store.CreateNote()
	println("---------------------- Inserted ----------------------------")
}

func main() {
	// Create an instance of the app structure
	app := NoteeApp()

	AppMenu := menu.NewMenu()

	FileMenu := AppMenu.AddSubmenu("Manage")
	FileMenu.AddText("New Note", keys.CmdOrCtrl("n"), openFile)
	FileMenu.AddText("New Folder", keys.CmdOrCtrl("d"), openFile)
	FileMenu.AddText("Find...", keys.CmdOrCtrl("f"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Enter", keys.Key("enter"), openFile)
	FileMenu.AddText("Back", keys.Key("escape"), openFile)

	AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut / on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Notee",
		Width:     1024,
		Height:    768,
		MinWidth:  500,
		MinHeight: 700,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		OnStartup:        app.startup,
		Menu:             AppMenu,
		Mac: &mac.Options{
			Appearance: mac.NSAppearanceNameDarkAqua,
			About: &mac.AboutInfo{
				Title:   "Notee",
				Message: "© 2021 yassin",
			},
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
