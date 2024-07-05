package info

type GetAppInfoResponse struct {
	Version   string         `json:"version"`
	BuildTime string         `json:"buildTime"`
	GoVersion string         `json:"goVersion"`
	Git       GitResponse    `json:"git"`
	Server    ServerResponse `json:"server"`
} //@Name GetAppInfoResponse

type GitResponse struct {
	Branch     string `json:"branch"`
	Commit     string `json:"commit"`
	Repository string `json:"repository"`
} //@Name GitResponse

type ServerResponse struct {
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	Hostname  string `json:"hostname"`
	Url       string `json:"url"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
	Interface string `json:"interface"`
	UpTime    string `json:"uptime"`
} //@Name ServerResponse
