package generator

type GenerationType int32

const (
	Metadata GenerationType = iota
)

func (s GenerationType) String() string {
	switch s {
	case Metadata:
		return "meta"
	}
	return "unknown"
}
