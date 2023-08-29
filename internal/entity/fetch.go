package entity

type FetchUserResp struct {
	Filename string
	Details  []FetchUserRespDetail
}

type FetchUserRespDetail struct {
	Name string
	Err  error
}
