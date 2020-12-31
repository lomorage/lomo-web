package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	snet "net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	rice "github.com/GeertJohan/go.rice"
	"github.com/chai2010/gettext-go"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"golang.org/x/text/language"
)

// LomoWebVersion version auto generated
const LomoWebVersion = "2020-08-16.11-25-18.0.d75c683"

const i18nMessage = `{
	"zh_CN": {
		"LC_MESSAGES": {
			"message.json": [
				{
					"msgctxt"     : "",
					"msgid"       : "Login",
					"msgid_plural": "",
					"msgstr"      : ["登陆"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Login form",
					"msgid_plural": "",
					"msgstr"      : ["登陆表单"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Username",
					"msgid_plural": "",
					"msgstr"      : ["用户名"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Password",
					"msgid_plural": "",
					"msgstr"      : ["密码"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Submit",
					"msgid_plural": "",
					"msgstr"      : ["提交"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Gallery",
					"msgid_plural": "",
					"msgstr"      : ["图库"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Inbox",
					"msgid_plural": "",
					"msgstr"      : ["收件箱"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Import",
					"msgid_plural": "",
					"msgstr"      : ["导入"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Logout",
					"msgid_plural": "",
					"msgstr"      : ["退出"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Lomorage Gallery",
					"msgid_plural": "",
					"msgstr"      : ["Lomorage图库"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Assets Import",
					"msgid_plural": "",
					"msgstr"      : ["资源导入"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Drag and drop image and video files or directory to import them into Lomorage",
					"msgid_plural": "",
					"msgstr"      : ["要导入图片或视频到Lomorage，请将其拖拽到下面区域"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Add files...",
					"msgid_plural": "",
					"msgstr"      : ["添加文件..."]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Cancel upload",
					"msgid_plural": "",
					"msgstr"      : ["取消上传"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Delete selected",
					"msgid_plural": "",
					"msgstr"      : ["删除选择"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Processing...",
					"msgid_plural": "",
					"msgstr"      : ["处理中..."]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Edit",
					"msgid_plural": "",
					"msgstr"      : ["编辑"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Start",
					"msgid_plural": "",
					"msgstr"      : ["开始"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Cancel",
					"msgid_plural": "",
					"msgstr"      : ["取消"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Error",
					"msgid_plural": "",
					"msgstr"      : ["错误"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Delete",
					"msgid_plural": "",
					"msgstr"      : ["删除"]
				}
			]
		}
	},

	"en_US": {
		"LC_MESSAGES": {
			"message.json": [
				{
					"msgctxt"     : "",
					"msgid"       : "Login",
					"msgid_plural": "",
					"msgstr"      : ["Login"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Login form",
					"msgid_plural": "",
					"msgstr"      : ["Login form"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Username",
					"msgid_plural": "",
					"msgstr"      : ["Username"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Password",
					"msgid_plural": "",
					"msgstr"      : ["Password"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Submit",
					"msgid_plural": "",
					"msgstr"      : ["Submit"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Gallery",
					"msgid_plural": "",
					"msgstr"      : ["Gallery"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Inbox",
					"msgid_plural": "",
					"msgstr"      : ["Inbox"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Import",
					"msgid_plural": "",
					"msgstr"      : ["Import"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Logout",
					"msgid_plural": "",
					"msgstr"      : ["Logout"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Lomorage Gallery",
					"msgid_plural": "",
					"msgstr"      : ["Lomorage Gallery"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Assets Import",
					"msgid_plural": "",
					"msgstr"      : ["Assets Import"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Drag and drop image and video files or directory to import them into Lomorage",
					"msgid_plural": "",
					"msgstr"      : ["Drag and drop image and video files or directory to import them into Lomorage"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Add files...",
					"msgid_plural": "",
					"msgstr"      : ["Add files..."]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Cancel upload",
					"msgid_plural": "",
					"msgstr"      : ["Cancel upload"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Delete selected",
					"msgid_plural": "",
					"msgstr"      : ["Delete selected"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Processing...",
					"msgid_plural": "",
					"msgstr"      : ["Processing..."]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Edit",
					"msgid_plural": "",
					"msgstr"      : ["Edit"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Start",
					"msgid_plural": "",
					"msgstr"      : ["Start"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Cancel",
					"msgid_plural": "",
					"msgstr"      : ["Cancel"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Error",
					"msgid_plural": "",
					"msgstr"      : ["Error"]
				},
				{
					"msgctxt"     : "",
					"msgid"       : "Delete",
					"msgid_plural": "",
					"msgstr"      : ["Delete"]
				}
			]
		}
	}
}`

var gText gettext.Gettexter

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

func changeLocale(locale string) {
	gText.SetLanguage(locale)
}

func translate(input string) string {
	return gText.PGettext("", input)
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

	funcMap := template.FuncMap{
		"gettext": translate,
	}
	t, _ := template.New("foo").Funcs(funcMap).Parse(templateString)
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// ChangePreferedLanguage change lang according to http header
func ChangePreferedLanguage(r *http.Request) {
	var matcher = language.NewMatcher([]language.Tag{
		language.English, // The first language is used as fallback.
		language.Chinese,
	})

	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(matcher, accept)
	if strings.HasPrefix(tag.String(), "zh") {
		changeLocale("zh_CN")
	} else {
		changeLocale("en_US")
	}
}

// Handlers

// LoginPageHandler for GET
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	ChangePreferedLanguage(request)
	var body, _ = LoadFile("login.html")
	io.WriteString(response, body)
}

// ImportPageHandler for GET
func ImportPageHandler(response http.ResponseWriter, request *http.Request) {
	ChangePreferedLanguage(request)
	var body, _ = LoadFile("import.html")
	io.WriteString(response, body)
}

// GalleryPageHandler for GET
func GalleryPageHandler(response http.ResponseWriter, request *http.Request) {
	ChangePreferedLanguage(request)
	var body, _ = LoadFile("gallery.html")
	io.WriteString(response, body)
}

// InboxPageHandler for GET
func InboxPageHandler(response http.ResponseWriter, request *http.Request) {
	ChangePreferedLanguage(request)
	var body, _ = LoadFile("inbox.html")
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
	RECEIVE_URI: 'receive',
	RECEIVE_USER_URI: 'receive/user',
	RECEIVE_GROUP_URI: 'receive/group',
	RECEIVE_ASSERT_URI: 'receive/asset',
	RECEIVE_PREVIEW_URI: 'receive/preview',

	getServiceUrl: function() {
		if (CONFIG.SERVICE_URL === "") {
			return window.location.protocol + "//" + window.location.hostname + ":8000";
		} else {
			return CONFIG.SERVICE_URL;
		}
	},

    getLoginUrl: function() {
        return CONFIG.getServiceUrl() + '/' + CONFIG.LOGIN_URI;
    },

    getUploadUrl: function() {
        return CONFIG.getServiceUrl() + '/' + CONFIG.ASSERT_URI;
    },

    getAssetUrl: function(name) {
        return CONFIG.getServiceUrl() + '/' + CONFIG.ASSERT_URI + '/' + name + "?token=" + sessionStorage.getItem("token") + "&orig=1";
    },

    getPreviewUrl: function(name) {
        return CONFIG.getServiceUrl() + '/' + CONFIG.PREVIEW_URI + '/' + name + "?token=" + sessionStorage.getItem("token");
	},

	getMonthLevelMerkleTreeUrl: function() {
		return CONFIG.getServiceUrl() + '/' + CONFIG.CATEGORY_URI;
	},

	getAssetLevelMerkleTreeUrl: function(year, month) {
		return CONFIG.getServiceUrl() + '/' + CONFIG.CATEGORY_URI + '/' + year + '/' + month;
	},

	getInboxUrl: function() {
		return CONFIG.getServiceUrl() + '/' + CONFIG.RECEIVE_URI + "?token=" + sessionStorage.getItem("token");
	},

	getUserInboxUrl: function(uid) {
		return CONFIG.getServiceUrl() + '/' + CONFIG.RECEIVE_USER_URI + '/' + uid + "?token=" + sessionStorage.getItem("token");
	},

	getGroupInboxUrl: function(gid) {
		return CONFIG.getServiceUrl() + '/' + CONFIG.RECEIVE_GROUP_URI + '/' + gid + "?token=" + sessionStorage.getItem("token");
	},

	getInboxAssetUrl: function(shareid) {
        return CONFIG.getServiceUrl() + '/' + CONFIG.RECEIVE_ASSERT_URI + '/' + shareid + "?token=" + sessionStorage.getItem("token") + "&orig=1";
    },

    getInboxPreviewUrl: function(shareid) {
        return CONFIG.getServiceUrl() + '/' + CONFIG.RECEIVE_PREVIEW_URI + '/' + shareid + "?token=" + sessionStorage.getItem("token");
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

	app.Authors = []cli.Author{
		cli.Author{
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

	log.Printf("Lomorage Service lomod url: %s", BaseURL)

	gText = gettext.New("message", "", i18nMessage).SetLanguage("en_US")

	var router = mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	router.HandleFunc("/static/lomo/js/conf.js", ConfJsHandler)

	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.PathPrefix("/static/").Handler(staticFileServer)
	router.HandleFunc("/", LoginPageHandler)          // GET
	router.HandleFunc("/import", ImportPageHandler)   // GET
	router.HandleFunc("/gallery", GalleryPageHandler) // GET
	router.HandleFunc("/inbox", InboxPageHandler)     // GET

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
