package diab

import (
	schemas "diab/database"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"crypto/sha256"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type docsChats struct {
	DID    int
	Amount int
	Ready  bool
}

// NewSHA256 handles sha256-ing passwords
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// NewUser handles POST request for creating new users (sign up)
func NewUser(c *gin.Context, db *gorm.DB) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	shaed := NewSHA256([]byte(password))
	// Creating new user
	user := schemas.User{
		Name:     name,
		Email:    email,
		Password: hex.EncodeToString(shaed),
	}
	db.Create(&user)
	var doctors []schemas.Doctor
	var chats []schemas.Chat
	var docIDs []int
	var dchats []docsChats
	db.Find(&doctors)
	for i := range doctors {
		docIDs = append(docIDs, int(doctors[i].ID))
		ready, _ := strconv.ParseBool(doctors[i].Ready)
		dchats = append(dchats, docsChats{DID: int(doctors[i].ID), Amount: 0, Ready: ready})
	}
	db.Where(map[string]interface{}{"doctorId": docIDs}).Find(&chats)
	for i := range doctors {
		for y := range chats {
			did, _ := strconv.ParseInt(chats[y].DoctorID, 10, 64)
			if int(did) == int(doctors[i].ID) {
				dchats[i].Amount = dchats[i].Amount + 1
			}
		}
	}
	min := []int{0, dchats[0].Amount}
	for i := range dchats {
		if min[1] > dchats[i].Amount {
			min = []int{i, dchats[i].Amount}
		}
	}
	chat := schemas.Chat{
		PatientID: strconv.Itoa(int(user.ID)),
		DoctorID:  strconv.Itoa(dchats[min[0]].DID),
	}
	db.Create(&chat)
	response, _ := json.Marshal(user)
	c.String(http.StatusOK, string(response))
}
