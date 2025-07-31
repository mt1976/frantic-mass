package contentProvider

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dao/lookup"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-mass/app/web/helpers"
	"github.com/mt1976/frantic-mass/app/web/styleHelper"
)

var Locales = lookup.Lookup{}

type AppContext struct {
	PageTitle       string
	PageDescription string // Description of the page, used for SEO and metadata
	PageSummary     string
	PageKeywords    string

	Author    string
	Copyright string // Copyright information for the application
	License   string // License type for the application
	Delimiter string // Delimiter used in the application, if applicable

	AppVersion     string // Version of the application
	AppPrefix      string // Prefix for the application, used in URLs
	AppEnvironment string // Environment in which the application is running (e.g., development, production)
	AppBuildTime   string // Build time of the application
	AppReleaseDate string // Release date of the application
	AppLogoPath    string // Path to the application logo
	AppFaviconPath string // Path to the application favicon
	AppName        string // Name of the application
	AppDescription string // Description of the application

	UserLocale string // Locale for the application, used for internationalization
	UserTheme  string

	UserErrorMessages   []string // Error messages for the user
	UserMessages        []string // General messages for the user
	UserAlerts          []string // Alerts or notifications for the user
	UserSuccessMessages []string // Messages indicating successful operations
	HasMessages         bool     // Flag to indicate if there are any messages to display

	WasSuccessful  bool
	HttpStatusCode int // HTTP status code

	SessionUserID         int    // ID of the currently logged-in user
	SessionUserName       string // Name of the currently logged-in user
	SessionUserAvatarPath string // Avatar URL of the currently logged-in user
	SessionID             string // Session ID for tracking user sessions

	Timestamp     string // Current date and time
	FormattedDate string // Current date and time in string format

	Hostname      string // Hostname or domain of the application
	HostIPAddress string // IP address of the server
	HostIdentity  string // Unique identifier for the system or application

	Page             string // Name of the current page, used for routing or identification
	TemplatePath     string // URL path of the current page
	TemplateName     string // Name of the template used for rendering the page
	TemplateFilePath string
	PageActions      helpers.Actions // Actions available on the current page, such as buttons or links
	PageHasChart     bool            // Flag to indicate if the page has a chart to display
	ChartID          string          // ID of the chart to be displayed on the page
	ChartData        template.JS     // Data for the chart to be displayed on the page
	ChartTitle       string          // Title of the chart to be displayed on the page
	Breadcrumbs      []Breadcrumb    // Breadcrumbs for navigation, each containing a title and URL
}

type Breadcrumb struct {
	Title string // Title of the breadcrumb item
	Hover string // Hover text for the breadcrumb item
	URL   string // URL of the breadcrumb item
}

func (c *AppContext) AddBreadcrumb(title, hover, url string) {
	// Add a breadcrumb to the context
	breadcrumb := Breadcrumb{
		Title: title,
		Hover: hover,
		URL:   url,
	}
	c.Breadcrumbs = append(c.Breadcrumbs, breadcrumb)
	logHandler.InfoLogger.Printf("Added breadcrumb: %s (%s) - %s", title, hover, url)

}

func (c *AppContext) RemoveBreadcrumb(title string) {
	// Remove a breadcrumb from the context by title
	for i, breadcrumb := range c.Breadcrumbs {
		if breadcrumb.Title == title {
			c.Breadcrumbs = append(c.Breadcrumbs[:i], c.Breadcrumbs[i+1:]...)
			logHandler.InfoLogger.Printf("Removed breadcrumb: %s", title)
			return
		}
	}
	logHandler.WarningLogger.Printf("Breadcrumb not found: %s", title)
}

var cache *commonConfig.Settings
var cacheChecksum string
var css styleHelper.CSS
var style styleHelper.CLASS

func init() {
	// Initialize the Common struct with default values

	cache = commonConfig.Get()
	if cache == nil {
		logHandler.ErrorLogger.Println("Failed to load common configuration")
		return
	}
	if cacheChecksum == "" {
		cacheChecksum = getCacheChecksum(cache)
	} else {
		if cacheChecksum != getCacheChecksum(cache) {
			logHandler.ErrorLogger.Println("Cache checksum mismatch, reloading configuration")
			cache = commonConfig.Get()
			cacheChecksum = getCacheChecksum(cache)
		}
	}
	//godump.Dump(cache, "Cache Configuration", cacheChecksum)

	// Get Supports Translation locales from the common configuration
	if cache.GetTranslation_PermittedLocales() != nil {
		// Parse the permitted locales from the configuration
		locales := cache.GetTranslation_PermittedLocales()
		for _, locale := range locales {
			// Add each locale to the Locales lookup
			lookupData := lookup.LookupData{
				Key:      locale.Key,
				Value:    locale.Name,
				Selected: locale.Key == cache.GetApplication_Locale(), // Mark as selected if it matches
			}
			Locales.Data = append(Locales.Data, lookupData)
		}
	} else {
		logHandler.ErrorLogger.Println("No permitted locales found in the configuration")
	}

}

func getCacheChecksum(cache *commonConfig.Settings) string {
	// Generate a checksum for the cache configuration
	checksum := fmt.Sprintf("%s-%s-%s-%s-%s-%s-%s-%s-%s",
		cache.GetApplication_Name(),
		cache.GetApplication_Description(),
		cache.GetApplication_Version(),
		cache.GetApplication_Author(),
		cache.GetApplication_Prefix(),
		cache.GetApplication_Environment(),
		cache.GetApplication_ReleaseDate(),
		cache.GetApplication_Copyright(),
		cache.GetApplication_License(),
	)
	checksum = stringHelpers.Encode(checksum)
	logHandler.InfoLogger.Println("Cache Checksum:", checksum)
	return checksum
}

func (c *AppContext) SetDefaults() {

	// Application Config

	c.PageTitle = cache.GetApplication_Name()
	if c.PageTitle == "" {
		c.PageTitle = "Frantic Mass"
	}
	c.AppDescription = cache.GetApplication_Description()
	if c.PageSummary == "" {
		c.PageSummary = "Frantic Mass Management Application"
	}
	c.PageSummary = "Unknown Summary"
	c.PageKeywords = "mass, management"

	c.Author = cache.GetApplication_Author()
	if c.Author == "" {
		c.Author = "Frantic Mass Team"
	}
	c.AppVersion = cache.GetApplication_Version()
	if c.AppVersion == "" {
		c.AppVersion = "1.0.0"
	}
	c.UserTheme = "default"
	c.AppLogoPath = "/static/images/logo.png"
	c.AppFaviconPath = "/static/images/favicon.ico"
	c.UserErrorMessages = []string{}
	c.UserMessages = []string{}
	c.UserAlerts = []string{}
	c.UserSuccessMessages = []string{}
	c.HasMessages = false

	c.WasSuccessful = true
	c.HttpStatusCode = 200 // OK
	c.SessionUserID = 0
	c.SessionUserName = "Guest"
	c.SessionUserAvatarPath = "/static/images/default-avatar.png"
	c.SessionID = ""
	c.Timestamp = time.Now().Format(time.RFC3339)
	c.FormattedDate = dateHelpers.FormatHuman(time.Now())
	c.Hostname = application.HostName()
	c.HostIPAddress = application.HostIP()
	c.HostIdentity = application.SystemIdentity()
	c.AppName = cache.GetApplication_Name()
	if c.AppName == "" {
		c.AppName = "Frantic Mass"
	}

	c.AppPrefix = cache.GetApplication_Prefix()
	if c.AppPrefix == "" {
		c.AppPrefix = "/"
	}
	c.AppEnvironment = cache.GetApplication_Environment()
	if c.AppEnvironment == "" {
		c.AppEnvironment = "development"
	}
	c.AppReleaseDate = cache.GetApplication_ReleaseDate()
	if c.AppReleaseDate == "" {
		c.AppReleaseDate = time.Now().Format("2006-01-02")
	}
	c.Copyright = cache.GetApplication_Copyright()
	if c.Copyright == "" {
		c.Copyright = "© 2023 Frantic Mass Team"
	}
	c.License = cache.GetApplication_License()
	if c.License == "" {
		c.License = "MIT"
	}
	c.UserLocale = cache.GetApplication_Locale()
	if c.UserLocale == "" {
		c.UserLocale = "en_GB"
	}
	c.AppBuildTime = time.Now().Format(time.RFC3339)
	c.Delimiter = cache.Delimiter()

	c.TemplatePath = paths.HTML().String()
	logHandler.InfoLogger.Printf("Template Path: %s", c.TemplatePath)
	c.TemplateName = "error" // Default template name, can be overridden by specific views
	c.PageActions = helpers.Actions{}
}

func (c *AppContext) AddError(err string) {
	c.UserErrorMessages = append(c.UserErrorMessages, err)
	c.HasMessages = true
	logHandler.ErrorLogger.Println("Error added to context:", err)
	c.WasSuccessful = false
}

func (c *AppContext) AddMessage(msg string) {
	c.UserMessages = append(c.UserMessages, msg)
	c.HasMessages = true
	logHandler.InfoLogger.Println("Message added to context:", msg)
}

func (c *AppContext) AddNotification(notification string) {
	c.UserAlerts = append(c.UserAlerts, notification)
	c.HasMessages = true
	logHandler.InfoLogger.Println("Notification added to context:", notification)
}

func (c *AppContext) AddSuccess(success string) {
	c.UserSuccessMessages = append(c.UserSuccessMessages, success)
	c.HasMessages = true
	logHandler.InfoLogger.Println("Success added to context:", success)
}

// AgeFromDOB calculates the age in years given a date of birth
func AgeFromDOB(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Check if the birthday has occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func IntToString(i int) string {
	if i == 0 {
		return ""
	}
	return strconv.Itoa(i)
}

func StringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil // Return 0 if the string is empty
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error converting string '%s' to int: %v", s, err)
		return 0, err // Return 0 and the error if conversion fails
	}
	return i, nil
}

// / Charting Fiddling
type DataPoint struct {
	Time  time.Time
	Value float64
}

type ScatterData struct {
	X    []string  `json:"x"`
	Y    []float64 `json:"y"`
	Type string    `json:"type"`
}

func GenerateScatterData(points []DataPoint) (ScatterData, error) {
	var x []string
	var y []float64

	for _, point := range points {
		x = append(x, point.Time.Format("2006-01-02 15:04:05")) // Match the required format
		y = append(y, point.Value)
	}

	return ScatterData{
		X:    x,
		Y:    y,
		Type: "scatter",
	}, nil
}

// ReplacePathParam replaces a placeholder like {key} in the path template with the provided value.
func ReplacePathParam(template, key, value string) string {
	placeholder := "{" + key + "}"
	return strings.ReplaceAll(template, placeholder, value)
}
