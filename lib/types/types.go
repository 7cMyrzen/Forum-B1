package types

type Configuration struct {
	MySQL MySQLConfig `json:"mysql"`
	Forum ForumConfig `json:"forum"`
}

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ForumConfig struct {
	Reload      string `json:"reload"`
	ReloadTime  int    `json:"reloadtime"`
	Description string `json:"description"`
}

type Account struct {
	Username string
	Email    string
	Password string
}

type HomePageData struct {
	Authenticated bool
}

type ProfilePageData struct {
	Image         string
	Username      string
	Email         string
	Description   string
	Creation      string
	Authenticated bool
	Nb_Posts      int
	Nb_Likes      int
	Nb_Comments   int
	Posts         []Posts
}

type ForumPageData struct {
	Authenticated bool
	Posts         []Posts
}

type NewTopicData struct {
	Authenticated bool
}

type Posts struct {
	Author    string
	AuthorPic string
	ID        int
	Title     string
	Content   string
	Creation  string
	NbLikes   int
	NbComs    int
	Comments  []Comments
}

type Comments struct {
	ID        int
	Content   string
	Author    string
	AuthorPic string
	Creation  string
}

type DeletePostRequest struct {
	TopicID string `json:"postId"`
}

type TopicPageData struct {
	Authenticated bool
	Post          Posts
	Comments      []Comments
}
