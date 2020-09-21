package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "mm", PageSize: gopdf.Rect{W: 57, H: 120}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("HDZB_5", "./ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("HDZB_5", "", 10)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "xxxx")
	pdf.Br(30)
	pdf.Cell(nil, "Exxxxxx")
	pdf.WritePdf("hello.pdf")

}
