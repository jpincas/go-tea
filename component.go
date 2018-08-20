package gotea

type Component struct {
	UniqueID string
}

func (c Component) UniqueMsg(msg string) string {
	return c.UniqueID + "_" + msg
}
