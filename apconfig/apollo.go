package apconfig

import (
	"context"
	"errors"
	"github.com/zouyx/agollo"
	"os"
)

const (
	DEFAULT_CLUSTER        = "default"
	DEFAULT_NAMESAPCE_NAME = "application"
)

type Apollo struct {
	AppConfig *agollo.AppConfig
}

func NewApollo(serverName string) (*Apollo, error) {
	metaServerURL := os.Getenv("APOLLO_META_SERVER_URL")
	if metaServerURL == "" {
		return nil, errors.New("Get ENV 'APOLLO_META_SERVER_URL' is empty.")
	}

	return &Apollo{
		AppConfig: &agollo.AppConfig{
			AppId:         serverName,
			Cluster:       DEFAULT_CLUSTER,
			NamespaceName: DEFAULT_NAMESAPCE_NAME,
			Ip:            metaServerURL,
		},
	}, nil
}

func NewCompleteApollo(serverName, cluster, namespaceName string) (*Apollo, error) {
	metaServerURL := os.Getenv("APOLLO_META_SERVER_URL")
	if metaServerURL == "" {
		return nil, errors.New("Get ENV 'APOLLO_META_SERVER_URL' is empty.")
	}

	return &Apollo{
		AppConfig: &agollo.AppConfig{
			AppId:         serverName,
			Cluster:       cluster,
			NamespaceName: namespaceName,
			Ip:            metaServerURL,
		},
	}, nil
}

func (a *Apollo) SetAppId(appId string) *Apollo {
	a.AppConfig.AppId = appId

	return a
}

func (a *Apollo) SetCluster(cluster string) *Apollo {
	a.AppConfig.Cluster = cluster

	return a
}

func (a *Apollo) SetNamespaceName(namespace string) *Apollo {
	a.AppConfig.NamespaceName = namespace

	return a
}

func (a *Apollo) Start() error {
	agollo.InitCustomConfig(func() (*agollo.AppConfig, error) {
		return a.AppConfig, nil
	})

	return agollo.Start()
}

func (a *Apollo) StartWithLogger(ctx context.Context, loggerInterface agollo.LoggerInterface) error {
	agollo.InitCustomConfig(func() (*agollo.AppConfig, error) {
		return a.AppConfig, nil
	})

	return agollo.StartWithLogger(loggerInterface)
}

func (a *Apollo) ListenChangeEvent() <-chan *agollo.ChangeEvent {
	return agollo.ListenChangeEvent()
}

func (a *Apollo) GetStringValue(key, defaultValue string) string {
	return agollo.GetStringValue(key, defaultValue)
}

func (a *Apollo) GetIntValue(key string, defaultValue int) int {
	return agollo.GetIntValue(key, defaultValue)
}

func (a *Apollo) GetFloatValue(key string, defaultValue float64) float64 {
	return agollo.GetFloatValue(key, defaultValue)
}

func (a *Apollo) GetBoolValue(key string, defaultValue bool) bool {
	return agollo.GetBoolValue(key, defaultValue)
}
