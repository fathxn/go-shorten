package qr

import "github.com/skip2/go-qrcode"

type GenerateQR interface {
	QRGenerator(shortURL string) error
}

type generateQR struct {
}

func NewGenerateQR() GenerateQR {
	return &generateQR{}
}

func (g *generateQR) QRGenerator(shortURL string) error {
	// var png []byte
	if err := qrcode.WriteFile(shortURL, qrcode.Medium, 256, "qr.png"); err != nil {
		return err
	}
	return nil
}
