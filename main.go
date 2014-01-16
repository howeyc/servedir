// Serve Directory
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type HumanReadableSize int64

func (hrSize HumanReadableSize) String() string {
	unit := ""
	fltSize := float64(hrSize)
	if fltSize > 1024 {
		unit = "K"
		fltSize /= 1024.0
		if fltSize > 1024 {
			unit = "M"
			fltSize /= 1024.0
			if fltSize > 1024 {
				unit = "G"
				fltSize /= 1024.0
			}
		}
	}
	return fmt.Sprintf("%0.1f%s", fltSize, unit)
}

func BrowseDirectory(w http.ResponseWriter, r *http.Request) {
	upath, _ := url.QueryUnescape(mux.Vars(r)["directory"])
	if upath == "" {
		upath = "."
	}
	if pinfo, sErr := os.Stat(upath); sErr != nil {
		http.Error(w, sErr.Error(), http.StatusNotFound)
	} else if pinfo.IsDir() {
		entries, err := ioutil.ReadDir(upath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		fmt.Fprintf(w, "<html><head><title>Index of %s</title></head>\n", upath)
		fmt.Fprintln(w, "<body>")
		fmt.Fprintf(w, "<h1>Index of %s</h1>\n<hr>\n", upath)
		fmt.Fprintln(w, "<table width=\"100%\">")
		fmt.Fprintln(w, "<tr><th width=\"60%\" align=\"left\">Name</th><th width=\"30%\" align=\"left\">Last Modified</th><th width=\"10%\" align=\"left\">Size</th></tr>")
		if upath != "." {
			parent, _ := filepath.Split(upath[:len(upath)-1])
			fmt.Fprintf(w, "<tr><td><a href=\"/%s\">Parent Directory</a></td><td></td><td> - </td></tr>", parent)
		}
		for _, info := range entries {
			hrSize := HumanReadableSize(info.Size())
			if info.IsDir() {
				fmt.Fprintf(w, "<tr><td><a href=\"%s/\">%s/</a></td><td>%s</td><td>%s</td></tr>", info.Name(), info.Name(), info.ModTime().Format("2006-01-02 15:04:05 -0700"), " - ")
			} else {
				fmt.Fprintf(w, "<tr><td><a href=\"%s\">%s</a></td><td>%s</td><td>%s</td></tr>", info.Name(), info.Name(), info.ModTime().Format("2006-01-02 15:04:05 -0700"), hrSize.String())
			}
		}
		fmt.Fprintln(w, "</table><hr>Golang <a href=\"http://github.com/howeyc/servedir\">servedir</a></body></html>")
	} else {
		http.ServeFile(w, r, upath)
	}
}

func main() {
	r := mux.NewRouter()
	r.Path("/{directory:.*}").HandlerFunc(BrowseDirectory)
	http.Handle("/", r)

	var port int
	var localhost bool
	flag.IntVar(&port, "port", 8080, "Port number.")
	flag.BoolVar(&localhost, "localhost", false, "Bind to 127.0.0.1 only.")
	flag.Parse()

	serveIP := ""
	if localhost {
		serveIP = "127.0.0.1"
	}
	http.ListenAndServe(fmt.Sprintf("%s:%d", serveIP, port), nil)
}
