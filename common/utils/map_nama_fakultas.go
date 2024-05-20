package utils

import (
	"strings"
)

func GetFullNamaFak(acronym string) string {
	facultyMap := map[string]string{
		"FIP":         "ILMU PENDIDIKAN",
		"FPIPS":       "PENDIDIKAN ILMU PENGETAHUAN SOSIAL",
		"FPBS":        "PENDIDIKAN BAHASA DAN SASTRA",
		"FPMIPA":      "PENDIDIKAN MATEMATIKA DAN ILMU PENGETAHUAN ALAM",
		"FPTK":        "PENDIDIKAN TEKNOLOGI DAN KEJURUAN",
		"FPOK":        "PENDIDIKAN OLAHRAGA DAN KESEHATAN",
		"CIBIRU":      "CIBIRU",
		"SUMEDANG":    "SUMEDANG",
		"PURWAKARTA":  "PURWAKARTA",
		"TASIKMALAYA": "TASIKMALAYA",
		"SERANG":      "SERANG",
		"FPEB":        "PENDIDIKAN EKONOMI DAN BISNIS",
		"FPSD":        "PENDIDIKAN SENI DAN DESAIN",
		"SPS":         "PASCASARJANA",
	}

	fullName, found := facultyMap[strings.ToUpper(acronym)]
	if !found {
		return acronym
	}

	return fullName
}
