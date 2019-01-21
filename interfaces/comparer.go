package interfaces


type comparison int8

const (
	LESSER = comparison(-1)
	GREATER = comparison(1)
	EQUAL = comparison(0)
	INCOMPARABLE = comparison(-2)
)

type Comparer interface {
	Equaler
	Compare(Comparer) (result comparison)
}


type Int64 int64


func (i Int64) Equal(j Equaler) (comparable bool, equal bool) {
	if val, ok := j.(Int64); !ok {
		return false, false
	} else {
		result := i.compare(val)
		return result != INCOMPARABLE, result == EQUAL
	}
}


func (i Int64) compare(j Int64) (result comparison) {
	var iVal, jVal = int64(i), int64(j)

	if iVal > jVal {
		return GREATER
	} else if jVal > iVal {
		return LESSER
	}

	return EQUAL
}


func (i Int64) Compare(j Comparer) (result comparison) {
	val, ok := j.(Int64)
	if !ok {
		return INCOMPARABLE
	}

	return i.compare(val)
}


type String string


func (s String) Equal(t Equaler) (comparable bool, equal bool) {
	if val, ok := t.(String); !ok {
		return false, false
	} else {
		result := s.compare(val)
		return result != INCOMPARABLE, result == EQUAL
	}
}

func (s String) compare(t String) (result comparison) {
	var sVal, tVal = string(s), string(t)

	if sVal < tVal {
		return LESSER
	} else if tVal < sVal {
		return GREATER
	}

	return EQUAL
}


func (s String) Compare(t Comparer) (result comparison) {
	val, ok := t.(String)
	if !ok {
		return INCOMPARABLE
	}

	return s.compare(val)
}
