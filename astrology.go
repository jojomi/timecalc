package timecalc

import "time"

type AstrologicalSign struct {
	Name          string
	LocalizedName string
	Symbol        string
}

var astrologicalSigns = map[string]AstrologicalSign{
	"Aries": {
		Name:          "Aries",
		LocalizedName: "Widder",
		Symbol:        "♈",
	},
	"Taurus": {
		Name:          "Taurus",
		LocalizedName: "Stier",
		Symbol:        "♉",
	},
	"Gemini": {
		Name:          "Gemini",
		LocalizedName: "Zwillinge",
		Symbol:        "♊",
	},
	"Cancer": {
		Name:          "Cancer",
		LocalizedName: "Krebs",
		Symbol:        "♋",
	},
	"Leo": {
		Name:          "Leo",
		LocalizedName: "Löwe",
		Symbol:        "♌",
	},
	"Virgo": {
		Name:          "Virgo",
		LocalizedName: "Jungfrau",
		Symbol:        "♍",
	},
	"Libra": {
		Name:          "Libra",
		LocalizedName: "Waage",
		Symbol:        "♎",
	},
	"Scorpio": {
		Name:          "Scorpio",
		LocalizedName: "Skorpion",
		Symbol:        "♏",
	},
	"Sagittarius": {
		Name:          "Sagittarius",
		LocalizedName: "Schütze",
		Symbol:        "♐",
	},
	"Capricornus": {
		Name:          "Capricornus",
		LocalizedName: "Steinbock",
		Symbol:        "♑",
	},
	"Aquarius": {
		Name:          "Aquarius",
		LocalizedName: "Wassermann",
		Symbol:        "♒",
	},
	"Pisces": {
		Name:          "Pisces",
		LocalizedName: "Fische",
		Symbol:        "♓",
	},
}

func (fd FullDate) AstrologicalSign() AstrologicalSign {
	switch fd.Month() {
	case time.January:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Capricornus")
		}
		return makeAstrologicalSign("Aquarius")
	case time.February:
		if fd.Day() <= 19 {
			return makeAstrologicalSign("Aquarius")
		}
		return makeAstrologicalSign("Pisces")
	case time.March:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Pisces")
		}
		return makeAstrologicalSign("Aries")
	case time.April:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Aries")
		}
		return makeAstrologicalSign("Taurus")
	case time.May:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Taurus")
		}
		return makeAstrologicalSign("Gemini")
	case time.June:
		if fd.Day() <= 20 {
			return makeAstrologicalSign("Gemini")
		}
		return makeAstrologicalSign("Cancer")
	case time.July:
		if fd.Day() <= 22 {
			return makeAstrologicalSign("Cancer")
		}
		return makeAstrologicalSign("Leo")
	case time.August:
		if fd.Day() <= 23 {
			return makeAstrologicalSign("Leo")
		}
		return makeAstrologicalSign("Virgo")
	case time.September:
		if fd.Day() <= 23 {
			return makeAstrologicalSign("Virgo")
		}
		return makeAstrologicalSign("Libra")
	case time.October:
		if fd.Day() <= 23 {
			return makeAstrologicalSign("Libra")
		}
		return makeAstrologicalSign("Scorpio")
	case time.November:
		if fd.Day() <= 22 {
			return makeAstrologicalSign("Scorpio")
		}
		return makeAstrologicalSign("Sagittarius")
	case time.December:
		if fd.Day() <= 21 {
			return makeAstrologicalSign("Sagittarius")
		}
		return makeAstrologicalSign("Capricornus")
	}
	return AstrologicalSign{}
}

func makeAstrologicalSign(name string) AstrologicalSign {
	if astrologicalSign, ok := astrologicalSigns[name]; ok {
		return astrologicalSign
	}
	return AstrologicalSign{}
}
