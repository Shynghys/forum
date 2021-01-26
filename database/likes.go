package data

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

type Data struct {
	object uuid.UUID
	user   uuid.UUID
}

func LikeBtn(object, user uuid.UUID) int {

	// object is either a post or a comment
	data := Data{object: object, user: user}
	likeUserSli, ok := data.checkLike()
	if ok {
		return len(likeUserSli)
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
		updateDislike(object, dislikeUserStr)
		updateDisPost(object, len(newDisUserSli))
		updateDisComment(object, len(newDisUserSli))
	}

	// if _, ok = data.checkDislike(); ok {
	// 	return len(likeUserSli)
	// }

	likeUserSli = append(likeUserSli, user.String())
	likeUserStr := strings.Join(likeUserSli, ",")
	updateLike(object, likeUserStr)
	updateLikePost(object, len(likeUserSli))
	updateLikeComment(object, len(likeUserSli))

	return len(likeUserSli)
}

func DislikeBtn(object, user uuid.UUID) int {

	// object is either a post or a comment
	data := Data{object: object, user: user}

	disUserSli, ok := data.checkDislike()
	if ok {
		return len(disUserSli)
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
		updateLike(object, likeUserStr)
		updateLikePost(object, len(newLikeUserSli))
		updateLikeComment(object, len(newLikeUserSli))

	}

	disUserSli = append(disUserSli, user.String())
	dislikeUserStr := strings.Join(disUserSli, ",")
	updateDislike(object, dislikeUserStr)
	updateDisPost(object, len(disUserSli))
	updateDisComment(object, len(disUserSli))

	return len(disUserSli)
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

func updateLike(object uuid.UUID, likeUserStr string) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update likes set authorsID=?, where id=?")
	_, err := smth.Exec(likeUserStr, object)
	CheckErr(err)
	tx.Commit()
}

func updateDislike(object uuid.UUID, disLikeUserStr string) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update dislikes set authorsID=?, where id=?")
	_, err := smth.Exec(disLikeUserStr, object)
	CheckErr(err)
	tx.Commit()
}

func updateLikePost(id uuid.UUID, nbr int) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update posts set likes=?, where id=?")
	_, err := smth.Exec(nbr, id)
	CheckErr(err)
	tx.Commit()
}

func updateDisPost(id uuid.UUID, nbr int) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update posts set dislikes=?, where id=?")
	_, err := smth.Exec(nbr, id)
	CheckErr(err)
	tx.Commit()
}

func updateLikeComment(id uuid.UUID, nbr int) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update comments set likes=?, where id=?")
	_, err := smth.Exec(nbr, id)
	CheckErr(err)
	tx.Commit()
}

func updateDisComment(id uuid.UUID, nbr int) {
	db := DbConn()
	defer db.Close()
	tx, _ := db.Begin()
	smth, _ := tx.Prepare("update comments set dislikes=?, where id=?")
	_, err := smth.Exec(nbr, id)
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
