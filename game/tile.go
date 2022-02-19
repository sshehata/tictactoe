package game

type tile int

const (
  undefined tile = iota
  otile
  xtile
)

func (t tile) String() string {
  switch t {
  case otile:
    return "O"
  case xtile:
    return "X"
  default:
    return "-"
  }
}
