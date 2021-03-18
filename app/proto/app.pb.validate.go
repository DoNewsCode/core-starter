// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: app.proto

package app_pb

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _app_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetOneUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetOneUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return GetOneUserRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// GetOneUserRequestValidationError is the validation error returned by
// GetOneUserRequest.Validate if the designated constraints aren't met.
type GetOneUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOneUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOneUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOneUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOneUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOneUserRequestValidationError) ErrorName() string {
	return "GetOneUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOneUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOneUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOneUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOneUserRequestValidationError{}

// Validate checks the field values on UserInfoReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserInfoReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for CreatedAt

	return nil
}

// UserInfoReplyValidationError is the validation error returned by
// UserInfoReply.Validate if the designated constraints aren't met.
type UserInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoReplyValidationError) ErrorName() string { return "UserInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoReplyValidationError{}