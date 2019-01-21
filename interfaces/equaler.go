package interfaces


type Equaler interface {
	Equal(Equaler) (comparable bool, equal bool)
}

