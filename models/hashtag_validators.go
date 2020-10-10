package models

type hashtagValidatorFunc func(*Hashtag) error

func runHashtagValidatorFuncs(hashtag *Hashtag, fns ...hashtagValidatorFunc) error {
	for _, f := range fns {
		err := f(hashtag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (hv *hashtagValidator) idGreaterThanZero(hashtag *Hashtag) error {
	if hashtag.ID <= 0 {
		return ErrInvalidID
	}
	return nil
}

func (hv *hashtagValidator) quackIDGreaterThanZero(hashtag *Hashtag) error {
	if hashtag.QuackID <= 0 {
		return ErrInvalidID
	}
	return nil
}

func (hv *hashtagValidator) quackExists(hashtag *Hashtag) error {
	// TODO
	return nil
}
