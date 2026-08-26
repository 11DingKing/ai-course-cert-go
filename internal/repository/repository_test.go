package repository

import (
	"context"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/storage"
	"path/filepath"
	"testing"
	"time"
)

func fixture(t *testing.T) (context.Context, *Users, *Courses, *Submissions, *Evidences, *Reviews) {
	t.Helper()
	db, e := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	if e = storage.Migrate(context.Background(), db); e != nil {
		t.Fatal(e)
	}
	return context.Background(), &Users{db}, &Courses{db}, &Submissions{db}, &Evidences{db}, &Reviews{db}
}
func TestUsersCreateLookup(t *testing.T) {
	ctx, u, _, _, _, _ := fixture(t)
	x, e := u.Create(ctx, "s@example.edu", "Student", domain.RoleStudent, "pw")
	if e != nil || x.ID == 0 {
		t.Fatal(e)
	}
	y, e := u.ByEmail(ctx, x.Email)
	if e != nil || y.Role != domain.RoleStudent {
		t.Fatalf("%+v %v", y, e)
	}
	if _, e = u.Create(ctx, x.Email, "Other", domain.RoleStudent, "pw"); e == nil {
		t.Fatal("duplicate accepted")
	}
}
func TestCourseCRUDAndPagination(t *testing.T) {
	ctx, _, c, _, _, _ := fixture(t)
	now := time.Now().Add(-time.Hour)
	for i := 0; i < 5; i++ {
		if _, e := c.Create(ctx, domain.Course{Code: string(rune('A' + i)), Title: "Language", Semester: "2026", OpensAt: now, ClosesAt: now.Add(24 * time.Hour), Capacity: 2}); e != nil {
			t.Fatal(e)
		}
	}
	xs, e := c.List(ctx, 2, 1)
	if e != nil || len(xs) != 2 {
		t.Fatalf("%d %v", len(xs), e)
	}
	if _, e = c.Get(ctx, 999); e == nil {
		t.Fatal("missing course")
	}
}
func TestSubmissionCapacityAndWindow(t *testing.T) {
	ctx, u, c, s, _, _ := fixture(t)
	stu, _ := u.Create(ctx, "a@x", "A", domain.RoleStudent, "p")
	now := time.Now().Add(-time.Minute)
	course, _ := c.Create(ctx, domain.Course{Code: "C", Title: "Course", Semester: "S", OpensAt: now, ClosesAt: now.Add(time.Hour), Capacity: 1})
	x, e := s.Create(ctx, domain.Submission{CourseID: course.ID, StudentID: stu.ID, Title: "essay", AIStatement: "used"})
	if e != nil || x.ID == 0 {
		t.Fatal(e)
	}
	if _, e = s.Create(ctx, domain.Submission{CourseID: course.ID, StudentID: stu.ID, Title: "essay2", AIStatement: "used"}); e == nil {
		t.Fatal("capacity ignored")
	}
}
func TestSubmissionTransitionVersion(t *testing.T) {
	ctx, u, c, s, _, _ := fixture(t)
	stu, _ := u.Create(ctx, "b@x", "B", domain.RoleStudent, "p")
	now := time.Now().Add(-time.Minute)
	course, _ := c.Create(ctx, domain.Course{Code: "D", Title: "Course", Semester: "S", OpensAt: now, ClosesAt: now.Add(time.Hour), Capacity: 2})
	x, _ := s.Create(ctx, domain.Submission{CourseID: course.ID, StudentID: stu.ID, Title: "t", AIStatement: "a"})
	y, e := s.Transition(ctx, x.ID, domain.StatusDraft, domain.StatusSubmitted, 1)
	if e != nil || y.Status != domain.StatusSubmitted {
		t.Fatal(e)
	}
	if _, e = s.Transition(ctx, x.ID, domain.StatusSubmitted, domain.StatusApproved, 1); e == nil {
		t.Fatal("stale version accepted")
	}
	if _, e = s.Transition(ctx, x.ID, domain.StatusDraft, domain.StatusApproved, 2); e == nil {
		t.Fatal("invalid transition accepted")
	}
}
func TestEvidenceVersions(t *testing.T) {
	ctx, u, c, s, evi, _ := fixture(t)
	stu, _ := u.Create(ctx, "c@x", "C", domain.RoleStudent, "p")
	now := time.Now().Add(-time.Minute)
	course, _ := c.Create(ctx, domain.Course{Code: "E", Title: "Course", Semester: "S", OpensAt: now, ClosesAt: now.Add(time.Hour), Capacity: 2})
	x, _ := s.Create(ctx, domain.Submission{CourseID: course.ID, StudentID: stu.ID, Title: "t", AIStatement: "a"})
	for i := 0; i < 3; i++ {
		if _, e := evi.Add(ctx, domain.Evidence{SubmissionID: x.ID, URI: "u", Checksum: string(rune('0' + i)), Description: "d"}); e != nil {
			t.Fatal(e)
		}
	}
	xs, e := evi.List(ctx, x.ID)
	if e != nil || len(xs) != 3 || xs[2].Version != 3 {
		t.Fatalf("%+v", xs)
	}
}
func TestReviewsList(t *testing.T) {
	ctx, u, c, s, _, rv := fixture(t)
	stu, _ := u.Create(ctx, "d@x", "D", domain.RoleStudent, "p")
	rev, _ := u.Create(ctx, "r@x", "R", domain.RoleReviewer, "p")
	now := time.Now().Add(-time.Minute)
	course, _ := c.Create(ctx, domain.Course{Code: "F", Title: "Course", Semester: "S", OpensAt: now, ClosesAt: now.Add(time.Hour), Capacity: 2})
	x, _ := s.Create(ctx, domain.Submission{CourseID: course.ID, StudentID: stu.ID, Title: "t", AIStatement: "a"})
	if _, e := rv.Add(ctx, domain.Review{SubmissionID: x.ID, ReviewerID: rev.ID, Decision: "return", Comment: "fix", Version: 1}); e != nil {
		t.Fatal(e)
	}
	xs, e := rv.List(ctx, x.ID)
	if e != nil || len(xs) != 1 {
		t.Fatal(e)
	}
}
