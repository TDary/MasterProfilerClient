package UnityServer

func CheckUnityVersion(data AnalyzeData) string {
	for i := 0; i < len(config.UnityPath); i++ {
		if data.UnityVersion == config.UnityPath[i].Version {
			return config.UnityPath[i].Path
		}
	}
	return ""
}
