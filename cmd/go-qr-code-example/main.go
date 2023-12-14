package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create the barcode scaled to 256x256 pixels
	err := createQRCode("https://www.linkedin.com/in/michael-donahue-15138314/", qr.M, qr.Auto, 256, 256, "linkedin-qrcode.png")
	if err != nil {
		panic(err)
	}

	err = createQRCode("https://www.google.com/", qr.M, qr.Auto, 256, 256, "google-qrcode.png")
	if err != nil {
		panic(err)
	}

	err = createQRCode("https://www.godaddy.com/", qr.M, qr.Auto, 256, 256, "godaddy-qrcode.png")
	if err != nil {
		panic(err)
	}

	err = createEAN13("078693597244", 256, 128, "ean13-barcode.png")
	if err != nil {
		panic(err)
	}
}

func createQRCode(content string, level qr.ErrorCorrectionLevel, mode qr.Encoding, width int, height int, fileName string) error {
	// Create the barcode
	code, err := qr.Encode(content, level, mode)
	if err != nil {
		return err
	}

	// Scale the barcode
	scale, serr := barcode.Scale(code, width, height)
	if serr != nil {
		return serr
	}

	// create the output file
	file, ferr := os.Create(fileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	// encode the barcode as png
	perr := png.Encode(file, scale)
	if perr != nil {
		return perr
	}

	return nil
}

func createEAN13(content string, width int, height int, fileName string) error {
	// Create the barcode
	code, err := ean.Encode(content)
	if err != nil {
		return err
	}

	// Scale the barcode
	scale, serr := barcode.Scale(code, width, height)
	if serr != nil {
		return serr
	}

	// create the output file
	file, ferr := os.Create(fileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	// encode the barcode as png
	perr := png.Encode(file, scale)
	if perr != nil {
		return perr
	}

	return nil
}
