package Jsonstruct

type ProblemInfo struct {
	Tag          string
	Title        string
	Description  string
	Input        string
	Output       string
	SampleInput  string
	SampleOutput string
	Code         string
}

type DescriptionInfo struct {
	Trustable bool `json:"trustable"`
	Sections  []struct {
		Title string `json:"title"`
		Value struct {
			Format  string `json:"format"`
			Content string `json:"content"`
		} `json:"value"`
	} `json:"sections"`
}

type ReqCodeInfo struct {
	Memory            int    `json:"memory"`
	Code              string `json:"code"`
	StatusType        int    `json:"statusType"`
	Author            string `json:"author"`
	Length            int    `json:"length"`
	Runtime           int    `json:"runtime"`
	Language          string `json:"language"`
	StatusCanonical   string `json:"statusCanonical"`
	AuthorID          int    `json:"authorId"`
	LanguageCanonical string `json:"languageCanonical"`
	ContestID         int    `json:"contestId"`
	SubmitTime        int64  `json:"submitTime"`
	IsOpen            int    `json:"isOpen"`
	ContestNum        string `json:"contestNum"`
	Processing        bool   `json:"processing"`
	RunID             int    `json:"runId"`
	Oj                string `json:"oj"`
	RemoteRunID       string `json:"remoteRunId"`
	ProbNum           string `json:"probNum"`
	Status            string `json:"status"`
}

type CodeQueryInfo struct {
	Memory            int    `json:"memory"`
	Access            int    `json:"access"`
	StatusType        int    `json:"statusType"`
	Runtime           int    `json:"runtime"`
	Language          string `json:"language"`
	StatusCanonical   string `json:"statusCanonical"`
	UserName          string `json:"userName"`
	UserID            int    `json:"userId"`
	LanguageCanonical string `json:"languageCanonical"`
	ContestID         int    `json:"contestId"`
	ContestNum        string `json:"contestNum"`
	Processing        bool   `json:"processing"`
	RunID             int    `json:"runId"`
	Time              int64  `json:"time"`
	ProblemID         int    `json:"problemId"`
	SourceLength      int    `json:"sourceLength"`
	Status            string `json:"status"`
}

type ContestInfo struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Type           int    `json:"type"`
	Openness       int    `json:"openness"`
	AuthStatus     int    `json:"authStatus"`
	Begin          int64  `json:"begin"`
	End            int64  `json:"end"`
	CreateTime     int64  `json:"createTime"`
	Started        bool   `json:"started"`
	Ended          bool   `json:"ended"`
	ManagerID      int    `json:"managerId"`
	ManagerName    string `json:"managerName"`
	GroupShortName string `json:"groupShortName"`
	GroupName      string `json:"groupName"`
	Fav            bool   `json:"fav"`
	Description    struct {
		Format  string `json:"format"`
		Content string `json:"content"`
	} `json:"description"`
	Announcement string `json:"announcement"`
	Problems     []struct {
		Pid               int    `json:"pid"`
		Title             string `json:"title"`
		Oj                string `json:"oj"`
		ProbNum           string `json:"probNum"`
		Num               string `json:"num"`
		PublicDescID      int    `json:"publicDescId"`
		PublicDescVersion int64  `json:"publicDescVersion"`
		Properties        []struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Hint    bool   `json:"hint"`
		} `json:"properties"`
		Weight int `json:"weight"`
	} `json:"problems"`
	ProblemsHash          string        `json:"problemsHash"`
	PrivatePeerContestIds []interface{} `json:"privatePeerContestIds"`
	EnableTimeMachine     bool          `json:"enableTimeMachine"`
	SumTime               bool          `json:"sumTime"`
	Penalty               int           `json:"penalty"`
	PartialScore          bool          `json:"partialScore"`
	CustomizedWeight      bool          `json:"customizedWeight"`
	ShowPeers             bool          `json:"showPeers"`
}
