package generator

type GenerationType int32

const (
	Metadata GenerationType = iota
	Archive
)

func (s GenerationType) String() string {
	switch s {
	case Metadata:
		return "meta"
	case Archive:
		return "archive"
	}
	return "unknown"
}
