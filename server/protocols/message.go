package protocols

type Message struct {
	UserName    string `json:"username"`
	GroupName   string `json:"group_name"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
	Login       string `json:"login"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	Status      bool   `json:"-"`
	UserIcon    string `json:"-"`
	Action      string `json:"action"`
}
