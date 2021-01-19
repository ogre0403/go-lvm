package lvm

const (
	B   = "B"
	KB  = "KB"
	MB  = "MB"
	GB  = "GB"
	TB  = "TB"
	PB  = "PB"
	KiB = "KiB"
	MiB = "MiB"
	GiB = "GiB"
	TiB = "TiB"
	PiB = "PiB"
)

var sizeUnit = map[string]float64{
	B:   1,
	KB:  1000,                             // kilobyte
	MB:  1000 * 1000,                      // megabyte
	GB:  1000 * 1000 * 1000,               // gigabyte
	TB:  1000 * 1000 * 1000 * 1000,        // terabyte
	PB:  1000 * 1000 * 1000 * 1000 * 1000, // petabyte
	KiB: 1024,                             // kibibyte
	MiB: 1024 * 1024,                      // mebibyte
	GiB: 1024 * 1024 * 1024,               // gibibyte
	TiB: 1024 * 1024 * 1024 * 1024,        // tebibyte
	PiB: 1024 * 1024 * 1024 * 1024 * 1024, // pebibyte
}

func BytesToHumanReadable(byte uint64, units string) float64 {
	return float64(byte) / sizeUnit[units]
}

func HumanReadableToBytes(size uint64, units string) uint64 {
	return size * uint64(sizeUnit[units])
}

func UnitTranslate(size uint64, srcUnit, targetUnit string) float64 {
	b := HumanReadableToBytes(size, srcUnit)
	return BytesToHumanReadable(b, targetUnit)
}
