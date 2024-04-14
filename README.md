# libreoffice

Golang bindings to [LibreOfficeKit](https://docs.libreoffice.org/libreofficekit.html)

## install libreofficekit
```bash
sudo apt-get install libreoffice libreofficekit-dev
```

## install package

```bash
go get github.com/atopx/libreoffice@latest
```

## example for convert

```go
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
```
