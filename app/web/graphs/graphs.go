package graphs

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
)

var TIMESERIES_FORMAT = "2006-01-02 15:04:05" // Default format for time series data, 2013-10-04 22:23:00

func GeneratePlotlyScript(traces []Trace, legend LegendConfig, divID string) (template.JS, error) {
	var sb strings.Builder

	// If no traces are provided, return an empty script
	if len(traces) == 0 {
		logHandler.WarningLogger.Println("No traces provided for Plotly script generation.")
		return "", fmt.Errorf("no traces provided")
	}
	//If no legend is provided, use default values
	if legend.YStepSize == 0 {
		legend.YStepSize = 0.5 // Default Y step size
	}
	if legend.TraceOrder == "" {
		legend.TraceOrder = "reversed" // Default trace order
	}
	if legend.FontSize == 0 {
		legend.FontSize = 16 // Default font size
	}
	if legend.YRef == "" {
		legend.YRef = "paper" // Default Y reference
	}
	// Check for a valid divID
	if divID == "" {
		logHandler.ErrorLogger.Println("Invalid divID provided for Plotly script generation.")
		return "", fmt.Errorf("invalid divID provided")
	}
	// Generate trace variables
	for i, trace := range traces {
		traceVar := fmt.Sprintf("trace%d", i+1)
		sb.WriteString(fmt.Sprintf("var %s = {\n", traceVar))
		if trace.XIsTime {
			sb.WriteString("  x: [\n")
			for _, x := range trace.X {
				sb.WriteString(fmt.Sprintf("    '%s',\n", x))
			}
			sb.WriteString("  ],\n")
		} else {
			sb.WriteString("  x: [\n")
			for _, x := range trace.X {
				sb.WriteString(fmt.Sprintf("    '%s',\n", strings.ReplaceAll(x, "'", "\\'")))
			}
			sb.WriteString("  ],\n")
		}

		//sb.WriteString(fmt.Sprintf("  x: %v,\n", trace.X))
		// Handle Y values
		sb.WriteString("  y: [\n")
		for _, y := range trace.Y {
			sb.WriteString(fmt.Sprintf("    %s,\n", y))
		}
		sb.WriteString("  ],\n")
		sb.WriteString("  mode: 'lines+markers',\n")
		sb.WriteString(fmt.Sprintf("  name: '%s',\n", trace.Name))

		if len(trace.Text) > 0 {
			texts := make([]string, len(trace.Text))
			for i, txt := range trace.Text {
				texts[i] = fmt.Sprintf("'%s'", strings.ReplaceAll(txt, "'", "\\'"))
			}
			sb.WriteString(fmt.Sprintf("  text: [%s],\n", strings.Join(texts, ", ")))
		}

		sb.WriteString(fmt.Sprintf("  line: {shape: '%s'},\n", trace.Shape))
		sb.WriteString("  type: 'scatter'\n")
		sb.WriteString("};\n\n")
	}

	// Assemble data array
	sb.WriteString("var data = [")
	for i := range traces {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("trace%d", i+1))
	}
	sb.WriteString("];\n\n")

	// Legend configuration
	sb.WriteString("var layout = {\n")
	sb.WriteString("  legend: {\n")
	sb.WriteString(fmt.Sprintf("    y: %v,\n", legend.YStepSize))
	sb.WriteString(fmt.Sprintf("    traceorder: '%s',\n", legend.TraceOrder))
	sb.WriteString(fmt.Sprintf("    font: {size: %d},\n", legend.FontSize))
	sb.WriteString(fmt.Sprintf("    yref: '%s'\n", legend.YRef))

	sb.WriteString("  }\n")
	sb.WriteString("};\n\n")
	// Configuration object
	sb.WriteString("var config = {\n")
	if legend.Responsive {
		sb.WriteString("  responsive: true,\n")
	} else {
		sb.WriteString("  responsive: false,\n")
	}
	sb.WriteString("};\n\n")

	// Plot call with dynamic divID
	sb.WriteString(fmt.Sprintf("Plotly.newPlot('%s', data, layout, config);", divID))

	godump.Dump(sb.String())

	// Minify the script
	minifiedScript, err := mini.String("text/javascript", sb.String())
	if err != nil {
		logHandler.ErrorLogger.Println("Error minifying Plotly script:", err)
		return "", err
	}

	godump.Dump(minifiedScript)

	return template.JS(minifiedScript), nil
}
