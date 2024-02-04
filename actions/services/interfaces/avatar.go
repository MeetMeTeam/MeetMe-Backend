package interfaces

type AvatarResponse struct {
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
}
type AvatarService interface {
}
