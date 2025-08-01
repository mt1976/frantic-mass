package contentProvider

import "github.com/mt1976/frantic-mass/app/web/glyphs"

var WeightWildcard = "{weightId}"           // Wildcard for the weight ID in the URI
var WeightURI = "/weight/" + WeightWildcard // Define the URI for the weight measurement
var WeightName = "Weight"                   // Name for the weight measurement
var WeightIcon = glyphs.Weight              // Icon for the weight measurement
var WeightHover = "User %s Weight"          // Description for the weight measurement
