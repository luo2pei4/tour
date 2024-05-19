package main

import (
	"context"
	"errors"
)

type OSDisk struct{}

// turn on/off disk alarm light
func (d *OSDisk) Light(ctx context.Context, turnon bool) error {
	return errors.New("not implemented") // TODO: Implement
}

// kick the specified disk
func (d *OSDisk) Kick(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// insert disk
func (d *OSDisk) Insert(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// pull out disk
func (d *OSDisk) Pullout(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// cancel the kick operation
func (d *OSDisk) CancelKick(ctx context.Context, deviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}

// replace disk
func (d *OSDisk) Replace(ctx context.Context, oldDeviceName string, newDeviceName string) error {
	return errors.New("not implemented") // TODO: Implement
}
