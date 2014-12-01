package Gonigsberg

type location struct {
	serial string
	name   string
}

func NewLocation(serial string, name string) location {
	return location{serial, name}
}

func (l *location) Serial() string {
	return l.serial
}

func (l *location) Name() string {
	return l.name
}
