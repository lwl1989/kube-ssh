package types

type PageCommonReq struct {
	Page  int `json:"page,optional"`
	Size  int `json:"size,optional"`
	PreId int `json:"pre_id,optional"`
}

func (req PageCommonReq) GetLimit() int {
	if req.Size > 0 {
		return req.Size
	}
	return 10
}

func (req PageCommonReq) GetPage() int {
	if req.Page > 0 {
		return req.Page
	}
	return 1
}

func (req PageCommonReq) GetOffset() int {
	return (req.GetPage() - 1) * req.GetLimit()
}

func (req PageCommonReq) GetPreId() int {
	return req.PreId
}
