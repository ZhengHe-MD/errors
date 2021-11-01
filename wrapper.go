package errors

func EConflict(args ...interface{}) error {
	return eWithSkip(2, append(args, Conflict)...)
}

func EInternal(args ...interface{}) error {
	return eWithSkip(2, append(args, Internal)...)
}

func EInvalid(args ...interface{}) error {
	return eWithSkip(2, append(args, Invalid)...)
}

func ENotFound(args ...interface{}) error {
	return eWithSkip(2, append(args, NotFound)...)
}

func EPermDenied(args ...interface{}) error {
	return eWithSkip(2, append(args, PermDenied)...)
}
