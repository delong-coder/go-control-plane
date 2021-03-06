// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/composite/v3/composite.proto

package envoy_extensions_filters_http_composite_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Composite with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Composite) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CompositeValidationError is the validation error returned by
// Composite.Validate if the designated constraints aren't met.
type CompositeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompositeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompositeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompositeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompositeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompositeValidationError) ErrorName() string { return "CompositeValidationError" }

// Error satisfies the builtin error interface
func (e CompositeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComposite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompositeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompositeValidationError{}

// Validate checks the field values on ExecuteFilterAction with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecuteFilterAction) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecuteFilterActionValidationError{
				field:  "TypedConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecuteFilterActionValidationError is the validation error returned by
// ExecuteFilterAction.Validate if the designated constraints aren't met.
type ExecuteFilterActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecuteFilterActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecuteFilterActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecuteFilterActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecuteFilterActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecuteFilterActionValidationError) ErrorName() string {
	return "ExecuteFilterActionValidationError"
}

// Error satisfies the builtin error interface
func (e ExecuteFilterActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecuteFilterAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecuteFilterActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecuteFilterActionValidationError{}
