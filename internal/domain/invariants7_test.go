package domain

import (
	"testing"
	"time"
)

func TestDomainInvariant70(t *testing.T) {
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
func TestDomainInvariant71(t *testing.T) {
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
func TestDomainInvariant72(t *testing.T) {
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
func TestDomainInvariant73(t *testing.T) {
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
func TestDomainInvariant74(t *testing.T) {
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
func TestDomainInvariant75(t *testing.T) {
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
func TestDomainInvariant76(t *testing.T) {
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
func TestDomainInvariant77(t *testing.T) {
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
func TestDomainInvariant78(t *testing.T) {
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
func TestDomainInvariant79(t *testing.T) {
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
