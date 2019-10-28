package service
import (
    "myService/crawlads/adsservice/dbclient"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    //"bytes"
)
var isHealthy = true 

func ScoreSort(w http.ResponseWriter, r *http.Request) {
    //Read the 'accountId' path parameter from the mux map
    var t = mux.Vars(r)["num"]
    var last = mux.Vars(r)["last"]
    if last == "" {
        t := time.Now()
        last = t.Format("2006-01")+"-01 00:00:00"
    }


    o := orm.NewOrm()

    var account []orm.Params

    sql := "SELECT * FROM book_info where last_update>'" +last+ "' order by score*score_count limit " + t

    // Read the account struct BoltDB
    o.Raw(sql).Values(&account)

    // If found, marshal into JSON, write headers and content
    data , _ := json.Marshal(account)

    //tmp := [][]byte{ []byte(t),json }
    //data := bytes.Join(tmp, []byte(""))

    writeJsonResponse(w, http.StatusOK, data)
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
    // Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection
    dbUp := dbclient.DbCheck()
    if dbUp && isHealthy {
        data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
        writeJsonResponse(w, http.StatusOK, data)
    } else {
        data, _ := json.Marshal(healthCheckResponse{Status: "Database unaccessible"})
        writeJsonResponse(w, http.StatusServiceUnavailable, data)
    }
}


func SetHealthyState(w http.ResponseWriter, r *http.Request) {
    // Read the 'state' path parameter from the mux map and convert to a bool
    var state, err = strconv.ParseBool(mux.Vars(r)["state"])
    
    // If we couldn't parse the state param, return a HTTP 400
    if err != nil {
        fmt.Println("Invalid request to SetHealthyState, allowed values are true or false")
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    // Otherwise, mutate the package scoped "isHealthy" variable.
    isHealthy = state
    w.WriteHeader(http.StatusOK)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Content-Length", strconv.Itoa(len(data)))
    w.WriteHeader(status)
    w.Write(data)
}

type healthCheckResponse struct {
    Status string `json:"status"`
}
