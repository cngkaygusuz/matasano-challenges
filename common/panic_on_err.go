package common

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}