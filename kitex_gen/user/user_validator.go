// Code generated by Validator v0.1.4. DO NOT EDIT.

package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}
func (p *User) IsValid() error {
	return nil
}
func (p *CheckUserRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *CheckUserResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *CreateUserRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *CreateUserResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *QueryUserRequest) IsValid() error {
	if p.TargetUserId <= int64(0) {
		return fmt.Errorf("field TargetUserId gt rule failed, current value: %v", p.TargetUserId)
	}
	return nil
}
func (p *QueryUserResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
