package method

type ErrorMsg struct {
	Msg string
}

func (em ErrorMsg) Error() string {
	return em.Msg
}