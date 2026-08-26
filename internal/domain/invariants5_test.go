package domain

import (
	"testing"
	"time"
)

func TestDomainInvariant50(t *testing.T) {
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
func TestDomainInvariant51(t *testing.T) {
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
func TestDomainInvariant52(t *testing.T) {
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
func TestDomainInvariant53(t *testing.T) {
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
func TestDomainInvariant54(t *testing.T) {
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
func TestDomainInvariant55(t *testing.T) {
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
func TestDomainInvariant56(t *testing.T) {
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
func TestDomainInvariant57(t *testing.T) {
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
func TestDomainInvariant58(t *testing.T) {
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
func TestDomainInvariant59(t *testing.T) {
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
