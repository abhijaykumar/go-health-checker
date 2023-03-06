package display

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"techinjektion.dev/tools/healthchecker/health"
)

func applicationsScreen(_ fyne.Window) fyne.CanvasObject {
	apps := health.Sites
	t := widget.NewTable(
		func() (int, int) { return len(apps) + 1, 5 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Name")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				switch id.Col {
				case 0:
					label.SetText("Name")
				case 1:
					label.SetText("URL")
				case 2:
					label.SetText("Http Method")
				case 3:
					label.SetText("Wait Seconds")
				case 4:
					label.SetText("Health Status")
				}
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				switch id.Col {
				case 0:
					label.SetText(apps[id.Row-1].Name)
				case 1:
					label.SetText(apps[id.Row-1].Url)
				case 2:
					label.SetText(apps[id.Row-1].HttpMethod)
				case 3:
					label.SetText(strconv.FormatInt(apps[id.Row-1].WaitSeconds, 10))
				case 4:
					label.SetText(strconv.FormatBool(apps[id.Row-1].Up))
				}

			}
		})
	setDefaultColumnsWidth(t)
	return t
}

func setDefaultColumnsWidth(table *widget.Table) {
	colWidths := []float32{100, 200, 100, 100, 100}
	for idx, colWidth := range colWidths {
		table.SetColumnWidth(idx, colWidth)
	}
}
