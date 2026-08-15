package workbench

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	ErrInvalidGoal       = errors.New("course goal is required")
	ErrProjectNotFound   = errors.New("course project not found")
	ErrInvalidOrder      = errors.New("chapter order must contain every chapter exactly once")
	ErrUnsupportedFormat = errors.New("export format must be pdf or pptx")
)

type Service struct {
	projects          []CourseProject
	exports           []ExportRecord
	nextProjectNumber int
	nextExportNumber  int
	sequence          int
}

func NewService() *Service {
	return &Service{
		projects:          fixedProjects(),
		exports:           fixedExports(),
		nextProjectNumber: 3,
		nextExportNumber:  2,
		sequence:          3,
	}
}

func (s *Service) CreateProject(goal string) (CourseProject, error) {
	goal = strings.TrimSpace(goal)
	if goal == "" {
		return CourseProject{}, ErrInvalidGoal
	}

	id := fmt.Sprintf("course-%03d", s.nextProjectNumber)
	s.nextProjectNumber++
	chapters := generatedChapters(id, goal)
	project := CourseProject{
		ID:                   id,
		Title:                "岗位实操能力提升课",
		Goal:                 goal,
		Institution:          "启程零售培训中心",
		Chapters:             chapters,
		TotalDurationMinutes: totalDuration(chapters),
		UpdatedSequence:      s.nextSequence(),
	}
	s.projects = append(s.projects, project)
	return cloneProject(project), nil
}

func (s *Service) Project(projectID string) (CourseProject, error) {
	project, err := s.findProject(projectID)
	if err != nil {
		return CourseProject{}, err
	}
	return cloneProject(*project), nil
}

func (s *Service) EditableDetail(projectID string) (EditableCourseDetail, error) {
	project, err := s.findProject(projectID)
	if err != nil {
		return EditableCourseDetail{}, err
	}
	return editableDetailFromProject(*project), nil
}

func (s *Service) ReorderChapters(projectID string, orderedChapterIDs []string) (CourseProject, error) {
	project, err := s.findProject(projectID)
	if err != nil {
		return CourseProject{}, err
	}
	if len(orderedChapterIDs) != len(project.Chapters) {
		return CourseProject{}, ErrInvalidOrder
	}

	byID := make(map[string]Chapter, len(project.Chapters))
	for _, chapter := range project.Chapters {
		byID[chapter.ID] = chapter
	}
	reordered := make([]Chapter, len(orderedChapterIDs))
	seen := make(map[string]bool, len(orderedChapterIDs))
	for index, chapterID := range orderedChapterIDs {
		chapter, ok := byID[chapterID]
		if !ok || seen[chapterID] {
			return CourseProject{}, ErrInvalidOrder
		}
		seen[chapterID] = true
		reordered[index] = chapter
	}

	project.Chapters = reordered
	project.TotalDurationMinutes = totalDuration(reordered)
	project.UpdatedSequence = s.nextSequence()
	return cloneProject(*project), nil
}

func (s *Service) CreateExport(projectID, format string) (ExportRecord, error) {
	project, err := s.findProject(projectID)
	if err != nil {
		return ExportRecord{}, err
	}
	format = strings.ToLower(strings.TrimSpace(format))
	if format != "pdf" && format != "pptx" {
		return ExportRecord{}, ErrUnsupportedFormat
	}

	record := ExportRecord{
		ID:        fmt.Sprintf("export-%03d", s.nextExportNumber),
		ProjectID: project.ID,
		FileName:  fmt.Sprintf("%s-v%d.%s", project.ID, project.UpdatedSequence, format),
		Format:    format,
		Status:    "已完成",
		Sequence:  s.nextSequence(),
	}
	s.nextExportNumber++
	s.exports = append(s.exports, record)
	return record, nil
}

func (s *Service) UserCenter() UserCenter {
	recent := make([]ProjectSummary, 0, len(s.projects))
	for _, project := range s.projects {
		recent = append(recent, ProjectSummary{
			ID:                   project.ID,
			Title:                project.Title,
			Goal:                 project.Goal,
			ChapterCount:         len(project.Chapters),
			TotalDurationMinutes: project.TotalDurationMinutes,
			UpdatedSequence:      project.UpdatedSequence,
		})
	}
	sort.SliceStable(recent, func(left, right int) bool {
		return recent[left].UpdatedSequence > recent[right].UpdatedSequence
	})

	exports := append([]ExportRecord(nil), s.exports...)
	sort.SliceStable(exports, func(left, right int) bool {
		return exports[left].Sequence > exports[right].Sequence
	})

	return UserCenter{
		TrainerName:   "林岚",
		Institution:   "启程零售培训中心",
		RecentCourses: recent,
		ExportRecords: exports,
	}
}

func (s *Service) findProject(projectID string) (*CourseProject, error) {
	for index := range s.projects {
		if s.projects[index].ID == projectID {
			return &s.projects[index], nil
		}
	}
	return nil, ErrProjectNotFound
}

func (s *Service) nextSequence() int {
	s.sequence++
	return s.sequence
}

func generatedChapters(projectID, goal string) []Chapter {
	return []Chapter{
		{
			ID:                      projectID + "-ch1",
			Title:                   "目标对齐与岗位情境",
			DurationMinutes:         8,
			InstructorScript:        []string{"本课程的岗位结果是：" + goal, "先用一个真实班次案例确认现状与目标之间的差距。"},
			DemoAssets:              []DemoAsset{{Name: "岗位任务开场案例", Kind: "情境图片", CuePoint: "标出任务角色、现场限制和验收标准"}},
			ExercisePrompt:          "用一句可观察、可验收的话复述课程目标。",
			ExpectedExerciseMinutes: 2,
		},
		{
			ID:                      projectID + "-ch2",
			Title:                   "标准操作拆解",
			DurationMinutes:         12,
			InstructorScript:        []string{"接下来按准备、执行、确认三个阶段拆解标准动作。", "每个动作都要对应现场证据，避免只记口号。"},
			DemoAssets:              []DemoAsset{{Name: "标准操作近景演示", Kind: "操作视频", CuePoint: "关键动作处暂停并叠加检查点"}},
			ExercisePrompt:          "把标准流程卡按正确顺序排列，并补充一个质量检查点。",
			ExpectedExerciseMinutes: 4,
		},
		{
			ID:                      projectID + "-ch3",
			Title:                   "异常判断与升级",
			DurationMinutes:         9,
			InstructorScript:        []string{"标准流程之外，岗位人员还要知道何时停止、隔离和上报。", "优先判断人员风险，再判断产品或设备影响范围。"},
			DemoAssets:              []DemoAsset{{Name: "异常情境分支演练", Kind: "交互演示", CuePoint: "依次选择停止、隔离、记录和升级"}},
			ExercisePrompt:          "面对信息不完整的异常案例，写出第一步动作和升级依据。",
			ExpectedExerciseMinutes: 3,
		},
		{
			ID:                      projectID + "-ch4",
			Title:                   "独立练习与现场验收",
			DurationMinutes:         11,
			InstructorScript:        []string{"最后一次练习由学员独立完成，讲师只按验收表观察和记录。", "未通过项要落到具体动作和再次验证时间。"},
			DemoAssets:              []DemoAsset{{Name: "岗位验收表示范", Kind: "表单演示", CuePoint: "填写观察证据、结果和复训动作"}},
			ExercisePrompt:          "完成一次限时岗位演练，并依据验收表进行同伴复核。",
			ExpectedExerciseMinutes: 5,
		},
	}
}

func editableDetailFromProject(project CourseProject) EditableCourseDetail {
	return EditableCourseDetail{
		ProjectID:            project.ID,
		Goal:                 project.Goal,
		Chapters:             chapterListFromModel(project.Chapters),
		TotalDurationMinutes: project.TotalDurationMinutes,
	}
}

func chapterListFromModel(chapters []Chapter) []Chapter {
	return chapters
}

func totalDuration(chapters []Chapter) int {
	total := 0
	for _, chapter := range chapters {
		total += chapter.DurationMinutes
	}
	return total
}

func cloneProject(project CourseProject) CourseProject {
	chapters := project.Chapters
	project.Chapters = make([]Chapter, len(chapters))
	for index, chapter := range chapters {
		project.Chapters[index] = cloneChapter(chapter)
	}
	return project
}

func cloneChapter(chapter Chapter) Chapter {
	chapter.InstructorScript = append([]string(nil), chapter.InstructorScript...)
	chapter.DemoAssets = append([]DemoAsset(nil), chapter.DemoAssets...)
	return chapter
}
