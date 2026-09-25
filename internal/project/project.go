package project

type Subproject string

const (
	SubprojectMicrofiction    Subproject = "MICROFICTION"
	SubprojectSystemOverview  Subproject = "SYSTEM_OVERVIEW"
	SubprojectPodcastGame     Subproject = "PODCAST_GAME_SESSION"
	SubprojectPodcastRolafter Subproject = "PODCAST_ROLAFTER"
	SubprojectClips           Subproject = "CLIPS"
)

type CreateRequest struct {
	Name         string
	Subprojects  []Subproject
	InitialClips int
}
