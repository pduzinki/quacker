package models

type followValidatorFunc func(*Follow) error

func runFollowValidatorFuncs(follow *Follow, fns ...followValidatorFunc) error {
	for _, f := range fns {
		err := f(follow)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fv *followValidator) idGreaterThanZero(follow *Follow) error {
	if follow.ID <= 0 {
		return ErrInvalidID
	}
	return nil
}

func (fv *followValidator) userIDGreaterThanZero(follow *Follow) error {
	if follow.UserID <= 0 {
		return ErrInvalidID
	}
	return nil
}

func (fv *followValidator) followsUserIDGraterThanZero(follow *Follow) error {
	if follow.FollowsUserID <= 0 {
		return ErrInvalidID
	}
	return nil
}
