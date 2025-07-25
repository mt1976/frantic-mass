package graphs

type Trace struct {
	Name    string
	X       []string
	Y       []string
	Shape   string
	Text    []string // Optional
	XIsTime bool     // Indicates if X values are time-based
}

// NewTrace creates a new trace with name, X/Y values and shape
func NewTrace(name string, x, y []string, shape string) Trace {
	return Trace{
		Name:  name,
		X:     x,
		Y:     y,
		Shape: shape,
	}
}

// WithText sets the optional text field (e.g., tooltip content)
func (t Trace) WithText(texts []string) Trace {
	t.Text = texts
	return t
}

// AddX appends one or more x-values to the Trace
func (t *Trace) AddX(values ...string) {
	t.X = append(t.X, values...)
}

// AddY appends one or more y-values to the Trace
func (t *Trace) AddY(values ...string) {
	t.Y = append(t.Y, values...)
}

func (t *Trace) AddXY(x, y string) {
	t.X = append(t.X, x)
	t.Y = append(t.Y, y)

}

func (t *Trace) AddXYText(x, y, text string) {
	t.X = append(t.X, x)
	t.Y = append(t.Y, y)
	t.Text = append(t.Text, text)
}
