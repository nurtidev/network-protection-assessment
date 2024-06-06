package charts

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/wcharczuk/go-chart/v2"
)

func GenerateChart(c *fiber.Ctx) error {
	graph := chart.BarChart{
		Title: "Эффективность методов защиты",
		Bars: []chart.Value{
			{Value: 0.85, Label: "Антивирус"},
			{Value: 0.90, Label: "Фаервол"},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}

	c.Set("Content-Type", "image/png")
	return c.Send(buffer.Bytes())
}
