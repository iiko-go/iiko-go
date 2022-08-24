package iiko

import "time"

// BaseURL is the base URL of iikoCloud API.
const BaseURL string = "https://api-ru.iiko.services"

// DefaultTimeout is the default timeout for iikoCloud API Requests.
// For each iiko client a custom timeout can be setted by calling client.SetTimeout(time.Duration).
// For each API request a custom timeout can be setted by putting iiko.WithCustomTimeout(time.Duration) option after all args.
const DefaultTimeout = 15 * time.Second

// IikoTimeLayout is the date-time format using by iiko.
const IikoTimeLayout = "2006-01-02 15:04:05"

// DefaultRefreshTokenInterval is the default timeout for refreshing iikoCloud API Token.
// For each iiko client a custom timeout can be setted by calling client.SetRefreshTokenTimeout(time.Duration).
const DefaultRefreshTokenInterval = 45 * time.Minute
