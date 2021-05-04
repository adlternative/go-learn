package main

import (
	// "bytes"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"
)

func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	_, _ = db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	// ...
}

// func writeHeader(w io.Writer, contentType string) error {
// 	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
// 		return err
// 	}
// 	if _, err := w.Write([]byte(contentType)); err != nil {
// 		return err
// 	}
// 	// ...
// }
func formatOneValue(x interface{}) string {
	switch x.(type) {
	case error:
		return (x.(error)).Error()
	case int:
		return fmt.Sprintf("%d", x.(int))
	default:
		return string("")
	}
	// if err, ok := x.(error); ok {
	// 	return err.Error()
	// }
	// if int, ok := x.(int); ok {
	// 	return fmt.Sprintf("%d", int)
	// }
	// return string("")
}

func main() {
	fmt.Print(formatOneValue(1))
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	// Output:
	// &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}

	var w io.Writer = os.Stdout
	_, ok := w.(*os.File) // success:  ok, f == os.Stdout
	fmt.Println(ok)
	_, ok = w.(*bytes.Buffer) // failure: !ok, b == nil
	fmt.Println(ok)
}
