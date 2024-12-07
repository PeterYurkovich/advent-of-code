package helpers

func Assert(passed bool, reason string) {
	if !passed {
		panic(reason)
	}
}

func AssertError(err error) {
	if err != nil {
		Assert(false, err.Error())
	}
}
