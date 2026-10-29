package tools

import (
	"fmt"
	"log"
	"testing"
)

func TestLogging(t *testing.T) {

	LogWithColor(true)
	log.SetFlags(0)

	//InfoLevel // Default

	Debug("DebugLevel   FAIL")
	Log("LogLevel   FAIL")
	Info("InfoLevel   Default")
	Warn("WarnLevel   Default")
	Error("ErrorLevel   Default")

	LoggingLevel = DebugLevel

	Debug("DebugLevel   Debug")
	Log("LogLevel   Debug")
	Info("InfoLevel   Debug")
	Warn("WarnLevel   Debug")
	Error("ErrorLevel   Debug")

	LoggingLevel = ErrorLevel

	Debug("DebugLevel   FAIL")
	Log("LogLevel   FAIL")
	Info("InfoLevel   FAIL")
	Warn("WarnLevel   FAIL")
	Error("ErrorLevel   Error")

	LogWithColor(false)
	log.SetFlags(log.Ldate | log.Ltime)
	LoggingLevel = DebugLevel

	Debug("DebugLevel   Debug")
	Log("LogLevel   Debug")
	Info("InfoLevel   Debug")
	Warn("WarnLevel   Debug")
	Error("ErrorLevel   Debug")

	LogWithColor(true)
	log.SetFlags(0)
}

func TestCheck(t *testing.T) {
	if !Check(true) {
		t.Fail()
	}
	if Check(false) {
		t.Fail()
	}
	Check(nil)
	CheckError(nil)
	Check(-1)
	Check(t)
	Check("")
	Check(0)
	Check("abc")
	CheckError(fmt.Errorf("erreur"))
	Check(fmt.Errorf("erreur"))
	Check(fmt.Errorf("erreur"), nil)
	obj := struct {
		aaa string
		bbb bool
	}{"aaa", true}
	Check(fmt.Errorf("erreur"), obj, nil, "az", 3)
	Check(nil, "message %s %d", "nb", 234)
}

func TestAssert(t *testing.T) {
	defer Recover()
	if Assert(false) {
		t.Fail()
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		recover()
	}()
	Panic("PanicLevel %s", "arg")
}

func TestFatal(t *testing.T) {
	t.Skip("for debug only")
	defer func() {
		recover()
	}()
	Fatal("Fatal!!!!!")
	t.Fail()
}
