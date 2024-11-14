package logger

import (
    "io"
    "log"
    "os"
)

func CreateLogger(logFilePath string) (*log.Logger, error) {
    logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return nil, err
    }

    // Use io.MultiWriter instead of log.MultiWriter
    multiWriter := io.MultiWriter(os.Stdout, logFile)
    return log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile), nil
}
