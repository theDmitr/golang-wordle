package page

import "main/ui/web/utils"

type File struct {
	Body *[]byte
}

func NewFile(filename string) *File {
	page := new(File)
	page.Body, _ = utils.LoadFile(&filename)
	return page
}
