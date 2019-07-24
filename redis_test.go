package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

type dummyClient struct {
	exists bool
	err    bool
}

func (c dummyClient) doErr() (err error) {
	if c.err {
		err = fmt.Errorf("some error")
	}

	return
}

func (c dummyClient) doBoolCmd() *redis.BoolCmd {
	return redis.NewBoolResult(c.exists, c.doErr())
}

func (c dummyClient) HGet(string, string) *redis.StringCmd {
	return redis.NewStringResult("", c.doErr())
}

func (c dummyClient) HExists(string, string) *redis.BoolCmd {
	return c.doBoolCmd()
}

func (c dummyClient) HSet(string, string, interface{}) *redis.BoolCmd {
	return c.doBoolCmd()
}

func (c dummyClient) Ping() *redis.StatusCmd {
	return redis.NewStatusResult("", c.doErr())
}

// this exists as a placeholder test, and to stop codecov moaning
//
// In an ideal world, with more time, this would be tested in an
// integration test
func TestNewRedis(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: %+v", err)
		}
	}()

	NewRedis("localhost:6379")
}

func TestRedis_Exists(t *testing.T) {
	for _, test := range []struct {
		name        string
		client      RedisClient
		expect      bool
		expectError bool
	}{
		{"happy path, exists", dummyClient{exists: true}, true, false},
		{"happy path, doesn't exist", dummyClient{}, false, false},
		{"erroring client", dummyClient{err: true}, false, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			exists, err := Redis{test.client}.Exists(Sine{})

			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if test.expect != exists {
				t.Errorf("expected %+v, received %+v", test.expect, exists)
			}
		})
	}
}

func TestRedis_Read(t *testing.T) {
	for _, test := range []struct {
		name        string
		client      RedisClient
		expectError bool
	}{
		{"happy path, exists", dummyClient{exists: true}, false},
		{"happy path, doesn't exist", dummyClient{}, false},
		{"erroring client", dummyClient{err: true}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := Redis{test.client}.Read(Sine{})

			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

		})
	}
}

func TestRedis_Create(t *testing.T) {
	for _, test := range []struct {
		name        string
		client      RedisClient
		expect      bool
		expectError bool
	}{
		{"happy path", dummyClient{exists: true}, true, false},
		{"write failed", dummyClient{}, false, false},
		{"erroring client", dummyClient{err: true}, false, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			exists, err := Redis{test.client}.Write(Sine{}, Chart{b: &bytes.Buffer{}})

			if test.expectError && err == nil {
				t.Errorf("expected error")
			}

			if !test.expectError && err != nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if test.expect != exists {
				t.Errorf("expected %+v, received %+v", test.expect, exists)
			}
		})
	}
}
