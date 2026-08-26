package repository

import ("errors"; "fmt")

var ErrReviewCapacity = errors.New("review capacity exhausted")

func CapacityFailure(courseID int64) error {
 return fmt.Errorf("course %d cannot assign reviewer: %v", courseID, ErrReviewCapacity)
}
