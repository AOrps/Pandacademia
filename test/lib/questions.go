package lib

// ===============================================
// Types
// ===============================================

type Question struct {
	QuestionText string
	NameElement  string
}

type QuestionPage struct {
	Ask       bool
	Questions []Question
	Code      int8
}

// ===============================================
// Functions
// ===============================================

// MakeQuesitonArr() -> makes an array of questions
func MakeQuestionArr() *[]Question {
	return &[]Question{
		{QuestionText: "Are you feeling sick?", NameElement: "sick"},
		{QuestionText: "Do you have muscle soreness or respiratory trouble that you can’t attribute to another medical condition?", NameElement: "sore"},
		{QuestionText: "Have you recently had close contact (within 6 feet of an infected person for at least 15 minutes) with someone with symptoms of COVID-19, tested for COVID-19, or diagnosed with COVID-19?", NameElement: "contact"},
		{QuestionText: "Have you recently been in a nursing home, healthcare facility, or homeless shelter?", NameElement: "location"},
	}
}
