package is

func Int(i interface{}) bool {
	_, ok := i.(int)
	return ok
}

func String(i interface{}) bool {
	_, ok := i.(string)
	return ok
}
