package swagger

import (
	"time"
)

func CreateUser() User {
	user := User{Name: "Test user name", Surname: "Barnard"}

	return user
}

func CreateUsers() []User {
	user1 := CreateUser()
	user2 := CreateUser()

	users := Users{}
	users.Users = []User{user1, user2}

	return []User{user1, user2}
}

func CreateMedium() MediumSized {
	medium := MediumSized{
		Name:      "some name",
		Surname:   "some surname",
		Address:   "some old cape road, Grabouw",
		Location:  "Western cape?",
		Country:   "South Africa",
		Age:       13,
		BirthDate: time.Now(),
		StartDate: time.Now().AddDate(-10, 0, 0),
		Type_:     "Male",
	}

	return medium
}

func CreateLargeSize() LargeSized {
	user := CreateUser()
	users := CreateUsers()
	medium := CreateMedium()

	large := LargeSized{
		One:           1.1,
		Two:           6,
		Three:         true,
		Four:          &user,
		Five:          "some string",
		Six:           users,
		Seven:         &medium,
		BiggerThan100: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Nine:          "asdfgrqeagsdfgsdfg",
		Ten:           "asodijfnaslkjdfnaklsdjfnklqajwenrflksajdnfkljasndfkljasndf",
		Eleven:        "some more text sthis is getting old",
		Twelve:        "asdasdasdasdasdasd",
		Thirteen:      "ccccccccccccccccccc",
		Fourteen:      "asdasdasvregsefdgsdfg",
		Fifteen:       "asdfsadf",
		Sixteen:       "even more text!",
		Seventeen:     "and more",
		Eighteen:      "and more",
		Nineteen:      "extra",
		Twenty:        "more more more",
		TwentyOne:     "This is getting old",
		TwentyTwo:     "really old",
		TwentyThree:   " and more more more",
		TwentyFour:    " aamper klaar",
		TwentyFive:    "klaar",
	}

	return large
}

func CreateLargeSizeGroup() LargeSizeGroup {
	one := CreateLargeSize()
	Two := CreateLargeSize()
	three := CreateLargeSize()
	q1 := CreateLargeSize()
	q2 := CreateLargeSize()
	q3 := CreateLargeSize()
	q4 := CreateLargeSize()
	q5 := CreateLargeSize()
	q6 := CreateLargeSize()
	q7 := CreateLargeSize()
	q8 := CreateLargeSize()
	q9 := CreateLargeSize()
	q10 := CreateLargeSize()
	q11 := CreateLargeSize()
	q12 := CreateLargeSize()

	large := LargeSizeGroup{Reason: "there isn't actually a reason for this",
		ManyLargeSizeItems: []LargeSized{one, Two, three, q1, q2, q3, q4, q5, q6, q7, q8, q9, q10, q11, q12}}

	return large
}
