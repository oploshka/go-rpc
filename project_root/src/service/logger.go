package service

import (
    "fmt"
    "github.com/sirupsen/logrus"
    "os"
    "project-my-test/src/rpc/rpcInterface"
    "strings"
)


type myFormatter struct {
    logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    // this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the logrus levels
    var levelColor int
    switch entry.Level {
    case logrus.DebugLevel, logrus.TraceLevel:
        levelColor = 31 // gray
    case logrus.WarnLevel:
        levelColor = 33 // yellow
    case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
        levelColor = 31 // red
    default:
        levelColor = 36 // blue
    }
    return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}


func CreateLogger() rpcInterface.Logger {
    
    log := &logrus.Logger{
        Out:   os.Stdout,
        Level: logrus.DebugLevel,
        Formatter: &myFormatter{logrus.TextFormatter{
            FullTimestamp:          true,
            TimestampFormat:        "2006-01-02 15:04:05",
            ForceColors:            true,
            DisableLevelTruncation: true,
        },
        },
    }
    // var log = logrus.New()
    // log.Out = os.Stdout
    
    return log
}


/*

// func Logger() *logrus.Logger {
//     f, _ := os.OpenFile("logrus.txt", os.O_CREATE|os.O_WRONLY, 0777)
//     logger := &logrus.Logger{
//         Out:   io.MultiWriter(os.Stderr, f),
//         Level: logrus.InfoLevel,
//         Formatter: &myFormatter{logrus.TextFormatter{
//             FullTimestamp:          true,
//             TimestampFormat:        "2006-01-02 15:04:05",
//             ForceColors:            true,
//             DisableLevelTruncation: true,
//         },
//         },
//     }
//     return logger
// }

func init() {
    // Log as JSON instead of the default ASCII formatter.
    logrus.SetFormatter(&logrus.JSONFormatter{})

    // Output to stdout instead of the default stderr
    // Can be any io.Writer, see below for File example
    logrus.SetOutput(os.Stdout)

    // Only log the warning severity or above.
    logrus.SetLevel(logrus.WarnLevel)
}

func GetLogger() rpcInterface.Logger {
    var log = logrus.New()

    // The API for setting attributes is a little different than the package level
    // exported logger. See Godoc.
    log.Out = os.Stdout

    // You could set this to any `io.Writer` such as a file
    // file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    // if err == nil {
    //  log.Out = file
    // } else {
    //  log.Info("Failed to log to file, using default stderr")
    // }

    return log
}
*/
