package util

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(bookingCode string) (string, error) {
	imagedQrCode, err := qrcode.New(bookingCode, qrcode.Medium)
	if err != nil {
		return err.Error(), err
	}

	buffer := new(bytes.Buffer)
	png.Encode(buffer, imagedQrCode.Image(256))
	base64QRCode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return base64QRCode, nil
}