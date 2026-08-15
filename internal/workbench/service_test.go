package workbench_test

import (
	"errors"
	"testing"

	"training-storyboard.local/workbench/internal/workbench"
)

func TestCreateProjectBuildsCompleteStoryboard(t *testing.T) {
	service := workbench.NewService()
	goal := "区域督导能够辅导店长完成一次客诉复盘，并形成可追踪的改善任务"

	project, err := service.CreateProject(goal)
	if err != nil {
		t.Fatalf("CreateProject() error = %v", err)
	}
	if project.ID != "course-003" {
		t.Fatalf("project ID = %q, want course-003", project.ID)
	}
	if project.Goal != goal {
		t.Fatalf("project goal = %q, want %q", project.Goal, goal)
	}
	if len(project.Chapters) != 4 {
		t.Fatalf("chapter count = %d, want 4", len(project.Chapters))
	}

	wantTotal := 0
	for _, chapter := range project.Chapters {
		wantTotal += chapter.DurationMinutes
		if len(chapter.InstructorScript) == 0 {
			t.Fatalf("chapter %q has no instructor script", chapter.ID)
		}
		if len(chapter.DemoAssets) == 0 {
			t.Fatalf("chapter %q has no demo asset", chapter.ID)
		}
		if chapter.ExercisePrompt == "" {
			t.Fatalf("chapter %q has no exercise prompt", chapter.ID)
		}
	}
	if project.TotalDurationMinutes != wantTotal {
		t.Fatalf("total duration = %d, want %d", project.TotalDurationMinutes, wantTotal)
	}
}

func TestReorderChaptersSynchronizesProject(t *testing.T) {
	service := workbench.NewService()
	detail, err := service.EditableDetail("course-001")
	if err != nil {
		t.Fatalf("EditableDetail() error = %v", err)
	}
	order := make([]string, len(detail.Chapters))
	for index := range detail.Chapters {
		order[len(order)-1-index] = detail.Chapters[index].ID
	}

	project, err := service.ReorderChapters("course-001", order)
	if err != nil {
		t.Fatalf("ReorderChapters() error = %v", err)
	}
	if project.Chapters[0].ID != "course-001-ch4" {
		t.Fatalf("first chapter = %q, want course-001-ch4", project.Chapters[0].ID)
	}
	if project.TotalDurationMinutes != 45 {
		t.Fatalf("total duration = %d, want 45", project.TotalDurationMinutes)
	}

	center := service.UserCenter()
	if center.RecentCourses[0].ID != "course-001" {
		t.Fatalf("latest course = %q, want course-001", center.RecentCourses[0].ID)
	}
	if center.RecentCourses[0].TotalDurationMinutes != 45 {
		t.Fatalf("center duration = %d, want 45", center.RecentCourses[0].TotalDurationMinutes)
	}
}

func TestUserCenterTracksExports(t *testing.T) {
	service := workbench.NewService()
	record, err := service.CreateExport("course-002", "PDF")
	if err != nil {
		t.Fatalf("CreateExport() error = %v", err)
	}
	if record.ID != "export-002" || record.Format != "pdf" || record.Status != "已完成" {
		t.Fatalf("export record = %#v", record)
	}

	center := service.UserCenter()
	if center.TrainerName != "林岚" || center.Institution != "启程零售培训中心" {
		t.Fatalf("trainer center = %#v", center)
	}
	if center.ExportRecords[0] != record {
		t.Fatalf("latest export = %#v, want %#v", center.ExportRecords[0], record)
	}
	if center.RecentCourses[0].ID != "course-002" {
		t.Fatalf("latest course = %q, want course-002", center.RecentCourses[0].ID)
	}
}

func TestWorkflowValidation(t *testing.T) {
	service := workbench.NewService()
	if _, err := service.CreateProject("  "); !errors.Is(err, workbench.ErrInvalidGoal) {
		t.Fatalf("CreateProject() error = %v, want %v", err, workbench.ErrInvalidGoal)
	}
	if _, err := service.ReorderChapters("course-001", []string{"course-001-ch1"}); !errors.Is(err, workbench.ErrInvalidOrder) {
		t.Fatalf("ReorderChapters() error = %v, want %v", err, workbench.ErrInvalidOrder)
	}
	if _, err := service.CreateExport("course-001", "docx"); !errors.Is(err, workbench.ErrUnsupportedFormat) {
		t.Fatalf("CreateExport() error = %v, want %v", err, workbench.ErrUnsupportedFormat)
	}
	if _, err := service.Project("course-999"); !errors.Is(err, workbench.ErrProjectNotFound) {
		t.Fatalf("Project() error = %v, want %v", err, workbench.ErrProjectNotFound)
	}
}

func TestEditableDetailRoundTrip(t *testing.T) {
	service := workbench.NewService()
	first, err := service.EditableDetail("course-001")
	if err != nil {
		t.Fatalf("first EditableDetail() error = %v", err)
	}
	wantTitle := first.Chapters[0].Title
	first.Chapters[0].Title = "临时调整：先做门店照片复盘"

	second, err := service.EditableDetail("course-001")
	if err != nil {
		t.Fatalf("second EditableDetail() error = %v", err)
	}
	if second.Chapters[0].Title != wantTitle {
		t.Fatalf("chapter title = %q, want %q", second.Chapters[0].Title, wantTitle)
	}
}
