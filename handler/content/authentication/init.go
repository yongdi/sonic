package authentication

import "sonic/injection"

func init() {
	injection.Provide(
		NewCategoryAuthentication,
		NewPostAuthentication,
	)
}
