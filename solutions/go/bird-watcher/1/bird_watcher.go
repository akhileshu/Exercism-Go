package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, birds := range birdsPerDay {
		count += birds
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
/*
birdsPerDay := []int{
	2, 5, 0, 7, 4, 1, 3, // week 1
	0, 2, 5, 0, 1, 3, 1, // week 2
}
*/
func BirdsInWeek(birdsPerDay []int, week int) int {
	count := 0
	birdsInAWeek := birdsPerDay[(week-1)*7 : (week-1)*7+7]
	for _, birds := range birdsInAWeek {
		count += birds
	}
	return count

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, _ := range birdsPerDay {
		isOddDay := index%2 == 0
		if isOddDay {
			birdsPerDay[index]++
		}
	}
	return birdsPerDay

}
