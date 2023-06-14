/*
Some nooding to convert an out-of-band set of OpenVPN profiles into the
location + format for AWS VPN to pick up.

The AWS VPN isn't great and requires walking through a basic wizard for each
profile.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

type connectionProfilesTopLevel struct {
	Version                  string
	LastSelectedProfileIndex int
	ConnectionProfiles       []connectionProfile
}

type connectionProfile struct {
	ProfileName          string
	OvpnConfigFilePath   string
	CvpnEndpointId       string
	CvpnEndpointRegion   string
	CompatibilityVersion string
	FederatedAuthType    int
}

func main() {
	/*
		Aim is to generate the AWS `~/.config/AWSVPNClient/ConnectionsProfile`
		so that I can just drop all AWS profiles from: `</path/to/vpn/configs/>`
		into: `~/.config/AWSVPNClient/OpenVpnConfigs` + remove the `.ovpn`
		extension.
	*/

	var connectionProfiles []connectionProfile

	reCvpn := regexp.MustCompile(`cvpn-endpoint-\w+`)
	userHomeDir, _ := os.UserHomeDir()
	ovpnConfigDir := path.Join(userHomeDir, ".config/AWSVPNClient/OpenVpnConfigs/")
	ovpnConfigFilePaths, err := os.ReadDir(ovpnConfigDir)
	fmt.Printf("%s, error: %s", ovpnConfigFilePaths, err)
	for _, file := range ovpnConfigFilePaths {
		if file.IsDir() {
			fmt.Printf("%s is a directory.", file.Name())
			continue
		}
		tmpFile := path.Join(ovpnConfigDir, file.Name())
		content, _ := ioutil.ReadFile(tmpFile)
		cvpnEndpointId := string(reCvpn.Find(content))
		tmpConnectionProfile := connectionProfile{
			ProfileName:          file.Name(),
			OvpnConfigFilePath:   tmpFile,
			CvpnEndpointId:       cvpnEndpointId,
			CvpnEndpointRegion:   "eu-west-1",
			CompatibilityVersion: "2",
			FederatedAuthType:    1,
		}
		connectionProfiles = append(connectionProfiles, tmpConnectionProfile)
	}
	final := &connectionProfilesTopLevel{
		Version:                  "1",
		LastSelectedProfileIndex: 1,
		ConnectionProfiles:       connectionProfiles,
	}
	// res_json, _ := json.Marshal(res)
	res_json, _ := json.MarshalIndent(final, "", "  ") // Pretty Print
	fmt.Println(string(res_json))

	fmt.Println("Done")
	// TODO: Write JSON to a file.
}
