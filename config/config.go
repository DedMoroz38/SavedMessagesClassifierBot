package config

type CallbackQuery int

const (
	AddFolder CallbackQuery = iota
	AddFile
)

var CallbackQueries = map[CallbackQuery]string{
	AddFolder: "add_folder",
	AddFile:   "add_file",
}
