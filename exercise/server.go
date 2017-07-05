package main

import (
	"fmt"
	cmn "exercise/common"
	"net/http"
	"os"
	"strings"
	"sync"
)

func init() {
	cmn.AllUsers = map[string]cmn.SingleUser{}
	cmn.UsersLock = sync.RWMutex{}

	cmn.Port = os.Getenv("SERVER_PORT")
	if len(cmn.Port) == 0 {
		cmn.Port = cmn.DefaultPort
	}

	cmn.ProfileURL = os.Getenv("PROFILE_URL")
	if len(cmn.ProfileURL) == 0 {
		cmn.ProfileURL = cmn.DefaultProfileURL
	} else {
		if !strings.HasSuffix(cmn.ProfileURL, "/") {
			cmn.ProfileURL = cmn.ProfileURL + "/"
		}
	}

	cmn.ContentURL = os.Getenv("CONTENT_URL")
	if len(cmn.ContentURL) == 0 {
		cmn.ContentURL = cmn.DefaultContentURL
	} else {
		if !strings.HasSuffix(cmn.ContentURL, "/") {
			cmn.ContentURL = cmn.ContentURL + "/"
		}
	}
}
func main() {
	
	fmt.Println("Starting Server on Port:", cmn.Port)
	fmt.Println("Profile URL:", cmn.ProfileURL)
	fmt.Println("Content URL:", cmn.ContentURL)
	http.HandleFunc("/", cmn.HandleRoot)
	http.ListenAndServe(":"+cmn.Port, nil)
}
