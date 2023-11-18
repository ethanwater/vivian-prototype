// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package login

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "vivianlab/internal/pkg/login/T",
		Iface: reflect.TypeOf((*T)(nil)).Elem(),
		Impl:  reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return t_local_stub{impl: impl.(T), tracer: tracer, generateAuthKey2FAMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "vivianlab/internal/pkg/login/T", Method: "GenerateAuthKey2FA", Remote: false}), loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "vivianlab/internal/pkg/login/T", Method: "Login", Remote: false}), verifyAuthKey2FAMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "vivianlab/internal/pkg/login/T", Method: "VerifyAuthKey2FA", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, generateAuthKey2FAMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "vivianlab/internal/pkg/login/T", Method: "GenerateAuthKey2FA", Remote: true}), loginMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "vivianlab/internal/pkg/login/T", Method: "Login", Remote: true}), verifyAuthKey2FAMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "vivianlab/internal/pkg/login/T", Method: "VerifyAuthKey2FA", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return t_reflect_stub{caller: caller}
		},
		RefData: "⟦1139106d:wEaVeReDgE:vivianlab/internal/pkg/login/T→vivianlab/internal/pkg/cache/Cache⟧\n⟦ef8616fa:wEaVeReDgE:vivianlab/internal/pkg/login/T→vivianlab/internal/pkg/auth/T⟧\n⟦830bb057:wEaVeReDgE:vivianlab/internal/pkg/login/T→vivianlab/database/Database⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[T] = (*impl)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*impl)(nil)

// Local stub implementations.

type t_local_stub struct {
	impl                      T
	tracer                    trace.Tracer
	generateAuthKey2FAMetrics *codegen.MethodMetrics
	loginMetrics              *codegen.MethodMetrics
	verifyAuthKey2FAMetrics   *codegen.MethodMetrics
}

// Check that t_local_stub implements the T interface.
var _ T = (*t_local_stub)(nil)

func (s t_local_stub) GenerateAuthKey2FA(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	begin := s.generateAuthKey2FAMetrics.Begin()
	defer func() { s.generateAuthKey2FAMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "login.T.GenerateAuthKey2FA", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GenerateAuthKey2FA(ctx)
}

func (s t_local_stub) Login(ctx context.Context, a0 string, a1 string) (r0 bool, err error) {
	// Update metrics.
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "login.T.Login", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Login(ctx, a0, a1)
}

func (s t_local_stub) VerifyAuthKey2FA(ctx context.Context, a0 string, a1 string) (r0 bool, err error) {
	// Update metrics.
	begin := s.verifyAuthKey2FAMetrics.Begin()
	defer func() { s.verifyAuthKey2FAMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "login.T.VerifyAuthKey2FA", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.VerifyAuthKey2FA(ctx, a0, a1)
}

// Client stub implementations.

type t_client_stub struct {
	stub                      codegen.Stub
	generateAuthKey2FAMetrics *codegen.MethodMetrics
	loginMetrics              *codegen.MethodMetrics
	verifyAuthKey2FAMetrics   *codegen.MethodMetrics
}

// Check that t_client_stub implements the T interface.
var _ T = (*t_client_stub)(nil)

func (s t_client_stub) GenerateAuthKey2FA(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.generateAuthKey2FAMetrics.Begin()
	defer func() { s.generateAuthKey2FAMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "login.T.GenerateAuthKey2FA", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

func (s t_client_stub) Login(ctx context.Context, a0 string, a1 string) (r0 bool, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.loginMetrics.Begin()
	defer func() { s.loginMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "login.T.Login", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Bool()
	err = dec.Error()
	return
}

func (s t_client_stub) VerifyAuthKey2FA(ctx context.Context, a0 string, a1 string) (r0 bool, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.verifyAuthKey2FAMetrics.Begin()
	defer func() { s.verifyAuthKey2FAMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "login.T.VerifyAuthKey2FA", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Bool()
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.22.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// Check that t_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*t_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GenerateAuthKey2FA":
		return s.generateAuthKey2FA
	case "Login":
		return s.login
	case "VerifyAuthKey2FA":
		return s.verifyAuthKey2FA
	default:
		return nil
	}
}

func (s t_server_stub) generateAuthKey2FA(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GenerateAuthKey2FA(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) login(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Login(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Bool(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) verifyAuthKey2FA(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.VerifyAuthKey2FA(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Bool(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type t_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that t_reflect_stub implements the T interface.
var _ T = (*t_reflect_stub)(nil)

func (s t_reflect_stub) GenerateAuthKey2FA(ctx context.Context) (r0 string, err error) {
	err = s.caller("GenerateAuthKey2FA", ctx, []any{}, []any{&r0})
	return
}

func (s t_reflect_stub) Login(ctx context.Context, a0 string, a1 string) (r0 bool, err error) {
	err = s.caller("Login", ctx, []any{a0, a1}, []any{&r0})
	return
}

func (s t_reflect_stub) VerifyAuthKey2FA(ctx context.Context, a0 string, a1 string) (r0 bool, err error) {
	err = s.caller("VerifyAuthKey2FA", ctx, []any{a0, a1}, []any{&r0})
	return
}
