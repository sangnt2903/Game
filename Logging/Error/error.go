package Error

type ErrorInterface interface {
	HasError(err error) bool
}

type Error struct {
}

var ErrorService ErrorInterface = &Error{}

func (e Error) HasError(err error) bool {
	return err != nil
}

var _ ErrorInterface = &Error{}
