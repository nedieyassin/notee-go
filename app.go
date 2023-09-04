package main

import (
	"context"
	"notee/store"
)

// App struct
type App struct {
	ctx context.Context
}

// NoteeApp NewApp creates a new App application struct
func NoteeApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	store.CreateSchema()
}

// CreateNote creates a new note in the App.
//
// It takes a store.Note parameter, which represents the note to be created.
// There is no return value.
func (a *App) CreateNote(note store.Note) uint {
	return store.CreateNote(&note)
}

func (a *App) DeleteNote(id uint) bool {
	return store.DeleteNote(&id)
}

func (a *App) RenameNote(id uint, title string) bool {
	return store.UpdateNoteTitle(&id, &title)
}

func (a *App) FavNote(id uint, isFav uint) bool {
	return store.UpdateFav(&id, &isFav)
}

func (a *App) UpdateNote(id uint, body string) bool {
	return store.UpdateNote(&id, &body)
}

// GetNotes returns an array of store.Note objects.
//
// This function does not take any parameters.
// It returns a slice of store.Note objects.
func (a *App) GetNotes() []store.Note {
	return *store.GetNotes()
}

func (a *App) GetNote(id uint) store.Note {
	return *store.GetNote(&id)
}
