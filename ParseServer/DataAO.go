package ParseServer

import (
	"MasterClient/UnityServer"
	"sync"
)

var analyzeData []UnityServer.AnalyzeState
var uploadMutex sync.Mutex
