package models

type quackValidatorFunc func(*Quack) error

func runQuackValidatorFuncs(quack *Quack, fns ...quackValidatorFunc) error {
	for _, f := range fns {
		err := f(quack)
		if err != nil {
			return err
		}
	}
	return nil
}

func (qv *quackValidator) idGreaterThanZero(quack *Quack) error {
	if quack.ID <= 0 {
		return ErrInvalidID
	}
	return nil
}
