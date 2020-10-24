package models

import (
	"unicode/utf8"
)

type tagValidatorFunc func(*Tag) error

func runTagValidatorFuncs(tag *Tag, fns ...tagValidatorFunc) error {
	for _, f := range fns {
		err := f(tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (hv *tagValidator) idGreaterThanZero(tag *Tag) error {
	if tag.ID <= 0 {
		return ErrInvalidID
	}
	return nil
}

func (hv *tagValidator) quackIDGreaterThanZero(tag *Tag) error {
	if tag.QuackID <= 0 {
		return ErrInvalidID
	}
	return nil
}

func (hv *tagValidator) quackExists(tag *Tag) error {
	// TODO
	return nil
}

func (hv *tagValidator) truncateHash(tag *Tag) error {
	_, i := utf8.DecodeRuneInString(tag.Text)
	tag.Text = tag.Text[i:]
	return nil
}
