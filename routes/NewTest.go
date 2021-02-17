package diab

import (
	schemas "diab/database"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func checkAge(age string) int {
	if age == "Younger than 40 years" {
		return 0
	} else if age == "40–49 years" {
		return 1
	} else if age == "50–59 years" {
		return 2
	}
	return 3
}

func checkGender(gender string) int {
	if gender == "Male" {
		return 1
	}
	return 0
}

func checkRace(race string) int {
	return 0
}

func checkPrevs(prevs string) int {
	if prevs == "Yes" {
		return 1
	}
	return 0
}

func checkRelatives(relatives string) int {
	if relatives == "Yes" {
		return 1
	}
	return 0
}

func checkBloodPressure(bp string) int {
	if bp == "Yes" {
		return 1
	}
	return 0
}

func checkActive(active string) int {
	if active == "Yes" {
		return 0
	}
	return 1
}

func checkHeightWeight(h string, w string) int {
	if h == `4'10"` {
		if w == "119-142" {
			return 1
		} else if w == "143-190" {
			return 2
		} else if w == "191+" {
			return 3
		}
		return 0
	} else if h == `4'11"` {
		if w == "124-147" {
			return 1
		} else if w == "148-197" {
			return 2
		} else if w == "198+" {
			return 3
		}
		return 0
	} else if h == `5'0"` {
		if w == "128-152" {
			return 1
		} else if w == "153-203" {
			return 2
		} else if w == "204+" {
			return 3
		}
		return 0
	} else if h == `5'1"` {
		if w == "132-157" {
			return 1
		} else if w == "158-210" {
			return 2
		} else if w == "211+" {
			return 3
		}
		return 0
	} else if h == `5'2"` {
		if w == "136-163" {
			return 1
		} else if w == "164-217" {
			return 2
		} else if w == "218+" {
			return 3
		}
		return 0
	} else if h == `5'3"` {
		if w == "141-168" {
			return 1
		} else if w == "169-224" {
			return 2
		} else if w == "225+" {
			return 3
		}
		return 0
	} else if h == `5'4"` {
		if w == "145-173" {
			return 1
		} else if w == "174-231" {
			return 2
		} else if w == "232+" {
			return 3
		}
		return 0
	} else if h == `5'5"` {
		if w == "150-179" {
			return 1
		} else if w == "180-239" {
			return 2
		} else if w == "240+" {
			return 3
		}
		return 0
	} else if h == `5'6"` {
		if w == "155-185" {
			return 1
		} else if w == "186-246" {
			return 2
		} else if w == "247+" {
			return 3
		}
		return 0
	} else if h == `5'7"` {
		if w == "159-190" {
			return 1
		} else if w == "191-254" {
			return 2
		} else if w == "255+" {
			return 3
		}
		return 0
	} else if h == `5'8"` {
		if w == "164-196" {
			return 1
		} else if w == "197-261" {
			return 2
		} else if w == "262+" {
			return 3
		}
		return 0
	} else if h == `5'9"` {
		if w == "169-202" {
			return 1
		} else if w == "203-269" {
			return 2
		} else if w == "270+" {
			return 3
		}
		return 0
	} else if h == `5'10"` {
		if w == "174-208" {
			return 1
		} else if w == "209-277" {
			return 2
		} else if w == "278+" {
			return 3
		}
		return 0
	} else if h == `5'11"` {
		if w == "179-214" {
			return 1
		} else if w == "215-285" {
			return 2
		} else if w == "286+" {
			return 3
		}
		return 0
	} else if h == `6'0"` {
		if w == "184-220" {
			return 1
		} else if w == "221-293" {
			return 2
		} else if w == "294+" {
			return 3
		}
		return 0
	} else if h == `6'1"` {
		if w == "189-226" {
			return 1
		} else if w == "227-301" {
			return 2
		} else if w == "302+" {
			return 3
		}
		return 0
	} else if h == `6'2"` {
		if w == "194-232" {
			return 1
		} else if w == "233-310" {
			return 2
		} else if w == "311+" {
			return 3
		}
		return 0
	} else if h == `6'3"` {
		if w == "200-239" {
			return 1
		} else if w == "240-318" {
			return 2
		} else if w == "319+" {
			return 3
		}
		return 0
	} else if h == `6'4"` {
		if w == "205-245" {
			return 1
		} else if w == "246-327" {
			return 2
		} else if w == "328+" {
			return 3
		}
		return 0
	}
	return 0
}

// NewTest handles POST request for creating new test for user
func NewTest(c *gin.Context, db *gorm.DB) {
	// Gather data from request
	uid := c.PostForm("uid")
	age := c.PostForm("age")
	gender := c.PostForm("gender")
	race := c.PostForm("race")
	prevDiagnosed := c.PostForm("prevDiagnosed")
	relatives := c.PostForm("relatives")
	bloodPressure := c.PostForm("bloodPressure")
	active := c.PostForm("active")
	height := c.PostForm("height")
	weight := c.PostForm("weight")
	// Calculate checks
	score := checkAge(age) + checkGender(gender) + checkRace(race) + checkPrevs(prevDiagnosed) + checkRelatives(relatives) + checkBloodPressure(bloodPressure) + checkActive(active) + checkHeightWeight(height, weight)
	// Creating new test
	test := schemas.Test{
		OwnerID:       uid,
		Age:           age,
		Gender:        gender,
		Race:          race,
		PrevDiagnosed: prevDiagnosed,
		Relatives:     relatives,
		BloodPressure: bloodPressure,
		Active:        active,
		Weight:        weight,
		Height:        height,
		Score:         score,
	}
	db.Create(&test)
	response, _ := json.Marshal(test)
	c.String(http.StatusOK, string(response))
}
