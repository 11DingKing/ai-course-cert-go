package domain

import (
	"testing"
	"time"
)

func TestDomainInvariant130(t *testing.T) {
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
func TestDomainInvariant131(t *testing.T) {
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
func TestDomainInvariant132(t *testing.T) {
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
func TestDomainInvariant133(t *testing.T) {
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
func TestDomainInvariant134(t *testing.T) {
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
func TestDomainInvariant135(t *testing.T) {
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
func TestDomainInvariant136(t *testing.T) {
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
func TestDomainInvariant137(t *testing.T) {
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
func TestDomainInvariant138(t *testing.T) {
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
func TestDomainInvariant139(t *testing.T) {
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
