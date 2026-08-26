package service

import (
	"context"
	"github.com/11DingKing/ai-course-cert-go/internal/audit"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"github.com/11DingKing/ai-course-cert-go/internal/storage"
	"path/filepath"
	"testing"
	"time"
)

func svcFixture(t *testing.T) (Service, domain.User, domain.User, domain.Course) {
	db, _ := storage.Open(filepath.Join(t.TempDir(), "s.db"))
	storage.Migrate(context.Background(), db)
	u := repository.Users{DB: db}
	stu, _ := u.Create(context.Background(), "student@x", "Student", domain.RoleStudent, "pw")
	rev, _ := u.Create(context.Background(), "reviewer@x", "Reviewer", domain.RoleReviewer, "pw")
	c, _ := repository.Courses{DB: db}.Create(context.Background(), domain.Course{Code: "SVC", Title: "AI", Semester: "2026", OpensAt: time.Now().Add(-time.Minute), ClosesAt: time.Now().Add(time.Hour), Capacity: 3})
	return Service{Users: u, Courses: repository.Courses{DB: db}, Submissions: repository.Submissions{DB: db}, Evidences: repository.Evidences{DB: db}, Reviews: repository.Reviews{DB: db}, Audit: audit.Logger{DB: db}}, stu, rev, c
}
func TestLogin(t *testing.T) {
	s, stu, _, _ := svcFixture(t)
	if _, e := s.Login(context.Background(), stu.Email, "bad"); e == nil {
		t.Fatal("bad password accepted")
	}
	if u, e := s.Login(context.Background(), stu.Email, "pw"); e != nil || u.ID != stu.ID {
		t.Fatal(e)
	}
}
func TestSubmitOwnershipAndAudit(t *testing.T) {
	s, stu, rev, c := svcFixture(t)
	x, _ := s.Submissions.Create(context.Background(), domain.Submission{CourseID: c.ID, StudentID: stu.ID, Title: "essay", AIStatement: "declared"})
	if _, e := s.Submit(context.Background(), rev.ID, x.ID, 1, "req"); e == nil {
		t.Fatal("reviewer submitted")
	}
	y, e := s.Submit(context.Background(), stu.ID, x.ID, 1, "req")
	if e != nil || y.Status != domain.StatusSubmitted {
		t.Fatal(e)
	}
	var n int
	s.Audit.DB.QueryRow(`SELECT COUNT(*) FROM audit_events`).Scan(&n)
	if n != 1 {
		t.Fatalf("audit=%d", n)
	}
}
func TestReviewWorkflow(t *testing.T) {
	s, stu, rev, c := svcFixture(t)
	x, _ := s.Submissions.Create(context.Background(), domain.Submission{CourseID: c.ID, StudentID: stu.ID, Title: "oral", AIStatement: "none"})
	s.Submit(context.Background(), stu.ID, x.ID, 1, "r")
	r, e := s.Review(context.Background(), rev.ID, x.ID, 2, "return", "add transcript", "q")
	if e != nil || r.Decision != "return" {
		t.Fatal(e)
	}
	z, _ := s.Submissions.Get(context.Background(), x.ID)
	if z.Status != domain.StatusReturned {
		t.Fatal(z.Status)
	}
}
func TestEvidenceValidation(t *testing.T) {
	s, stu, _, c := svcFixture(t)
	x, _ := s.Submissions.Create(context.Background(), domain.Submission{CourseID: c.ID, StudentID: stu.ID, Title: "read", AIStatement: "yes"})
	if _, e := s.AddEvidence(context.Background(), stu.ID, x.ID, "", "sum", "d"); e == nil {
		t.Fatal("missing metadata accepted")
	}
	if _, e := s.AddEvidence(context.Background(), stu.ID, x.ID, "uri", "sum", "d"); e != nil {
		t.Fatal(e)
	}
}
