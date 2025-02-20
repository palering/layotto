// Code generated by github.com/seeflood/protoc-gen-p6 .

// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package email

import (
	fmt "fmt"

	info "mosn.io/layotto/components/pkg/info"
	ref "mosn.io/layotto/components/ref"
)

const (
	serviceName = "email"
)

// Config is the component's configuration
type Config struct {
	ref.Config
	Type     string            `json:"type"`
	Metadata map[string]string `json:"metadata"`
}

type Registry interface {
	Register(fs ...*Factory)
	Create(compType string) (EmailService, error)
}

type Factory struct {
	CompType      string
	FactoryMethod func() EmailService
}

func NewFactory(compType string, f func() EmailService) *Factory {
	return &Factory{
		CompType:      compType,
		FactoryMethod: f,
	}
}

type registry struct {
	stores map[string]func() EmailService
	info   *info.RuntimeInfo
}

func NewRegistry(info *info.RuntimeInfo) Registry {
	info.AddService(serviceName)
	return &registry{
		stores: make(map[string]func() EmailService),
		info:   info,
	}
}

func (r *registry) Register(fs ...*Factory) {
	for _, f := range fs {
		r.stores[f.CompType] = f.FactoryMethod
		r.info.RegisterComponent(serviceName, f.CompType)
	}
}

func (r *registry) Create(compType string) (EmailService, error) {
	if f, ok := r.stores[compType]; ok {
		r.info.LoadComponent(serviceName, compType)
		return f(), nil
	}
	return nil, fmt.Errorf("service component %s is not registered", compType)
}
