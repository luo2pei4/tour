package main

import (
	"context"
	"errors"
)

type DataDisk struct{}

// turn on/off disk alarm light
func (d *DataDisk) Light(ctx context.Context, turnon bool) error {
	return errors.New("not implemented") // TODO: Implement
}

// kick the specified disk
func (d *DataDisk) Kick(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// insert disk
func (d *DataDisk) Insert(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// pull out disk
func (d *DataDisk) Pullout(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// cancel the kick operation
func (d *DataDisk) CancelKick(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// replace disk
func (d *DataDisk) Replace(ctx context.Context, oldDeviceName string, newDeviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}
