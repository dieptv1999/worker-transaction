package responses

type LoginResponse struct {
	Sessionkey string `json:"sessionkey"`
	Error      *Err   `json:"error"`
	Avatar     string `json:"avatar"`
	Uid        int64  `json:"uid"`
}

type UserImageResponse struct {
	Avatar     string `json:"avatar" xml:"avatar"`
	Background string `json:"background" xml:"background"`
}
