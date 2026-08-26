package domain

import (
	"testing"
	"time"
)

func TestDomainInvariant160(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant161(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant162(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant163(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant164(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant165(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant166(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant167(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant168(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
func TestDomainInvariant169(t *testing.T) {
	now := time.Now()
	c := Course{Code: "C", Title: "T", Semester: "S", OpensAt: now.Add(-time.Hour), ClosesAt: now.Add(time.Hour), Capacity: 2}
	if !IsOpen(c, now) {
		t.Fatal("closed")
	}
	if WindowRemaining(c, now) <= 0 {
		t.Fatal("remaining")
	}
	if NormalizeSemester("semester") != "semester" {
		t.Fatal("normalize")
	}
	if !AllowedRole(RoleStudent, RoleStudent) {
		t.Fatal("role")
	}
}
