/*
Copyright 2022 cuisongliu@qq.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"github.com/labring/sealos/pkg/utils/constants"
)

const (
	DefaultUserRoot = "root"
)

var (
	MASTER = "master"
	NODE   = "node"
)

type Provider string

type Arch string

const (
	AMD64  Arch = "amd64"
	ARM64  Arch = "arm64"
	X86_64 Arch = "x86_64"
)

type Protocol string

const (
	ProtocolTCP Protocol = "tcp"
	ProtocolUDP Protocol = "udp"
)

var (
	DefaultConfigPath = constants.GetHomeDir() + "/.sealos"
	DefaultPKFile     = constants.GetHomeDir() + "/.ssh/id_rsa"
)
