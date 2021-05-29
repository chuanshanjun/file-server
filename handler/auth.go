package handler

import "net/http"

func HttpIntercept(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()

			username := r.Form.Get("username")
			token := r.Form.Get("token")

			if len(username) < 3 || !ValidToken(username, token) {
				w.WriteHeader(http.StatusForbidden)
				// 或者跳转到登陆页面
				http.Redirect(w, r, "/user/signin", http.StatusFound)
				return
			}
			// 校验完毕后再执行原有函数
			h(w, r)
		},
	)
}
