package common

// CheckErr is panic error if error is not nil
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
