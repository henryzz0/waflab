// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package operator

import (
	"bufio"
	"net"
	"os"
	"path"
	"strings"

	"github.com/waflab/waflab/autogen/utils"
)

func randomIPfromNetworkSegments(segments []string) (string, error) {
	ipString := segments[utils.RandomIntWithRange(0, len(segments))]

	// Since @ipMatch takes both ip address and ip segments, we can return ip address directly
	// if the ipString is an ip address instead of segment
	if !strings.Contains(ipString, "/") {
		return ipString, nil
	}

	ipAddr, ipNet, err := net.ParseCIDR(ipString)
	if err != nil {
		return "", err
	}

	// randomly generate an IP address from choosen network segment
	for i := len(ipAddr) - 1; i >= 0; i-- {
		if ipNet.Mask[i] == 255 {
			break
		}
		ipAddr[i] = byte(utils.RandomIntWithRange(0, 256-int(ipNet.Mask[i]))) // generate within range [0, 255-Mask]
	}

	return ipAddr.String(), nil
}

func reverseIPMatch(argument string, not bool) (string, error) {
	// pick a random ip segments from argument
	networkSegments := strings.Split(argument, ",")
	return randomIPfromNetworkSegments(networkSegments)
}

func reverseIPMatchFromFile(argument string, not bool) (string, error) {
	// According to ModSecurity Handbook
	// Same as @ipMatch, but uses a file that contains one IP address or network segment per line.
	file, err := os.Open(argument)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	segments := []string{}

	for scanner.Scan() {
		segments = append(segments, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return randomIPfromNetworkSegments(segments)
}

func pickRandomLineFromFile(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return utils.PickRandomString(text), nil
}

func reverseDetectSQLi(argument string, not bool) (string, error) {
	return pickRandomLineFromFile(path.Join("autogen", "operator", "data", "Generic-SQLi.txt"))
}

func reverseDetectXSS(argument string, not bool) (string, error) {
	return pickRandomLineFromFile(path.Join("autogen", "operator", "data", "XSS-BruteLogic.txt"))
}
