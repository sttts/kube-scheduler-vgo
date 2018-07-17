package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Foo struct {
	metav1.TypeMeta
}
