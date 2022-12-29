package helper

func CatchPanic(err error) {
	if err != nil {
		AppearJSON("Terjadi kesalahan")
		AppearJSON(err)
	}
}
