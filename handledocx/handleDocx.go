package Handledocx

import "github.com/nguyenthenguyen/docx"

func On(outContent string, outCode string) {
	rr, err := docx.ReadDocxFile("./text.docx")
	if err != nil {
		panic(err)
	}

	docx1 := rr.Editable()

	docx1.Replace("{{content}}", outContent, 1)
	docx1.Replace("{{code}}", outCode, 1)

	docx1.WriteToFile("./vlh_out.docx")

	rr.Close()
}
