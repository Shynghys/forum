package data

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Data struct {
	object uuid.UUID
	user   uuid.UUID
}

func LikeBtn(object, user uuid.UUID, count int) int {

	// object is either a post or a comment
	data := Data{object: object, user: user}
	likeUserSli, ok := data.checkLike()
	if ok {
		return count
	}

	disUserSli, ok := data.checkDislike()
	if ok {
		var newDisUserSli []string
		for _, val := range disUserSli {
			if val != user.String() {
				newDisUserSli = append(newDisUserSli, val)
			}
		}

		dislikeUserStr := strings.Join(newDisUserSli, ",")
		updateDislike(object.String(), dislikeUserStr)
		count++
	}
	likeUserSli = append(likeUserSli, user.String())
	likeUserStr := strings.Join(likeUserSli, ",")
	updateLike(object.String(), likeUserStr)

	return count + 1
}

func DislikeBtn(object, user uuid.UUID, count int) int {

	// object is either a post or a comment
	data := Data{object: object, user: user}

	disUserSli, ok := data.checkDislike()
	if ok {
		return count
	}

	likeUserSli, ok := data.checkLike()
	if ok {
		var newLikeUserSli []string
		for _, val := range likeUserSli {
			if val != user.String() {
				newLikeUserSli = append(newLikeUserSli, val)
			}
		}

		likeUserStr := strings.Join(newLikeUserSli, ",")
		updateLike(object.String(), likeUserStr)
		count--
	}
	disUserSli = append(disUserSli, user.String())
	dislikeUserStr := strings.Join(disUserSli, ",")
	updateDislike(object.String(), dislikeUserStr)

	return count - 1
}

func isIn(str string, sli []string) bool {
	for _, val := range sli {
		if val == str {
			return true
		}
	}
	return false
}

func (val *Data) checkLike() ([]string, bool) {
	db := DbConn()
	defer db.Close()

	row, err := db.Query("SELECT authors FROM likes WHERE id LIKE ?", val.object)
	CheckErr(err)
	defer row.Close()
	var authors string
	for row.Next() {
		row.Scan(&authors)
	}

	sli := strings.Split(authors, ",")

	userStr := val.user.String()

	return sli, isIn(userStr, sli)

}

func (val *Data) checkDislike() ([]string, bool) {
	db := DbConn()
	defer db.Close()

	row, err := db.Query("SELECT authors FROM dislikes WHERE id LIKE ?", val.object)
	CheckErr(err)
	defer row.Close()
	var authors string
	for row.Next() {
		row.Scan(&authors)
	}

	sli := strings.Split(authors, ",")

	userStr := val.user.String()

	return sli, isIn(userStr, sli)
}

func updateLike(object, likeUserStr string) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update likes set authorsID=?, where id=?")
	_, err := smth.Exec(likeUserStr, object)
	CheckErr(err)
	tx.Commit()
}

func updateDislike(object, likeUserStr string) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update dislikes set authorsID=?, where id=?")
	_, err := smth.Exec(likeUserStr, object)
	CheckErr(err)
	tx.Commit()
}

/*
func checkLike(object, user uuid.UUID) bool {
	db := DbConn()
	defer db.Close()

	// object is either a post or a comment
	row, err := db.Query("SELECT authors FROM likes WHERE id LIKE ?", object)
	CheckErr(err)
	defer row.Close()
	var authors string
	for row.Next() {
		row.Scan(&authors)
	}

	sli := strings.Split(authors, ",")

	userStr := user.String()

	return isIn(userStr, sli)

}

func checkDislike(object, user uuid.UUID) bool {
	db := DbConn()
	defer db.Close()

	// object is either a post or a comment
	row, err := db.Query("SELECT authors FROM dislikes WHERE id LIKE ?", object)
	CheckErr(err)
	defer row.Close()
	var authors string
	for row.Next() {
		row.Scan(&authors)
	}

	sli := strings.Split(authors, ",")

	userStr := user.String()

	return isIn(userStr, sli)
}
*/
