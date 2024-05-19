package main

import "context"

type DiskOperation interface {
	// turn on/off disk alarm light
	Light(ctx context.Context, turnon bool) error
	// kick the specified disk
	Kick(ctx context.Context, deviceName string) error
	// insert disk
	Insert(ctx context.Context, deviceName string) error
	// pull out disk
	Pullout(ctx context.Context, deviceName string) error
	// cancel the kick operation
	CancelKick(ctx context.Context, deviceName string) error
	// replace disk
	Replace(ctx context.Context, oldDeviceName, newDeviceName string) error
}
