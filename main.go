package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	snet "net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

// LomoWebVersion version auto generated
const LomoWebVersion = "2019_09_11.23_59_28.0.b5aeabd"

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
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Handlers

// LoginPageHandler for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("templates/login.html")
	fmt.Fprintf(response, body)
}

// ImportPageHandler for GET
func ImportPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("templates/import.html")
	io.WriteString(response, body)
}

// ConfJsTemplate conf.js template
var ConfJsTemplate = `
var CONFIG = {
    SERVICE_URL: '%v',
    LOGIN_URI: 'login',
    ASSERT_URI: 'asset',
    PREVIEW_URI: 'preview',

    getLoginUrl: function() {
        return CONFIG.SERVICE_URL + '/' + CONFIG.LOGIN_URI;
    },

    getUploadUrl: function() {
        return CONFIG.SERVICE_URL + '/' + CONFIG.ASSERT_URI;
    },

    getAssetUrl: function(data) {
        return CONFIG.SERVICE_URL + '/' + CONFIG.ASSERT_URI + '/' + data.result.Name + "?token=" + sessionStorage.getItem("token");
    },

    getPreviewUrl: function(data) {
        return CONFIG.SERVICE_URL + '/' + CONFIG.PREVIEW_URI + '/' + data.result.Name + "?token=" + sessionStorage.getItem("token");
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
	app.Email = "fuji246@gmail.com"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "baseurl, b",
			Usage: "Base url of lomo-backend",
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
			fmt.Printf("error while list ips: %v\n", err)
		} else if len(ipList) > 0 {
			fmt.Printf("ip[0]: %v\n", ipList[0])
			BaseURL = fmt.Sprintf("http://%v:8000", ipList[0])
		}
	}

	if BaseURL == "" {
		return errors.New("invalid baseurl")
	}

	var router = mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	router.HandleFunc("/static/js/conf.js", ConfJsHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/", LoginPageHandler)        // GET
	router.HandleFunc("/import", ImportPageHandler) // GET

	http.Handle("/", router)
	http.ListenAndServe(":80", nil)

	return nil
}
