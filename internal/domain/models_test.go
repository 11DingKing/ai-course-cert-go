package domain

import "testing"

func TestSubmissionTransitions(t *testing.T) {
	cases := []struct {
		from, to SubmissionStatus
		ok       bool
	}{{StatusDraft, StatusSubmitted, true}, {StatusSubmitted, StatusReturned, true}, {StatusSubmitted, StatusApproved, true}, {StatusReturned, StatusSubmitted, true}, {StatusApproved, StatusArchived, true}, {StatusDraft, StatusApproved, false}, {StatusArchived, StatusDraft, false}}
	for _, c := range cases {
		if c.from.CanTransition(c.to) != c.ok {
			t.Fatalf("%s -> %s", c.from, c.to)
		}
	}
}
func TestRoleValues(t *testing.T) {
	for _, r := range []Role{RoleStudent, RoleTeacher, RoleReviewer, RoleAdmin} {
		if r == "" {
			t.Fatal("empty")
		}
	}
}
