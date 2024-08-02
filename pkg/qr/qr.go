package qr

type GenerateQR interface {
	QRGenerator(shortURL string) error
}

type generateQR struct {
}

func NewGenerateQR() GenerateQR {
	return &generateQR{}
}

func (g *generateQR) QRGenerator(shortURL string) error {
	return nil
}
