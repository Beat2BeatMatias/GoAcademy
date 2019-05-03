package errors

type MiError struct {
	mensaje string
	radicando float64
}

func New(m string, r float64) error  {
	return &MiError{m, r}
}

func (e *MiError) Error() string {
	return e.mensaje
}

func (e *MiError) IsNegativo() bool {
	return e.radicando < 0
}