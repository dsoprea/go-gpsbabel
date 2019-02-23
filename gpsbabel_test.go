package gpsbabel

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"testing"
	"time"

	"github.com/dsoprea/go-gpx"
	"github.com/dsoprea/go-gpx/reader"
	"github.com/dsoprea/go-logging"
)

type gpxPoint struct {
	timestamp   time.Time
	coordinates [2]float64
}

func testGpxOutput(t *testing.T, r io.Reader) {
	collected := make([]gpxPoint, 0)

	tpc := func(tp *gpxcommon.TrackPoint) (err error) {
		p := gpxPoint{
			tp.Time,
			[2]float64{tp.LatitudeDecimal, tp.LongitudeDecimal},
		}

		collected = append(collected, p)

		return nil
	}

	err := gpxreader.EnumerateTrackPoints(r, tpc)
	log.PanicIf(err)

	if len(collected) != 9 {
		t.Fatalf("The right number of records was not found.")
	}

	firstItem := collected[0]

	actualDescription := fmt.Sprintf("%s %.6f %.6f", firstItem.timestamp, firstItem.coordinates[0], firstItem.coordinates[1])
	expectedDescription := "2019-02-05 08:07:05 +0000 UTC 38.760708 -9.112968"

	if actualDescription != expectedDescription {
		t.Fatalf("First item does not match: [%s] != [%s]", actualDescription, expectedDescription)
	}

	lastItem := collected[len(collected)-1]

	actualDescription = fmt.Sprintf("%s %.6f %.6f", lastItem.timestamp, lastItem.coordinates[0], lastItem.coordinates[1])
	expectedDescription = "2019-02-05 08:07:18 +0000 UTC 38.760751 -9.112848"

	if actualDescription != expectedDescription {
		t.Fatalf("Last item does not match: [%s] != [%s]", actualDescription, expectedDescription)
	}
}

func TestBabel_Convert(t *testing.T) {
	b := NewBabel("v900", "gpx")

	filepath := path.Join(TestAssetPath, "19020501_Portugal2.CSV.head")

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	buffer := new(bytes.Buffer)

	err = b.Convert(f, buffer)
	if err != nil {
		if log.Is(err, ErrConversionFailed) == true {
			fmt.Printf("STDOUT:\n\n%s\n", buffer.String())
		}

		log.Panic(err)
	}

	testGpxOutput(t, buffer)
}

func TestConvert(t *testing.T) {
	filepath := path.Join(TestAssetPath, "19020501_Portugal2.CSV.head")

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	buffer := new(bytes.Buffer)

	err = Convert("v900", "gpx", f, buffer)
	if err != nil {
		if log.Is(err, ErrConversionFailed) == true {
			fmt.Printf("STDOUT:\n\n%s\n", buffer.String())
		}

		log.Panic(err)
	}

	testGpxOutput(t, buffer)
}

func TestConvertToGpx(t *testing.T) {
	filepath := path.Join(TestAssetPath, "19020501_Portugal2.CSV.head")

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	buffer := new(bytes.Buffer)

	err = ConvertToGpx("v900", f, buffer)
	if err != nil {
		if log.Is(err, ErrConversionFailed) == true {
			fmt.Printf("STDOUT:\n\n%s\n", buffer.String())
		}

		log.Panic(err)
	}

	testGpxOutput(t, buffer)
}
