package handlers

import (
	"fmt"
	"github.com/go-resty/resty"
	"github.com/hokaccha/go-prettyjson"
	"net"
)

// Performs search by of ip's location via https://www.ipvigilante.com/ API
// ip is the ip to search. second parameter (number of results) is ignored as there is either one result or no result.
func searchIpLocation(ip string, _ int) []*SearchResult {
	result := make([]*SearchResult, 0, 1)

	if !isValidIp4(ip) {
		fmt.Printf("%s is not a valid IPv4 address\n", ip)
		return result
	}

	url := "https://ipvigilante.com/json/" + ip
	resp, err := resty.R().SetQueryParams(map[string]string{}).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		fmt.Printf("Problem has been occured while retrieving IP location:%s\n", err.Error())
		return result
	}

	sr := SearchResult{Extras: make(map[string]string), ExtrasOrder: make([]string, 0, 1)}
	location := "IP location of " + ip + ":"

	value, err := prettyjson.Format(resp.Body())
	if err != nil {
		fmt.Printf("Cannot format JSON:%s", err.Error())
		return result
	}

	sr.Extras[location] = string(value)
	sr.ExtrasOrder = append(sr.ExtrasOrder, location)
	result = append(result, &sr)

	return result
}

func isValidIp4(ip string) bool {
	ip4 := net.ParseIP(ip)
	if ip4.To4() == nil {
		return false
	}

	return true
}
