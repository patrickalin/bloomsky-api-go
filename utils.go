package bloomskyStructure

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func logFatal(err error, fct string, msg string, params ...string) {
	logrus.WithFields(logrus.Fields{
		"param": strings.Join(params[:], ","),
		"error": err,
		"fct":   fct,
	}).Fatal(msg)
	log.WithFields(logrus.Fields{
		"param": strings.Join(params[:], ","),
		"error": err,
		"fct":   fct,
	}).Fatal(msg)
}

func logDebug(fct string, msg string, params ...string) {
	log.WithFields(logrus.Fields{
		"param": strings.Join(params[:], ","),
		"fct":   fct,
	}).Debug(msg)
}

func logWarn(fct string, msg string, params ...string) {
	logrus.WithFields(logrus.Fields{
		"param": fmt.Sprintf(strings.Join(params[:], ",")),
	}).Warn(msg)
	log.WithFields(logrus.Fields{
		"param": fmt.Sprintf(strings.Join(params[:], ",")),
		"fct":   fct,
	}).Warn(msg)
}

func checkErr(err error, fct string, msg string, params ...string) {
	if err != nil {
		logFatal(err, msg, fct, fmt.Sprintf(strings.Join(params[:], ",")))
	}
}
