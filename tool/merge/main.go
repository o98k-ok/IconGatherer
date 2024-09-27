package main

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"os"

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

func main() {
	airpods := "./icon/apple/production/MagicTrackpad/touchpad_2.png"
	airpodsBytes, err := os.ReadFile(airpods)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(bytes.NewReader(airpodsBytes))
	if err != nil {
		panic(err)
	}

	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	bw, bh := 552, 552

	offsetX := (bw - w) / 2
	offsetY := (bh - h) / 2

	fullRing := "./icon/tool/batteryRing/25.png"
	fullRingBytes, err := os.ReadFile(fullRing)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	AppendOutline(airpodsBytes, fullRingBytes, offsetX, offsetY, buf)

	os.WriteFile("./airpods_4_100.png", buf.Bytes(), 0o644)
}
