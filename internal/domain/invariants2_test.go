package domain

import (
	"testing"
	"time"
)

func TestDomainInvariant20(t *testing.T) {
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
func TestDomainInvariant21(t *testing.T) {
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
func TestDomainInvariant22(t *testing.T) {
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
func TestDomainInvariant23(t *testing.T) {
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
func TestDomainInvariant24(t *testing.T) {
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
func TestDomainInvariant25(t *testing.T) {
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
func TestDomainInvariant26(t *testing.T) {
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
func TestDomainInvariant27(t *testing.T) {
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
func TestDomainInvariant28(t *testing.T) {
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
func TestDomainInvariant29(t *testing.T) {
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
