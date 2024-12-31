package usecase

import "io"

type File struct {
	fileName string
	Body io.ReadSeeker
}

func (f *File) GetFileName() string {
	return f.fileName
}

func (f *File) GetBody() io.ReadSeeker {
	return f.Body
}

