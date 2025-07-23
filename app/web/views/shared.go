package views

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/stringHelpers"
	"github.com/mt1976/frantic-mass/app/web/helpers"
)

type AppContext struct {
	PageTitle    string
	PageSummary  string
	PageKeywords string

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

	UserErrorMessages []string
	UserMessages      []string
	UserAlerts        []string

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
}

var cache *commonConfig.Settings
var cacheChecksum string

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
	c.WasSuccessful = false
}

func (c *AppContext) AddMessage(msg string) {
	c.UserMessages = append(c.UserMessages, msg)
}

func (c *AppContext) AddNotification(notification string) {
	c.UserAlerts = append(c.UserAlerts, notification)
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
