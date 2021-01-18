package lvm

const (
	B   = "B"
	KB  = "KB"
	MB  = "MB"
	TB  = "TB"
	PB  = "PB"
	KiB = "KiB"
	MiB = "MiB"
	GiB = "GiB"
	TiB = "TiB"
	PiB = "PiB"
)

var sizeUnit = map[string]float64{
	"B":   1,
	"KB":  1000,                             // kilobyte
	"MB":  1000000,                          // megabyte
	"GB":  1000000000,                       // gigabyte
	"TB":  1000000000000,                    // terabyte
	"PB":  1000000000000000,                 // petabyte
	"KiB": 1024,                             // kibibyte
	"MiB": 1024 * 1024,                      // mebibyte
	"GiB": 1024 * 1024 * 1024,               // gibibyte
	"TiB": 1024 * 1024 * 1024 * 1024,        // tebibyte
	"PiB": 1024 * 1024 * 1024 * 1024 * 1024, // pebibyte
}

func BytesToHumanReadable(byte uint64, units string) float64 {
	return float64(byte) / sizeUnit[units]
}

func HumanReadableToBytes(byte int64, units string) int64 {
	return byte * int64(sizeUnit[units])
}
