package ParseServer

import (
	"MasterClient/UnityServer"
	"sync"
)

var analyzeData []UnityServer.AnalyzeState
var logicMutex sync.Mutex
