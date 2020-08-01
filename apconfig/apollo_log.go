package apconfig

import (
	"context"
	"fmt"
	"gitee.com/kelvins-io/common/log"
)

type ApolloLog struct {
	Logger     *log.LoggerContext
	ServerName string
}

func NewApolloLog(serverName string) (*ApolloLog, error) {
	logger, err := log.GetBusinessLogger("apollo_config")
	if err != nil {
		return nil, err
	}

	return &ApolloLog{
		Logger:     logger,
		ServerName: serverName,
	}, nil
}

func (l *ApolloLog) Debugf(format string, params ...interface{}) {
	l.Logger.Debugf(context.Background(), format, params)
}

func (l *ApolloLog) Infof(format string, params ...interface{}) {
	l.Logger.Infof(context.Background(), format, params)
}

func (l *ApolloLog) Warnf(format string, params ...interface{}) error {
	l.Logger.Warnf(context.Background(), format, params)
	return nil
}

func (l *ApolloLog) Errorf(format string, params ...interface{}) error {
	l.Logger.Errorf(context.Background(), format, params)
	return nil
}

func (l *ApolloLog) Debug(v ...interface{}) {
	l.Logger.Debug(context.Background(), fmt.Sprint(v))
}

func (l *ApolloLog) Info(v ...interface{}) {
	l.Logger.Info(context.Background(), fmt.Sprint(v))
}

func (l *ApolloLog) Warn(v ...interface{}) error {
	l.Logger.Warn(context.Background(), fmt.Sprint(v))
	return nil
}

func (l *ApolloLog) Error(v ...interface{}) error {
	l.Logger.Error(context.Background(), fmt.Sprint(v))
	return nil
}
