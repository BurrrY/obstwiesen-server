package imgHandler

import (
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image/jpeg"
	"os"
)

func ResizeImage(src_path string, dst_path string, width int) (string, error) {

	srcImage, err := imaging.Open(src_path, imaging.AutoOrientation(true))
	dstImage800 := imaging.Resize(srcImage, width, 0, imaging.Lanczos)

	// Create the file to which the resized image will be written.
	out, err := os.Create(dst_path)
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer out.Close()

	// Encode and write the resized image to the new file.
	err = jpeg.Encode(out, dstImage800, nil)
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Info("Resized Image ", src_path, " to ", dst_path)

	return dst_path, nil
}
