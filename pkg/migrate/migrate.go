package migrate

import "context"

type IMigrate interface {
	Migrate(context.Context, int) error
}
