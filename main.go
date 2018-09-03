package main


import (
    "net/http"
    "log"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Cluster struct {
    ID            string   `json:"id,omitempty"`
    Clustername   string   `json:"clustername,omitempty"`
    HostList    []Host     `json:"hosts,omitempty"`
}

type Host struct {
    HostType     string   `json:"hosttype,omitempty"`
    HostName     string   `json:"hostname,omitempty"`
    HostVersion  string   `json:"hostversion,omitempty"`
}

var cluster []Cluster


// fun main()
func main() {

    initData()

// Init rest server
    router := mux.NewRouter()
    router.HandleFunc("/cluster", ClusterList).Methods("GET")
    router.HandleFunc("/cluster/{id}", GetCluster).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func ClusterList(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(cluster)
}

func GetCluster(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range cluster {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Cluster{})
}



func initData() {
// Populate variables with data
    var hostlist []Host
    hostlist = append(hostlist, Host{
                                  HostType: "etcd",
                                  HostName: "vmware0392",
                                  HostVersion: "3.3.0" } )
    hostlist = append(hostlist, Host{
                                  HostType: "master",
                                  HostName: "vmware0393",
                                  HostVersion: "1.10.7" })
    cluster = append(cluster, Cluster{ ID: "1", Clustername: "Vestas TSW", HostList: hostlist } )

    hostlist = []Host{}

    hostlist = append(hostlist, Host{
                                  HostType: "etcd",
                                  HostName: "VM032",
                                  HostVersion: "3.3.3" } )
    hostlist = append(hostlist, Host{
                                  HostType: "etcd",
                                  HostName: "VM033",
                                  HostVersion: "3.3.3" } )
    hostlist = append(hostlist, Host{
                                  HostType: "master",
                                  HostName: "VM033",
                                  HostVersion: "1.9.9" })
    cluster = append(cluster, Cluster{ ID: "2", Clustername: "Terma IT", HostList: hostlist } )
}

