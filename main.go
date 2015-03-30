package main

import (
	"database/sql"
	"flag"
	"fmt"
	c "github.com/event-crawler/config"
	s "github.com/event-crawler/strategy"
	u "github.com/event-crawler/util"
	_ "github.com/lib/pq"
	"github.com/stvp/go-toml-config"
	"os"
)

const (
	DEV_CONFIG_PATH  = "config/dev.conf"
	PROD_CONFIG_PATH = "config/prod.conf"
)

func main() {

	config.Parse(PROD_CONFIG_PATH)

	db, _ := sql.Open("postgres", *c.DbConStr)

	initPtr := flag.Bool("init", false, "Creates database tables if not exists.")
	fsgetPtr := flag.String("fsget", "foursquareid", "Foursquare id for getting venue info.")
	bxslinkPtr := flag.String("bxslink", "bxvenuelink", "Biletix venue link. Should be used with fsget.")
	gcsePtr := flag.Bool("gcse", false, "Use Google CSE?")
	fbpagePtr := flag.String("fbpage", "fbpageid", "Facebook id for getting venue info.")

	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "-help")
	}

	flag.Parse()
	strategyflag := false
	msgs := make(chan string)
	quit := make(chan string)
	switch {
	case *initPtr:
		retmsg := "Database tables created successfully."
		err := u.CreateTables(db)
		if err != nil {
			retmsg = fmt.Sprintf("Error: %s", err)
		}
		fmt.Println(retmsg)
		return

	case *fsgetPtr != "foursquareid" && *bxslinkPtr != "bxvenuelink":
		fmt.Println("Manuel Biletix Strategy Started")
		strategyflag = true
		go s.FsGetWithBxLink(msgs, quit, db, *fsgetPtr, *bxslinkPtr)

	case *fsgetPtr != "foursquareid" && *gcsePtr:
		fmt.Println("Google Search Engine Biletix Strategy Started")
		strategyflag = true
		go s.FsGetWithGCSE(msgs, quit, db, *fsgetPtr)
	case *fbpagePtr != "fbpageid":
		fmt.Println("Manuel Biletix Strategy Started")
		strategyflag = true
		go s.FbCallPageApi(msgs, quit, db, *fbpagePtr)
	}
	if strategyflag {
		for {
			select {
			case s := <-msgs:
				fmt.Printf("'%s'\n", s)
			case err := <-quit:
				fmt.Printf("'%s'\n", err)
				return
			}
		}
	}
	return

	fmt.Println("Need help? please use -help flag")
}
