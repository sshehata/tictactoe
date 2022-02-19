package game

import (
  "fmt"
)

type position struct {
  x, y int
}

func (p *position) String() string {
  return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *position) X() int {
  return p.x
}

func (p *position) Y() int {
  return p.y
}
