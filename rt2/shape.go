package rt2

type Shape interface {
	Hit(r Ray) bool
}
