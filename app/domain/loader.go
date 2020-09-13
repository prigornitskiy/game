package domain

import (
	"github.com/faiface/pixel"
	"image"
	"os"
)

type LoaderInterface interface {
	LoadPicture(name string, path string) error
}

type Loader struct {
	repository PictureRepositoryInterface
}

func NewLoader(pictureRepository PictureRepositoryInterface) *Loader {
	return &Loader{
		repository: pictureRepository,
	}
}

func (l *Loader) LoadPicture(name string, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	l.repository.AddPicture(name, pixel.PictureDataFromImage(img))

	return nil
}

type PictureRepositoryInterface interface {
	AddPicture(name string, picture pixel.Picture)
	GetPicture(name string) pixel.Picture
}

type PictureRepository struct {
	PictureRepositoryInterface
	pictures map[string]pixel.Picture
}

func NewPictureRepository() *PictureRepository {
	return &PictureRepository{
		pictures: make(map[string]pixel.Picture),
	}
}

func (p *PictureRepository) AddPicture(name string, picture pixel.Picture) {
	p.pictures[name] = picture
}

func (p *PictureRepository) GetPicture(name string) pixel.Picture {
	return p.pictures[name]
}
