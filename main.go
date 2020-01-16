package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	snet "net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

// LomoWebVersion version auto generated
const LomoWebVersion = "2020-01-15.21-58-04.0.cb4d26f"

// ListIPs list available ipv4 addresses
func ListIPs() ([]snet.IP, error) {
	addrs, err := snet.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	ret := []snet.IP{}
	for _, a := range addrs {
		if ipnet, ok := a.(*snet.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				ret = append(ret, ipnet.IP)
			}
		}
	}
	return ret, nil
}

// LoadFile load html file
func LoadFile(fileName string) (string, error) {
	// find a rice.Box
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Printf("error finding templates: %v\n", err)
		return "", err
	}
	// get file contents as string
	templateString, err := templateBox.String(fileName)
	if err != nil {
		log.Printf("error reading templates: %v\n", err)
		return "", err
	}

	return templateString, nil
}

// Handlers

// LoginPageHandler for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("login.html")
	io.WriteString(response, body)
}

// ImportPageHandler for GET
func ImportPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("import.html")
	io.WriteString(response, body)
}

// GalleryPageHandler for GET
func GalleryPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("gallery.html")
	io.WriteString(response, body)
}

// ConfJsTemplate conf.js template
var ConfJsTemplate = `
var CONFIG = {
    SERVICE_URL: '%v',
    LOGIN_URI: 'login',
    ASSERT_URI: 'asset',
	PREVIEW_URI: 'preview',
	CATEGORY_URI: 'category',

    getLoginUrl: function() {
        return CONFIG.SERVICE_URL + '/' + CONFIG.LOGIN_URI;
    },

    getUploadUrl: function() {
        return CONFIG.SERVICE_URL + '/' + CONFIG.ASSERT_URI;
    },

    getAssetUrl: function(name) {
        return CONFIG.SERVICE_URL + '/' + CONFIG.ASSERT_URI + '/' + name + "?token=" + sessionStorage.getItem("token");
    },

    getPreviewUrl: function(name) {
        return CONFIG.SERVICE_URL + '/' + CONFIG.PREVIEW_URI + '/' + name + "?token=" + sessionStorage.getItem("token");
	},

	getMonthLevelMerkleTreeUrl: function() {
		return CONFIG.SERVICE_URL + '/' + CONFIG.CATEGORY_URI;
	},

	getAssetLevelMerkleTreeUrl: function(year, month) {
		return CONFIG.SERVICE_URL + '/' + CONFIG.CATEGORY_URI + '/' + year + '/' + month;
	}
}
`

// BaseURL base url from command line or local ip address
var BaseURL = ""

// ConfJsHandler server conf.js
func ConfJsHandler(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, fmt.Sprintf(ConfJsTemplate, BaseURL))
}

func main() {

	app := cli.NewApp()

	app.Version = LomoWebVersion
	app.Usage = "Lomorage web app"

	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Jeromy Fu",
			Email: "fuji246@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "baseurl",
			Usage: "URL of Lomorage Service (lomod), default will be http://[lomoweb-host-ip]:8000",
		},

		&cli.UintFlag{
			Name:  "baseport",
			Usage: "lomod listen port",
			Value: 8000,
		},

		&cli.UintFlag{
			Name:  "port",
			Usage: "lomo-web listen port",
			Value: 80,
		},
	}

	app.Action = bootService

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func bootService(ctx *cli.Context) error {
	BaseURL = ctx.String("baseurl")

	if BaseURL == "" {
		ipList, err := ListIPs()
		if err != nil {
			log.Printf("error while list ips: %v\n", err)
		} else if len(ipList) > 0 {
			log.Printf("ip[0]: %v\n", ipList[0])
			BaseURL = fmt.Sprintf("http://%v:%v", ipList[0], ctx.String("baseport"))
		}
	}

	if BaseURL == "" {
		return errors.New("invalid baseurl")
	}
	log.Printf("Lomorage Service lomod url: %s", BaseURL)

	var router = mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	router.HandleFunc("/static/lomo/js/conf.js", ConfJsHandler)

	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.PathPrefix("/static/").Handler(staticFileServer)
	router.HandleFunc("/", LoginPageHandler)          // GET
	router.HandleFunc("/import", ImportPageHandler)   // GET
	router.HandleFunc("/gallery", GalleryPageHandler) // GET

	log.Println("Server started. Press Ctrl-C to stop server")

	// Catch the Ctrl-C and SIGTERM from kill command
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <-ch
		signal.Stop(ch)
		log.Println("Exit command received. Exiting...")

		// this is a good place to flush everything to disk
		// before terminating.
		log.Println("Signal type : ", signalType)
		os.Exit(0)
	}()

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", ctx.Int("port")), router))

	return nil
}
