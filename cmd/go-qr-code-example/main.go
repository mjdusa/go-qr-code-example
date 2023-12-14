package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create the barcode scaled to 256x256 pixels
	err := createQR("https://www.linkedin.com/in/michael-donahue-15138314/", qr.M, qr.Auto, 256, 256, "linkedin-qrcode.png")
	if err != nil {
		panic(err)
	}

	err = createQR("https://www.google.com/", qr.M, qr.Auto, 256, 256, "google-qrcode.png")
	if err != nil {
		panic(err)
	}

	err = createQR("https://www.godaddy.com/", qr.M, qr.Auto, 256, 256, "godaddy-qrcode.png")
	if err != nil {
		panic(err)
	}
}

func createQR(content string, level qr.ErrorCorrectionLevel, mode qr.Encoding, width int, height int, fileName string) error {
	// Create the barcode
	qrCode, err := qr.Encode(content, level, mode)
	if err != nil {
		return err
	}

	// Scale the barcode
	qrCode, err = barcode.Scale(qrCode, width, height)
	if err != nil {
		return err
	}

	// create the output file
	file, ferr := os.Create(fileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	// encode the barcode as png
	perr := png.Encode(file, qrCode)
	if perr != nil {
		return perr
	}

	return nil
}
