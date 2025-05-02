package wrapper

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMap struct {
	d *v1.ConfigMap
}

// WrapCronJob wraps a v1.CronJob behind a KubernetesObject interface
func WrapConfigMap(i interface{}) (KubernetesObject, error) {
	d, ok := i.(*v1.ConfigMap)

	if !ok {
		return nil, fmt.Errorf("expected v1.ConfigMap received %T", i)
	}

	return &ConfigMap{
		d: d,
	}, nil
}

func (d *ConfigMap) GetMetadata() v1meta.ObjectMeta {
	return d.d.ObjectMeta
}

func (d *ConfigMap) GetObjectSpec() interface{} {
	return d.d.Data
}

func (d *ConfigMap) GetType() string {
	return "ConfigMap"
}
