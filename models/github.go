package models

import "time"

type GitHubPushPayload struct {
	Ref        string     `json:"ref"`
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Repository Repository `json:"repository"`
	Pusher     User       `json:"pusher"`
	Sender     User       `json:"sender"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	Forced     bool       `json:"forced"`
	BaseRef    *string    `json:"base_ref"`
	Compare    string     `json:"compare"`
	Commits    []Commit   `json:"commits"`
	HeadCommit Commit     `json:"head_commit"`
}

type Repos []Repository
type Repository struct {
	ID                       int64                `json:"id"`
	NodeID                   string               `json:"node_id"`
	Name                     string               `json:"name"`
	FullName                 string               `json:"full_name"`
	Private                  bool                 `json:"private"`
	Owner                    User                 `json:"owner"`
	HTMLURL                  string               `json:"html_url"`
	Description              *string              `json:"description,omitempty"`
	Fork                     bool                 `json:"fork"`
	URL                      string               `json:"url"`
	ArchiveURL               string               `json:"archive_url"`
	AssigneesURL             string               `json:"assignees_url"`
	BlobsURL                 string               `json:"blobs_url"`
	BranchesURL              string               `json:"branches_url"`
	CollaboratorsURL         string               `json:"collaborators_url"`
	CommentsURL              string               `json:"comments_url"`
	CommitsURL               string               `json:"commits_url"`
	CompareURL               string               `json:"compare_url"`
	ContentsURL              string               `json:"contents_url"`
	ContributorsURL          string               `json:"contributors_url"`
	DeploymentsURL           string               `json:"deployments_url"`
	DownloadsURL             string               `json:"downloads_url"`
	EventsURL                string               `json:"events_url"`
	ForksURL                 string               `json:"forks_url"`
	GitCommitsURL            string               `json:"git_commits_url"`
	GitRefsURL               string               `json:"git_refs_url"`
	GitTagsURL               string               `json:"git_tags_url"`
	GitURL                   string               `json:"git_url"`
	SSHURL                   string               `json:"ssh_url"`
	CloneURL                 string               `json:"clone_url"`
	SVNURL                   string               `json:"svn_url"`
	Homepage                 *string              `json:"homepage"`
	Size                     int                  `json:"size"`
	StargazersCount          int                  `json:"stargazers_count"`
	WatchersCount            int                  `json:"watchers_count"`
	Language                 *string              `json:"language"`
	HasIssues                bool                 `json:"has_issues"`
	HasProjects              bool                 `json:"has_projects"`
	HasDownloads             bool                 `json:"has_downloads"`
	HasWiki                  bool                 `json:"has_wiki"`
	HasPages                 bool                 `json:"has_pages"`
	HasDiscussions           bool                 `json:"has_discussions"`
	ForksCount               int                  `json:"forks_count"`
	MirrorURL                *string              `json:"mirror_url"`
	Archived                 bool                 `json:"archived"`
	Disabled                 bool                 `json:"disabled"`
	OpenIssuesCount          int                  `json:"open_issues_count"`
	License                  *License             `json:"license"` // ✅ fixed
	AllowForking             bool                 `json:"allow_forking"`
	IsTemplate               bool                 `json:"is_template"`
	WebCommitSignoffRequired bool                 `json:"web_commit_signoff_required"`
	Topics                   []string             `json:"topics"`
	Visibility               string               `json:"visibility"`
	Forks                    int                  `json:"forks"`
	OpenIssues               int                  `json:"open_issues"`
	Watchers                 int                  `json:"watchers"`
	DefaultBranch            string               `json:"default_branch"`
	MasterBranch             *string              `json:"master_branch"`                   // ✅ optional
	Permissions              *Permissions         `json:"permissions,omitempty"`           // ✅ optional
	SecurityAndAnalysis      *SecurityAndAnalysis `json:"security_and_analysis,omitempty"` // ✅ optional
	Organization             *User                `json:"organization,omitempty"`          // ✅ optional
	Parent                   *Repository          `json:"parent,omitempty"`                // ✅ optional (for forks)
	Source                   *Repository          `json:"source,omitempty"`                // ✅ optional (for forks)
	NetworkCount             *int                 `json:"network_count,omitempty"`
	SubscribersCount         *int                 `json:"subscribers_count,omitempty"`
	CreatedAt                time.Time            `json:"created_at"` // ✅ string (ISO 8601)
	UpdatedAt                time.Time            `json:"updated_at"`
	PushedAt                 time.Time            `json:"pushed_at"`
}
type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Commit struct {
	ID        string   `json:"id"`
	TreeID    string   `json:"tree_id"`
	Distinct  bool     `json:"distinct"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
	Author    GitUser  `json:"author"`
	Committer GitUser  `json:"committer"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

type GitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type User struct {
	Name              *string `json:"name,omitempty"`
	Email             *string `json:"email,omitempty"`
	Login             string  `json:"login"` // required in schema
	ID                *int64  `json:"id,omitempty"`
	NodeID            *string `json:"node_id,omitempty"`
	AvatarURL         string  `json:"avatar_url"` // required
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               string  `json:"url"`                 // required
	HTMLURL           string  `json:"html_url"`            // required
	FollowersURL      string  `json:"followers_url"`       // required
	FollowingURL      string  `json:"following_url"`       // required
	GistsURL          string  `json:"gists_url"`           // required
	StarredURL        string  `json:"starred_url"`         // required
	SubscriptionsURL  string  `json:"subscriptions_url"`   // required
	OrganizationsURL  string  `json:"organizations_url"`   // required
	ReposURL          string  `json:"repos_url"`           // required
	EventsURL         string  `json:"events_url"`          // required
	ReceivedEventsURL string  `json:"received_events_url"` // required
	Type              string  `json:"type"`                // required
	SiteAdmin         bool    `json:"site_admin"`          // required
	StarredAt         *string `json:"starred_at,omitempty"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type Permissions struct {
	Admin    *bool `json:"admin,omitempty"`
	Maintain *bool `json:"maintain,omitempty"`
	Push     *bool `json:"push,omitempty"`
	Triage   *bool `json:"triage,omitempty"`
	Pull     *bool `json:"pull,omitempty"`
}

type CodeOfConduct struct {
	Key     *string `json:"key,omitempty"`
	Name    *string `json:"name,omitempty"`
	URL     *string `json:"url,omitempty"`
	Body    *string `json:"body,omitempty"`
	HTMLURL *string `json:"html_url,omitempty"`
}

type License struct {
	Key    *string `json:"key,omitempty"`
	Name   *string `json:"name,omitempty"`
	SpdxID *string `json:"spdx_id,omitempty"`
	URL    *string `json:"url,omitempty"`
	NodeID *string `json:"node_id,omitempty"`
}

type AnalysisStatus struct {
	Status *string `json:"status,omitempty"` // "enabled" | "disabled"
}

type SecurityAndAnalysis struct {
	AdvancedSecurity                  *AnalysisStatus `json:"advanced_security,omitempty"`
	CodeSecurity                      *AnalysisStatus `json:"code_security,omitempty"`
	DependabotSecurityUpdates         *AnalysisStatus `json:"dependabot_security_updates,omitempty"`
	SecretScanning                    *AnalysisStatus `json:"secret_scanning,omitempty"`
	SecretScanningPushProtection      *AnalysisStatus `json:"secret_scanning_push_protection,omitempty"`
	SecretScanningNonProviderPatterns *AnalysisStatus `json:"secret_scanning_non_provider_patterns,omitempty"`
	SecretScanningAIDetection         *AnalysisStatus `json:"secret_scanning_ai_detection,omitempty"`
}
