package Golang_Logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
}

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Hello Logger")
	logger.Debug("Hello Logger")
	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
	//logger.Fatal("Hello Logger")
	//logger.Panic("Hello Logger")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger.SetOutput(file)

	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.Info("Hello Logger")

	logger.WithField("username", "Hanif").Info("Hello Logger")

	logger.WithField("username", "Hanif").
		WithField("name", "Hanif").
		Info("Hello Logger")

	logger.WithFields(logrus.Fields{
		"username": "Hanif",
		"name":     "Hanif",
	}).Info("Hello Logger")

	logger.WithFields(logrus.Fields{
		"username": "Hanif",
		"name":     "Hanif",
	}).WithField("age", 20).Info("Hello Logger")

	logger.WithFields(logrus.Fields{
		"username": "Hanif",
		"name":     "Hanif",
	}).WithFields(logrus.Fields{
		"age":     20,
		"address": "Jakarta",
	}).Info("Hello Logger")

	logger.WithFields(logrus.Fields{
		"username": "Hanif",
		"name":     "Hanif",
	}).WithFields(logrus.Fields{
		"age":     20,
		"address": "Jakarta",
	}).WithField("phone", "08123456789").Info("Hello Logger")

}

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.Info("Hello Logger")
	entry.WithField("username", "Hanif").Info("Hello Logger")
}

type SampleHook struct {
}

func (s SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(SampleHook{})

	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
}
