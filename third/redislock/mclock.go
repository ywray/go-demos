package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/bsm/redislock"
)

type McLock struct {
	rLocker *redislock.Client
	logLock map[string]*redislock.Lock
}

var (
	mLock  *McLock
	mLOnce sync.Once
)

func GetMcLock() *McLock {
	mLOnce.Do(func() {
		mLock = new(McLock)
		client := redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr:    "127.0.0.1:6379",
		})

		mLock.rLocker = redislock.New(client)
		mLock.logLock = make(map[string]*redislock.Lock)
	})
	return mLock
}

func (m *McLock) LogLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	tmpLogLock, err := m.rLocker.Obtain(ctx, key, ttl, nil)
	if err == redislock.ErrNotObtained {
		return false, nil
	} else if err != nil {
		return false, err
	}
	m.logLock[key] = tmpLogLock
	return true, nil
}

func (m *McLock) LogUnlock(ctx context.Context, key string) (err error) {
	if l, ok := m.logLock[key]; ok {
		err = l.Release(ctx)
		if err != nil {
			return err
		}
		delete(m.logLock, key)
		return nil
	}
	return errors.New("release log lock nil")
}

func (m *McLock) LogTtl(ctx context.Context, key string) (ttl time.Duration, err error) {
	if l, ok := m.logLock[key]; ok {
		return l.TTL(ctx)
	}
	return 0, errors.New("release log lock nil")
}

func testLogLock(ctx context.Context) {
	b, err := GetMcLock().LogLock(ctx, "123", 30*time.Second)
	fmt.Println("b", b, "err", err)
	b, err = GetMcLock().LogLock(ctx, "1234", 3*time.Second)
	fmt.Println("b", b, "err", err)
	time.Sleep(4 * time.Second)
	//GetMcLock().LogUnlock(ctx, "123")
	b, err = GetMcLock().LogLock(ctx, "123", 3*time.Second)
	fmt.Println("b1", b, "err1", err)
	b, err = GetMcLock().LogLock(ctx, "1234", 3*time.Second)
	fmt.Println("b2", b, "err1", err)
}
