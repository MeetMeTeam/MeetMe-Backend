package interfaces

type AvatarResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Assets  []string `json:"assets"`
	Preview string   `json:"preview"`
}
type AvatarService interface {
}
