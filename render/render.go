package render

import (
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)

var gui *tview.Application
var table *tview.Table

const (
  width = 3
  height = 3
)


func init() {
  table = tview.NewTable().
  SetBorders(true)
  for r := 0; r < height; r++ {
    for c := 0; c < width; c++ {
      color := tcell.ColorWhite

      cell := tview.NewTableCell("-").
        SetTextColor(color).
        SetAlign(tview.AlignCenter)
      table.SetCell(r, c, cell)
    }
  }
  gui = tview.NewApplication()
  if err := gui.SetRoot(table, true).Run(); err != nil {
    panic(err)
  }
}

// Render draw current game state
func Render() {
  for r := 0; r < height; r++ {
    for c := 0; c < width; c++ {
      color := tcell.ColorRed

      cell := tview.NewTableCell("-").
        SetTextColor(color).
        SetAlign(tview.AlignCenter)
      table.SetCell(r, c, cell)
    }
  }

  gui.QueueUpdateDraw(func () {
    gui.Draw()
  })

}
