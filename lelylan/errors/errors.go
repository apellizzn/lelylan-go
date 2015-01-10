package errors

type LelylanError struct{
	Code string
	Description string
}

type LelylanHttpFail struct{
	Error LelylanError
	Status int
}