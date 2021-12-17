package k8s

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// Client is a kubernetes client interface used internally. It copies functions from
// sigs.k8s.io/controller-runtime/pkg/client
//
//counterfeiter:generate . Client
type Client interface {
	Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error
	Get(ctx context.Context, key client.ObjectKey, obj client.Object) error

	Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error
	Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error
	DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error
	List(ctx context.Context, obj client.ObjectList, opts ...client.ListOption) error
	Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error

	RESTMapper() meta.RESTMapper
	Scheme() *runtime.Scheme

	Status() client.StatusWriter
}

// StatusWriter is a kubernetes status writer interface used internally. It copies functions from
// sigs.k8s.io/controller-runtime/pkg/client
//
//counterfeiter:generate . StatusWriter
type StatusWriter interface {
	Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error
	Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error
}