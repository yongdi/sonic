package theme

import "sonic/injection"

func init() {
	injection.Provide(
		NewFileScanner,
		NewPropertyScanner,
		NewMultipartZipThemeFetcher,
	)
}
