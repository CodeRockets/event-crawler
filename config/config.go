package config

import (
	"github.com/stvp/go-toml-config"
)

var (
	FsExploreEndpoint = config.String("fs_explore_endpoint", "")
	FsGetEndpoint     = config.String("fs_get_endpoint", "")
	FsClientId        = config.String("foursquare.client_id", "")
	FsClientSecret    = config.String("foursquare.client_secret", "")
	DbConStr          = config.String("database.conn_str", "")
	BxIdCookie        = config.String("biletix.cookie_pass", "")
	BxDomainRoot      = config.String("biletix.domain_root", "")
	BxIdCookieName    = config.String("biletix.cookie_name", "")
	BxDomain          = config.String("biletix.domain", "")
	GoogleCSEEndpoint = config.String("google.cse_endpoint", "")
	GoogleCSEApiKey   = config.String("google.cse_apikey", "")
	GoogleCSEAppCsx   = config.String("google.cse_appcsx", "")
)
