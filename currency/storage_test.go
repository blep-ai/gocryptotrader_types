package currency

import "testing"

func TestRunUpdater(t *testing.T) {
	var newStorage Storage

	emptyMainConfig := MainConfiguration{}
	err := newStorage.RunUpdater(BotOverrides{}, &emptyMainConfig, "", false)
