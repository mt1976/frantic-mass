package graphs

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

var mini *minify.M

func init() {
	// Initialize the minifier with JavaScript support
	mini = minify.New()
	mini.AddFunc("text/javascript", js.Minify)
}
