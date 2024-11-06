package assert

/*New- If condition == false it raises assertion error*/
func New(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
