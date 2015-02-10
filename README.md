Event-Crawler is a command line application that collects venue and event data by combining some of web services and applications. 

## SetUp

 - Create a postgresql database
 - Create a TOML configuration file under `config` folder named `prod.conf` using template below.```
```
# Global vars
fs_explore_endpoint = "https://api.foursquare.com/v2/venues/explore?mode=url"
fs_get_endpoint = "https://api.foursquare.com/v2/venues/"

# Database (replace with your own connection string)
[database] 
conn_str = "postgres://username:password@server:port/database"

# Foursquare (replace with your own values)
[foursquare]
client_id = "FOURSQUARECLIENTID"
client_secret = "FOURSQUARECLIENTSECRET"

# Biletix(Make a manual request and replace cookie_pass with it)
[biletix] 
cookie_name = "BXID"
cookie_pass = "AAAAAAWVXpUmkwqSGPBoLkqNMmzPrhweUSjn8Y8U+RxUn7CL0A=="
domain_root = "www.biletix.com"
domain = "http://www.biletix.com"

# google (replace cse_apikey and cse_appcsx with your own values)
[google] 
cse_endpoint = "https://www.googleapis.com/customsearch/v1?"
cse_apikey = "Google_CustomSearchEngine_ApiKey"
cse_appcsx = "Google_CustomSearchEngine:appcsx"

```
 - Build project (windows)
```
 $ go build -o event-crawler.exe main.go
```

## Initialization
After build, you should create db tables with -init flag
```
$ event-crawler -init
```
## Usage
There are two main strategy for now
### 1. Foursquare - Biletix manual mashup
You should give foursquare id and biletix link of the venue as argument. Event-crawler will get venue information from both two services and collect all events from biletix.

```
//Sample Venue;
//TIM Show Center foursquare id: 4b6ee07bf964a52096ce2ce3
//TIM Show Center biletix link: http://www.biletix.com/mekan/TM/TURKIYE/tr
$ event-crawler -fsget=4b6ee07bf964a52096ce2ce3 -bxslink=http://www.biletix.com/mekan/TM/TURKIYE/tr

//OUTPUT
//Manuel Biletix Strategy Started
//'TİM Show Center venue info has collected from foursquare'
//'TİM Show Center inserted into database'
//'TİM Show Center biletix link updated manually'
//'TİM Show Center venue info has collected from biletix'
//'TİM Show Center venue info updated'
//'      Lukomorie Müzikali event info crawled'
//'      Lukomorie Müzikali event info inserted into database'
//'      Lukomorie Müzikali event info crawled'
//'      .....'
//'      .....'
//'      .....'
//'...DONE...'
```

### 2. Foursquare - Google Custom Search - Biletix mashup
This strategy  almost same with first one. But this time you don't need to give biletix link of the venue as argument. Event-crawler will find the biletix link automatically.
```
//Sample Venue;
//TIM Show Center foursquare id: 4b6ee07bf964a52096ce2ce3
//TIM Show Center biletix link: http://www.biletix.com/mekan/TM/TURKIYE/tr
$ event-crawler -fsget=4b6ee07bf964a52096ce2ce3 -gcse

//OUTPUT
//Manuel Biletix Strategy Started
//'TİM Show Center venue info has collected from foursquare'
//'TİM Show Center inserted into database'
//'TİM Show Center biletix link updated manually'
//'TİM Show Center venue info has collected from biletix'
//'TİM Show Center venue info updated'
//'      Lukomorie Müzikali event info crawled'
//'      Lukomorie Müzikali event info inserted into database'
//'      Lukomorie Müzikali event info crawled'
//'      .....'
//'      .....'
//'      .....'
//'...DONE...'
```

