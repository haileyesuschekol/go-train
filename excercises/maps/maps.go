package main

import "fmt"

const (
	online = 0
	offline = 1
	maintenance = 2
	retired = 3
)

func printServerStatus(servers map[string] int){
fmt.Println("\nThere are",len(servers), "servers")
stats := make(map[int] int)
for _, status := range servers{
switch status{
	case online:
		stats[online] +=1
	case offline:
		stats[offline] +=1
	case maintenance:
		stats[maintenance] +=1
	case retired:
		stats[retired] +=1
	default:
	panic("unhandled server status")			
}}

fmt.Println(stats[online], "servers are online")
fmt.Println(stats[offline], "servers are offline")
fmt.Println(stats[maintenance], "servers are maintenance")
fmt.Println(stats[retired], "servers are retired")
}



func main(){
	servers := []string{"darkstar", "aiur", "omicorn","w-359","baseline"}
	serverStatus := make(map[string] int)

	for _,server := range servers{
		serverStatus[server] = online
	}
	 
	printServerStatus(serverStatus)

	serverStatus["darkstar"] = retired
	serverStatus["aiur"] = offline

	printServerStatus(serverStatus)

	for server, _ := range serverStatus{
		serverStatus[server] =maintenance
	}
	printServerStatus(serverStatus)

}