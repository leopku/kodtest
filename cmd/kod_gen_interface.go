// Code generated by kod struct2interface; DO NOT EDIT.

package cmd

import (
	"context"
)

// demo2Impl is a component that implements IMigration.
type IMigration interface {
	Migrate(context.Context, int) error
}
