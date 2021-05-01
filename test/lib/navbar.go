package lib

// ===============================================
// Types
// ===============================================

type Link struct {
	Linkname string
	Href     string
}

// ===============================================
// Functions
// ===============================================

// GetNavBar -> Gets the elements for the navbar
func GetNavBar() []Link {
	yes := []Link{
		{Linkname: "Questions", Href: "/questions"},
		{Linkname: "Data Analytics", Href: "/analysis"},
		{Linkname: "Data Visulization", Href: "/viz"},
		{Linkname: "Info", Href: "/info"},
	}
	return yes
}
