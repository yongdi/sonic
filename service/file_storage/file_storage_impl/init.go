package file_storage_impl

import "sonic/injection"

func init() {
	injection.Provide(
		NewMinIO,
		NewLocalFileStorage,
		NewAliyun,
	)
}
