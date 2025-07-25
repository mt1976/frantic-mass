package graphs

type LegendConfig struct {
	YStepSize  float64
	TraceOrder string
	FontSize   int
	YRef       string
}

// legend := NewLegendConfig(0.5, "reversed", 16, "paper")

// // Or using fluent chaining
// legend = LegendConfig{}.
// 	WithY(0.5).
// 	WithTraceOrder("reversed").
// 	WithFontSize(16).
// 	WithYRef("paper")

func NewLegendConfig(yStepSize float64, order string, fontSize int, yref string) LegendConfig {
	return LegendConfig{
		YStepSize:  yStepSize,
		TraceOrder: order,
		FontSize:   fontSize,
		YRef:       yref,
	}
}

func (l LegendConfig) WithYSize(yStepSize float64) LegendConfig {
	l.YStepSize = yStepSize
	return l
}

func (l LegendConfig) WithTraceOrder(order string) LegendConfig {
	l.TraceOrder = order
	return l
}

func (l LegendConfig) WithFontSize(size int) LegendConfig {
	l.FontSize = size
	return l
}

func (l LegendConfig) WithYRef(yref string) LegendConfig {
	l.YRef = yref
	return l
}
