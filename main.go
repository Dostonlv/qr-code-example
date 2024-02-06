package main

import qrcode "qr-code/qr-code"

func main() {
	resp, err := qrcode.QRCode("sabuuasub", 2)
	if err != nil {
		panic(err)
	}

	println(resp)

}
