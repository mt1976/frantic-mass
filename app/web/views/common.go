package views

import (
	"fmt"
	"time"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/application"
	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/dateHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/stringHelpers"
)

type Common struct {
	Title         string
	Description   string
	Keywords      string
	Author        string
	Version       string
	Theme         string
	Logo          string
	Favicon       string
	Errors        []string
	Messages      []string
	Notifications []string

	Success           bool
	Status            int    // HTTP status code
	CurrentUserID     int    // ID of the currently logged-in user
	CurrentUserName   string // Name of the currently logged-in user
	CurrentUserAvatar string // Avatar URL of the currently logged-in user
	SessionID         string // Session ID for tracking user sessions

	DateTime        string // Current date and time
	DateString      string // Current date and time in string format
	Host            string // Hostname or domain of the application
	IP              string // IP address of the server
	Identity        string // Unique identifier for the system or application
	ApplicationName string // Name of the application

	Prefix      string // Prefix for the application, used in URLs
	Environment string // Environment in which the application is running (e.g., development, production)
	ReleaseDate string // Release date of the application
	Copyright   string // Copyright information for the application
	License     string // License type for the application
	Locale      string // Locale for the application, used for internationalization
	BuildTime   string // Build time of the application
	Delimiter   string // Delimiter used in the application, if applicable
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
	godump.Dump(cache, "Cache Configuration", cacheChecksum)
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

func (c *Common) SetDefaults() {

	// Application Config

	c.Title = cache.GetApplication_Name()
	if c.Title == "" {
		c.Title = "Frantic Mass"
	}
	c.Description = cache.GetApplication_Description()
	if c.Description == "" {
		c.Description = "Frantic Mass Management Application"
	}
	c.Keywords = "mass, management"

	c.Author = cache.GetApplication_Author()
	if c.Author == "" {
		c.Author = "Frantic Mass Team"
	}
	c.Version = cache.GetApplication_Version()
	if c.Version == "" {
		c.Version = "1.0.0"
	}
	c.Theme = "default"
	c.Logo = "/static/images/logo.png"
	c.Favicon = "/static/images/favicon.ico"
	c.Errors = []string{}
	c.Messages = []string{}
	c.Notifications = []string{}

	c.Success = true
	c.Status = 200 // OK
	c.CurrentUserID = 0
	c.CurrentUserName = "Guest"
	c.CurrentUserAvatar = "/static/images/default-avatar.png"
	c.SessionID = ""
	c.DateTime = time.Now().Format(time.RFC3339)
	c.DateString = dateHelpers.FormatHuman(time.Now())
	c.Host = application.HostName()
	c.IP = application.HostIP()
	c.Identity = application.SystemIdentity()
	c.ApplicationName = cache.GetApplication_Name()
	if c.ApplicationName == "" {
		c.ApplicationName = "Frantic Mass"
	}

	c.Prefix = cache.GetApplication_Prefix()
	if c.Prefix == "" {
		c.Prefix = "/"
	}
	c.Environment = cache.GetApplication_Environment()
	if c.Environment == "" {
		c.Environment = "development"
	}
	c.ReleaseDate = cache.GetApplication_ReleaseDate()
	if c.ReleaseDate == "" {
		c.ReleaseDate = time.Now().Format("2006-01-02")
	}
	c.Copyright = cache.GetApplication_Copyright()
	if c.Copyright == "" {
		c.Copyright = "© 2023 Frantic Mass Team"
	}
	c.License = cache.GetApplication_License()
	if c.License == "" {
		c.License = "MIT"
	}
	c.Locale = cache.GetApplication_Locale()
	if c.Locale == "" {
		c.Locale = "en_GB"
	}
	c.BuildTime = time.Now().Format(time.RFC3339)
	c.Delimiter = cache.Delimiter()
}

func (c *Common) AddError(err string) {
	c.Errors = append(c.Errors, err)
	c.Success = false
}

func (c *Common) AddMessage(msg string) {
	c.Messages = append(c.Messages, msg)
}

func (c *Common) AddNotification(notification string) {
	c.Notifications = append(c.Notifications, notification)
}
