package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".html"))
	// fmt.Println(mime.AddExtensionType(".docx", "Docx"))
	fmt.Println(mime.TypeByExtension(".docx"))
	fmt.Println(mime.QEncoding.Encode("utf-8", "Hello!"))
	fmt.Println(mime.BEncoding.Encode("UTF-8", "¡Hola, señor!"))
	fmt.Println(mime.QEncoding.Encode("ISO-8859-1", "Caf\xE9"))
}

// https://stackoverflow.com/questions/29838185/how-to-detect-additional-mime-type-in-golang
/*
TypeByExtension returns the MIME type associated with the file extension ext. The extension ext should begin with a leading dot, as in ".html". When ext has no associated type, TypeByExtension returns "".

Extensions are looked up first case-sensitively, then case-insensitively.

The built-in table is small but on unix it is augmented by the local system's mime.types file(s) if available under one or more of these names:

/etc/mime.types
/etc/apache2/mime.types
/etc/apache/mime.types
On Windows, MIME types are extracted from the registry.

Text types have the charset parameter set to "utf-8" by default.
*/
