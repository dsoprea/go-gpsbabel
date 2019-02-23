package gpsbabel

import (
	"errors"
	"fmt"
	"io"

	"os/exec"

	"github.com/dsoprea/go-logging"
)

const (
	FormatGpx = "gpx"
)

var (
	ErrConversionFailed = errors.New("conversion failed")
)

type Babel struct {
	fromFormat string
	toFormat   string
}

func NewBabel(fromFormat, toFormat string) *Babel {
	return &Babel{
		fromFormat: fromFormat,
		toFormat:   toFormat,
	}
}

func (b *Babel) Convert(r io.Reader, w io.Writer) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	parameters := []string{
		"-i", b.fromFormat,
		"-f", "/dev/stdin",
		"-o", b.toFormat,
		"-F", "-",
	}

	cmd := exec.Command("gpsbabel", parameters...)
	cmd.Stdin = r
	cmd.Stdout = w

	err = cmd.Run()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok == true {
			fmt.Printf("STDERR:\n\n%s\n", string(ee.Stderr))

			log.Panic(ErrConversionFailed)
		}

		log.Panic(err)
	}

	return nil
}

func Convert(fromFormat, toFormat string, r io.Reader, w io.Writer) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	b := NewBabel(fromFormat, toFormat)

	err = b.Convert(r, w)
	log.PanicIf(err)

	return nil
}

func ConvertToGpx(fromFormat string, r io.Reader, w io.Writer) (err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	b := NewBabel(fromFormat, FormatGpx)

	err = b.Convert(r, w)
	log.PanicIf(err)

	return nil
}
