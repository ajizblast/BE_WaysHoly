package donationdto

type DonationRequest struct{
	Title string `json:"title" form:"title"`
	Goal  int `json:"goal" form:"goal"`
	Description string `json:"description" form:"description"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}