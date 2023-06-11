package proto

// Result of a QR/barcode scan
type EventScanData struct {
	// Bytes encoded into base64 string
	Scan string `json:"scan"`
}
