package workbench

type DemoAsset struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	CuePoint string `json:"cue_point"`
}

type Chapter struct {
	ID                      string      `json:"id"`
	Title                   string      `json:"title"`
	DurationMinutes         int         `json:"duration_minutes"`
	InstructorScript        []string    `json:"instructor_script"`
	DemoAssets              []DemoAsset `json:"demo_assets"`
	ExercisePrompt          string      `json:"exercise_prompt"`
	ExpectedExerciseMinutes int         `json:"expected_exercise_minutes"`
}

type CourseProject struct {
	ID                   string    `json:"id"`
	Title                string    `json:"title"`
	Goal                 string    `json:"goal"`
	Institution          string    `json:"institution"`
	Chapters             []Chapter `json:"chapters"`
	TotalDurationMinutes int       `json:"total_duration_minutes"`
	UpdatedSequence      int       `json:"updated_sequence"`
}

type EditableCourseDetail struct {
	ProjectID            string    `json:"project_id"`
	Goal                 string    `json:"goal"`
	Chapters             []Chapter `json:"chapters"`
	TotalDurationMinutes int       `json:"total_duration_minutes"`
}

type ProjectSummary struct {
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Goal                 string `json:"goal"`
	ChapterCount         int    `json:"chapter_count"`
	TotalDurationMinutes int    `json:"total_duration_minutes"`
	UpdatedSequence      int    `json:"updated_sequence"`
}

type ExportRecord struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	FileName  string `json:"file_name"`
	Format    string `json:"format"`
	Status    string `json:"status"`
	Sequence  int    `json:"sequence"`
}

type UserCenter struct {
	TrainerName   string           `json:"trainer_name"`
	Institution   string           `json:"institution"`
	RecentCourses []ProjectSummary `json:"recent_courses"`
	ExportRecords []ExportRecord   `json:"export_records"`
}
