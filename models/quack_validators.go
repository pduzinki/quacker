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

func (qv *quackValidator) userIDGreaterThanZero(quack *Quack) error {
	if quack.userID == 0 {
		return ErrInvalidID
	}
	return nil
}

func (qv *quackValidator) TextShorterThan160chars(quack *Quack) error {
	if len([]rune(quack.text)) > 160 {
		return ErrQuackTooLong
	}
	return nil
}
