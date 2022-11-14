package ParseServer

import (
	"MasterClient/UnityServer"
	"strings"
)

func AcceptData(data string) {
	var uuid string
	current := strings.Split(data, "&")
	for i := 0; i < len(current); i++ {
		if strings.Contains(current[i], "uuid") {
			cdata := strings.Split(current[i], "=")
			uuid = cdata[1]
		}
	}
	UnityServer.DeleteFiles(uuid)
}
