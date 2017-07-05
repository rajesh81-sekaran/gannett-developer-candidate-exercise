package common

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// HandleRoot Handler function for /
func HandleRoot(w http.ResponseWriter, req *http.Request) {
	var pid ProfileIDStruct
	var useProfileID int
	var arrays AllArticles
	defer req.Body.Close()

	userNameNC := req.FormValue("user")
	if len(userNameNC) == 0 {
		tmplUser, errUser := template.ParseFiles("common/user.html")
		if errUser != nil {
			msg := fmt.Sprintf("Error %s while parsing the file\n", errUser.Error())
			fmt.Printf("%s", msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		errUser = tmplUser.ExecuteTemplate(w, "user.html", nil)
		if errUser != nil {
			msg := fmt.Sprintf("Error %s in tmplUser.Execute\n", errUser.Error())
			fmt.Printf("%s", msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		return
	}

	//Converting user name string to lower case, case insensitive
	userName := strings.ToLower(userNameNC)

	UsersLock.RLock()
	userVal, present := AllUsers[userName]
	if present && AllUsers[userName].timeExpiration.Sub(time.Now()).Nanoseconds() >= 0 {
		useProfileID = userVal.profileID
		UsersLock.RUnlock()
	} else {
		UsersLock.RUnlock()
		responsePID, err := http.Get(ProfileURL)
		if err != nil {
			msg := fmt.Sprintf("Error %s in http.Get\n", err.Error())
			fmt.Printf("%s", msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		defer responsePID.Body.Close()
		if err = json.NewDecoder(responsePID.Body).Decode(&pid); err != nil {
			msg := fmt.Sprintf("Error %s while decoding json\n", err.Error())
			fmt.Printf("%s", msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		useProfileID = pid.ProfileID
		UsersLock.Lock()
		t := time.Now()

		AllUsers[userName] = SingleUser{useProfileID, t, t.AddDate(0, 0, 365)}
		UsersLock.Unlock()
	}

	contentURL := fmt.Sprintf("%s%d", ContentURL, useProfileID)
	responseContents, err := http.Get(contentURL)
	if err != nil {
		msg := fmt.Sprintf("Error %s in http.Get\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	defer responseContents.Body.Close()
	if err = json.NewDecoder(responseContents.Body).Decode(&arrays); err != nil {
		msg := fmt.Sprintf("Error %s while decoding json\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	tmpl, err := template.New("").Funcs(fm).ParseFiles("common/data.html")
	if err != nil {
		msg := fmt.Sprintf("Error %s while parsing the file\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "data.html", arrays)
	if err != nil {
		msg := fmt.Sprintf("Error %s in tmpl.Execute\n", err.Error())
		fmt.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
}
