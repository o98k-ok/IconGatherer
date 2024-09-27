package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	gim "github.com/ozankasikci/go-image-merge"
)

func AppendOutline(front, back []byte, x, y int, out io.Writer) error {
	var frontFile, backFile *os.File
	var err error
	{
		_, format, err := image.Decode(bytes.NewReader(front))
		if err != nil {
			return err
		}

		fileFormat := "front*.png"
		switch format {
		case "png":
			fileFormat = "front*.png"
		case "jpg", "jpeg":
			fileFormat = "front*.jpg"
		default:
		}

		frontFile, err = os.CreateTemp("", fileFormat)
		if err != nil {
			return err
		}
		frontFile.Write(front)
		frontFile.Close()
	}

	{

		_, format, err := image.Decode(bytes.NewReader(back))
		if err != nil {
			return err
		}

		fileFormat := "back*.jpg"
		switch format {
		case "png":
			fileFormat = "back*.png"
		case "jpg", "jpeg":
			fileFormat = "back*.jpg"
		default:
		}

		backFile, err = os.CreateTemp("", fileFormat)
		if err != nil {
			return err
		}
		backFile.Write(back)
		backFile.Close()
	}

	grids := []*gim.Grid{
		{
			ImageFilePath: backFile.Name(),
			Grids: []*gim.Grid{
				{
					ImageFilePath: frontFile.Name(),
					OffsetX:       x, OffsetY: y,
				},
			},
		},
	}

	rgba, err := gim.New(grids, 1, 1).Merge()
	if err != nil {
		return err
	}
	return png.Encode(out, rgba)
}

func BatteryRing(front, back string, out string) {
	frontBytes, err := os.ReadFile(front)
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(bytes.NewReader(frontBytes))
	if err != nil {
		panic(front + " " + err.Error())
	}
	w, h := img.Bounds().Dx(), img.Bounds().Dy()

	backBytes, err := os.ReadFile(back)
	if err != nil {
		panic(err)
	}
	backImg, err := png.Decode(bytes.NewReader(backBytes))
	if err != nil {
		panic(err)
	}
	bw, bh := backImg.Bounds().Dx(), backImg.Bounds().Dy()

	offsetX := (bw - w) / 2
	offsetY := (bh - h) / 2

	buf := bytes.NewBuffer(nil)
	AppendOutline(frontBytes, backBytes, offsetX, offsetY, buf)

	os.WriteFile(out, buf.Bytes(), 0o644)
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: go run main.go <icon_path> <ring_path> <out_path>")
	}
	iconPath := os.Args[1]
	ringPath := os.Args[2]
	outPath := os.Args[3]

	iconFiles, err := os.ReadDir(iconPath)
	if err != nil {
		log.Fatal(err)
	}
	ringFiles, err := os.ReadDir(ringPath)
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll(outPath, 0o755)

	for _, iconFile := range iconFiles {
		for _, ringFile := range ringFiles {
			iconFilePath := filepath.Join(iconPath, iconFile.Name())
			ringFilePath := filepath.Join(ringPath, ringFile.Name())

			// remove extension
			iconFileName := strings.TrimSuffix(iconFile.Name(), filepath.Ext(iconFile.Name()))
			ringFileName := strings.TrimSuffix(ringFile.Name(), filepath.Ext(ringFile.Name()))
			outFilePath := filepath.Join(outPath, fmt.Sprintf("%s_%s.png", iconFileName, ringFileName))
			BatteryRing(iconFilePath, ringFilePath, outFilePath)
		}
	}
}
