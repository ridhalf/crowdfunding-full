package helper

func ResultOrError[T any](result T, err error) (T, error) {
	if err != nil {
		return result, err
	}
	return result, nil
}
