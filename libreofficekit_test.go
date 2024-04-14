package libreoffice_test

import (
	"testing"

	"github.com/atopx/libreoffice"
)

func TestConvert(t *testing.T) {
	office, err := libreoffice.NewOffice("/usr/lib/libreoffice/program")
	if err != nil {
		t.Fatal(err)
	}
	doc, err := office.LoadDocumentWith("test.doc", "en-US")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		doc.Close()
		office.Close()
	}()
	doc.SaveAs("test.pdf", "pdf", "")
}
