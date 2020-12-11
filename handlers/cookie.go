package handlers

import "net/http"

func GetCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		return ""
	}
	return c.Value
}
func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   COOKIE_NAME,
		MaxAge: -1}
	http.SetCookie(w, &c)

}
