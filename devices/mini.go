package devices

import (
	"image"

	streamdeck "github.com/xkisu/go-streamdeck"
)

var (
	miniName                     string
	miniButtonWidth              uint
	miniButtonHeight             uint
	miniImageReportPayloadLength uint
	miniImageReportHeaderLength  uint
	miniImageReportLength        uint
)

// GetImageHeaderMini returns the USB comms header for a button image for the XL
func GetImageHeaderMini(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	var thisLength uint
	if miniImageReportPayloadLength < bytesRemaining {
		thisLength = miniImageReportPayloadLength
	} else {
		thisLength = bytesRemaining
	}
	header := []byte{
		'\x02',
		'\x01',
		byte(pageNumber),
		0,
		get_header_element(thisLength, bytesRemaining),
		byte(btnIndex + 1),
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
	}

	return header
}

func get_header_element(thisLength, bytesRemaining uint) byte {
	if thisLength == bytesRemaining {
		return '\x01'
	} else {
		return '\x00'
	}
}

func init() {
	miniName = "Streamdeck Mini"
	miniButtonWidth = 80
	miniButtonHeight = 80
	miniImageReportPayloadLength = 1024
	streamdeck.RegisterDevicetype(
		miniName, // Name
		image.Point{X: int(miniButtonWidth), Y: int(miniButtonHeight)}, // Width/height of a button
		0x63,                         // USB productID
		resetPacket17(),              // Reset packet
		6,                            // Number of buttons
		2,                            // Number of rows
		3,                            // Number of cols
		brightnessPacket17(),         // Brightness packet
		1,                            // Button read offset
		"BMP",                        // Image format
		miniImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderMini,           // Function to get the comms image header
	)
}
