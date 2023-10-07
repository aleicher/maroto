package fixture

import (
	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func BarcodeProp() props.Barcode {
	prop := props.Barcode{
		Top:     10,
		Left:    10,
		Percent: 98,
		Proportion: props.Proportion{
			Width:  16,
			Height: 9,
		},
		Center: false,
	}
	prop.MakeValid()
	return prop
}

func RectProp() props.Rect {
	prop := props.Rect{
		Top:     10,
		Left:    10,
		Percent: 98,
		Center:  false,
	}
	prop.MakeValid()
	return prop
}

func CellEntity() entity.Cell {
	return entity.Cell{
		X:      10,
		Y:      15,
		Width:  100,
		Height: 150,
	}
}

func CellProp() props.Cell {
	prop := props.Cell{
		BackgroundColor: &props.Color{
			Red:   255,
			Green: 100,
			Blue:  50,
		},
		BorderColor: &props.Color{
			Red:   200,
			Green: 80,
			Blue:  60,
		},
		BorderType:      border.Left,
		BorderThickness: 0.6,
		LineStyle:       linestyle.Dashed,
	}
	return prop
}

func Node(rootType string) *node.Node[core.Structure] {
	marotoNode := node.New[core.Structure](core.Structure{
		Type: rootType,
	})
	pageNode := node.New[core.Structure](core.Structure{
		Type: "page",
	})

	marotoNode.AddNext(pageNode)
	return marotoNode
}

func ColorProp() props.Color {
	return props.Color{
		Red:   100,
		Green: 50,
		Blue:  200,
	}
}

func LineProp() props.Line {
	prop := props.Line{
		Color:         ColorProp(),
		Style:         linestyle.Dashed,
		Thickness:     1.1,
		Orientation:   orientation.Vertical,
		OffsetPercent: 50,
		SizePercent:   20,
	}
	prop.MakeValid()
	return prop
}
