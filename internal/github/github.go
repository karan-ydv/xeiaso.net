package github

import (
	"bytes"
	"strconv"
	"time"
)

// User was autogenerated by go generate. To see more details about this
// payload type visit https://developer.github.com/v3/activity/events/types.
type User struct {
	AvatarURL         string `json:"avatar_url"`
	Email             string `json:"email"`
	EventsURL         string `json:"events_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	GravatarID        string `json:"gravatar_id"`
	HTMLURL           string `json:"html_url"`
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	OrganizationsURL  string `json:"organizations_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ReposURL          string `json:"repos_url"`
	SiteAdmin         bool   `json:"site_admin"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	Type              string `json:"type"`
	URL               string `json:"url"`
	Username          string `json:"username"`
}

// Commits was autogenerated by go generate. To see more details about this
// payload type visit https://developer.github.com/v3/activity/events/types.
type Commits struct {
	Added     []string `json:"added"`
	Author    User     `json:"author"`
	Committer User     `json:"committer"`
	Distinct  bool     `json:"distinct"`
	ID        string   `json:"id"`
	Message   string   `json:"message"`
	Modified  []string `json:"modified"`
	Removed   []string `json:"removed"`
	Timestamp Time     `json:"timestamp"`
	URL       string   `json:"url"`
}

// HeadCommit was autogenerated by go generate. To see more details about this
// payload type visit https://developer.github.com/v3/activity/events/types.
type HeadCommit struct {
	Added     []string `json:"added"`
	Author    User     `json:"author"`
	Committer User     `json:"committer"`
	Distinct  bool     `json:"distinct"`
	ID        string   `json:"id"`
	Message   string   `json:"message"`
	Modified  []string `json:"modified"`
	Removed   []string `json:"removed"`
	Timestamp Time     `json:"timestamp"`
	URL       string   `json:"url"`
}

// PushEvent was autogenerated by go generate. To see more details about this
// payload type visit https://developer.github.com/v3/activity/events/types.
type PushEvent struct {
	After      string     `json:"after"`
	BaseRef    string     `json:"base_ref"`
	Before     string     `json:"before"`
	Commits    []Commits  `json:"commits"`
	Compare    string     `json:"compare"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	Forced     bool       `json:"forced"`
	HeadCommit HeadCommit `json:"head_commit"`
	Pusher     User       `json:"pusher"`
	Ref        string     `json:"ref"`
	Repository Repository `json:"repository"`
	Sender     User       `json:"sender"`
}

// Repository was autogenerated by go generate. To see more details about this
// payload type visit https://developer.github.com/v3/activity/events/types.
type Repository struct {
	ArchiveURL       string `json:"archive_url"`
	AssigneesURL     string `json:"assignees_url"`
	BlobsURL         string `json:"blobs_url"`
	BranchesURL      string `json:"branches_url"`
	CloneURL         string `json:"clone_url"`
	CollaboratorsURL string `json:"collaborators_url"`
	CommentsURL      string `json:"comments_url"`
	CommitsURL       string `json:"commits_url"`
	CompareURL       string `json:"compare_url"`
	ContentsURL      string `json:"contents_url"`
	ContributorsURL  string `json:"contributors_url"`
	CreatedAt        Time   `json:"created_at"`
	DefaultBranch    string `json:"default_branch"`
	Description      string `json:"description"`
	DownloadsURL     string `json:"downloads_url"`
	EventsURL        string `json:"events_url"`
	Fork             bool   `json:"fork"`
	Forks            int    `json:"forks"`
	ForksCount       int    `json:"forks_count"`
	ForksURL         string `json:"forks_url"`
	FullName         string `json:"full_name"`
	GitCommitsURL    string `json:"git_commits_url"`
	GitRefsURL       string `json:"git_refs_url"`
	GitTagsURL       string `json:"git_tags_url"`
	GitURL           string `json:"git_url"`
	HTMLURL          string `json:"html_url"`
	HasDownloads     bool   `json:"has_downloads"`
	HasIssues        bool   `json:"has_issues"`
	HasPages         bool   `json:"has_pages"`
	HasWiki          bool   `json:"has_wiki"`
	Homepage         string `json:"homepage"`
	HooksURL         string `json:"hooks_url"`
	ID               int    `json:"id"`
	IssueCommentURL  string `json:"issue_comment_url"`
	IssueEventsURL   string `json:"issue_events_url"`
	IssuesURL        string `json:"issues_url"`
	KeysURL          string `json:"keys_url"`
	LabelsURL        string `json:"labels_url"`
	Language         string `json:"language"`
	LanguagesURL     string `json:"languages_url"`
	MasterBranch     string `json:"master_branch"`
	MergesURL        string `json:"merges_url"`
	MilestonesURL    string `json:"milestones_url"`
	MirrorURL        string `json:"mirror_url"`
	Name             string `json:"name"`
	NotificationsURL string `json:"notifications_url"`
	OpenIssues       int    `json:"open_issues"`
	OpenIssuesCount  int    `json:"open_issues_count"`
	Owner            User   `json:"owner"`
	Private          bool   `json:"private"`
	PullsURL         string `json:"pulls_url"`
	PushedAt         Time   `json:"pushed_at"`
	ReleasesURL      string `json:"releases_url"`
	SSHURL           string `json:"ssh_url"`
	Size             int    `json:"size"`
	Stargazers       int    `json:"stargazers"`
	StargazersCount  int    `json:"stargazers_count"`
	StargazersURL    string `json:"stargazers_url"`
	StatusesURL      string `json:"statuses_url"`
	SubscribersURL   string `json:"subscribers_url"`
	SubscriptionURL  string `json:"subscription_url"`
	SvnURL           string `json:"svn_url"`
	TagsURL          string `json:"tags_url"`
	TeamsURL         string `json:"teams_url"`
	TreesURL         string `json:"trees_url"`
	URL              string `json:"url"`
	UpdatedAt        Time   `json:"updated_at"`
	Watchers         int    `json:"watchers"`
	WatchersCount    int    `json:"watchers_count"`
}

// Time embeds time.Time. The wrapper allows for unmarshalling time from JSON
// null value or unix timestamp.
type Time struct {
	time.Time
}

var null = []byte("null")

// MarshalJSON implements the json.Marshaler interface. The time is a quoted
// string in RFC 3339 format or "null" if it's a zero value.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return null, nil
	}
	return t.Time.MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaler interface. The time is expected
// to be a quoted string in RFC 3339 format, a unix timestamp or a "null" string.
func (t *Time) UnmarshalJSON(p []byte) (err error) {
	if bytes.Compare(p, null) == 0 {
		t.Time = time.Time{}
		return nil
	}
	if err = t.Time.UnmarshalJSON(p); err == nil {
		return nil
	}
	n, e := strconv.ParseInt(string(bytes.Trim(p, `"`)), 10, 64)
	if e != nil {
		return err
	}
	t.Time = time.Unix(n, 0)
	return nil
}