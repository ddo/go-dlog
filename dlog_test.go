package dlog

import (
	"testing"
)

func TestJSON(t *testing.T) {
	_dlog := New("dlog", &Option{Type: "json"})

	_dlog.Debug("Debug")
	_dlog.Debug("Debug", "Debug1")
	_dlog.Debug("Debug", "Debug1", "Debug2")
	_dlog.Debug("Debug", "Debug1", "Debug2", "Debug3")

	_dlog.Info("Info")
	_dlog.Info("Info", "Info1")
	_dlog.Info("Info", "Info1", "Info2")
	_dlog.Info("Info", "Info1", "Info2", "Info3")

	_dlog.Done("Done")
	_dlog.Done("Done", "Done1")
	_dlog.Done("Done", "Done1", "Done2")
	_dlog.Done("Done", "Done1", "Done2", "Done3")

	_dlog.Fail("Fail")
	_dlog.Fail("Fail", "Fail1")
	_dlog.Fail("Fail", "Fail1", "Fail2")
	_dlog.Fail("Fail", "Fail1", "Fail2", "Fail3")

	_dlog.Warn("Warn")
	_dlog.Warn("Warn", "Warn1")
	_dlog.Warn("Warn", "Warn1", "Warn2")
	_dlog.Warn("Warn", "Warn1", "Warn2", "Warn3")

	_dlog.Error("Error")
	_dlog.Error("Error", "Error1")
	_dlog.Error("Error", "Error1", "Error2")
	_dlog.Error("Error", "Error1", "Error2", "Error3")
}

func TestSimple(t *testing.T) {
	_dlog := New("dlog", &Option{Type: "simple"})

	_dlog.Debug("Debug")
	_dlog.Debug("Debug", "Debug1")
	_dlog.Debug("Debug", "Debug1", "Debug2")
	_dlog.Debug("Debug", "Debug1", "Debug2", "Debug3")

	_dlog.Info("Info")
	_dlog.Info("Info", "Info1")
	_dlog.Info("Info", "Info1", "Info2")
	_dlog.Info("Info", "Info1", "Info2", "Info3")

	_dlog.Done("Done")
	_dlog.Done("Done", "Done1")
	_dlog.Done("Done", "Done1", "Done2")
	_dlog.Done("Done", "Done1", "Done2", "Done3")

	_dlog.Fail("Fail")
	_dlog.Fail("Fail", "Fail1")
	_dlog.Fail("Fail", "Fail1", "Fail2")
	_dlog.Fail("Fail", "Fail1", "Fail2", "Fail3")

	_dlog.Warn("Warn")
	_dlog.Warn("Warn", "Warn1")
	_dlog.Warn("Warn", "Warn1", "Warn2")
	_dlog.Warn("Warn", "Warn1", "Warn2", "Warn3")

	_dlog.Error("Error")
	_dlog.Error("Error", "Error1")
	_dlog.Error("Error", "Error1", "Error2")
	_dlog.Error("Error", "Error1", "Error2", "Error3")
}

func TestColor(t *testing.T) {
	_dlog := New("dlog", nil)

	_dlog.Debug("Debug")
	_dlog.Debug("Debug", "Debug1")
	_dlog.Debug("Debug", "Debug1", "Debug2")
	_dlog.Debug("Debug", "Debug1", "Debug2", "Debug3")

	_dlog.Info("Info")
	_dlog.Info("Info", "Info1")
	_dlog.Info("Info", "Info1", "Info2")
	_dlog.Info("Info", "Info1", "Info2", "Info3")

	_dlog.Done("Done")
	_dlog.Done("Done", "Done1")
	_dlog.Done("Done", "Done1", "Done2")
	_dlog.Done("Done", "Done1", "Done2", "Done3")

	_dlog.Fail("Fail")
	_dlog.Fail("Fail", "Fail1")
	_dlog.Fail("Fail", "Fail1", "Fail2")
	_dlog.Fail("Fail", "Fail1", "Fail2", "Fail3")

	_dlog.Warn("Warn")
	_dlog.Warn("Warn", "Warn1")
	_dlog.Warn("Warn", "Warn1", "Warn2")
	_dlog.Warn("Warn", "Warn1", "Warn2", "Warn3")

	_dlog.Error("Error")
	_dlog.Error("Error", "Error1")
	_dlog.Error("Error", "Error1", "Error2")
	_dlog.Error("Error", "Error1", "Error2", "Error3")
}
