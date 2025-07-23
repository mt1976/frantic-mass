# common.toml Configuration Guide

This file (`common.toml`) contains the main configuration for the Frantic Mass application. It defines application metadata, database settings, server parameters, translation options, asset references, date formats, host definitions, security/session settings, messaging keys, display options, status codes, communication integrations, and logging preferences.

## Sections Overview

### [Application]
- General metadata: name, version, description, environment, author, license, locale, etc.

### [Database]
- Database type (e.g., storm), version, and connection parameters.

### [Server]
- Host, port, protocol, and environment for the web server.

### [Translation]
- Locale, host, port, protocol, and permitted origins/locales for translation features.

### [Assets]
- Logo and favicon file names.

### [Dates]
- Date and time format strings for display, backup, and internal use.

### [History]
- Maximum number of history entries to retain.

### [Hosts]
- List of known hosts with name, FQDN, IP, and zone.

### [Security]
- Session expiry, session key names, and service user details.

### [Message]
- Message key names for inter-component communication.

### [Display]
- UI display delimiter.

### [Status]
- Status codes for unknown, online, offline, error, and warning states.

### [Communications]
- Pushover and Email integration settings (API keys, SMTP server, etc.).

### [Logging]
- Enable/disable logging for various subsystems, and log file rotation settings.

## Usage
- Edit this file to configure the Frantic Mass application for your environment.
- Sensitive information (API keys, passwords) should be protected and not shared publicly.
- Restart the application after making changes for them to take effect.

## Example
```toml
[Application]
name = "frantic-mass"
version = "0.0.1"
...

[Database]
type = "storm"
...

[Server]
host = "127.0.0.1"
port = 5059
...
```

For more details on each section, refer to the inline comments or documentation.
