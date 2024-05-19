package main

import (
	"context"
	"errors"
)

type CacheDisk struct{}

// turn on/off disk alarm light
func (d *CacheDisk) Light(ctx context.Context, turnon bool) error {
	return errors.New("not implemented") // TODO: Implement
}

// kick the specified disk
func (d *CacheDisk) Kick(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// insert disk
func (d *CacheDisk) Insert(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// pull out disk
func (d *CacheDisk) Pullout(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// cancel the kick operation
func (d *CacheDisk) CancelKick(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// replace disk
func (d *CacheDisk) Replace(ctx context.Context, oldDeviceName string, newDeviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}
