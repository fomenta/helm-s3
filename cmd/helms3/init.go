package main

import (
	"context"

	"github.com/pkg/errors"

	"github.com/fomenta/helm-s3/internal/awss3"
	"github.com/fomenta/helm-s3/internal/awsutil"
	"github.com/fomenta/helm-s3/internal/index"
)

type initAction struct {
	uri string
	acl string
}

func (act initAction) Run(ctx context.Context) error {
	r, err := index.New().Reader()
	if err != nil {
		return errors.WithMessage(err, "get index reader")
	}

	sess, err := awsutil.Session()
	if err != nil {
		return err
	}
	storage := awss3.New(sess)

	if err := storage.PutIndex(ctx, act.uri, act.acl, r); err != nil {
		return errors.WithMessage(err, "upload index to s3")
	}

	// TODO:
	// do we need to automatically do `helm repo add <name> <uri>`,
	// like we are doing `helm repo update` when we push a chart
	// with this plugin?

	return nil
}
