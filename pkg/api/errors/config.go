/*
Copyright 2024  The Dapr Authors
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

package errors

import "github.com/dapr/kit/errors"

type ConfigurationError struct {
	name string
}

func Configuration(name string) *ConfigurationError {
	return &ConfigurationError{
		name: name,
	}
}

// func (c *ConfigurationError) StoresNotConfigured() error {
// 	return c.build(
//     errors.NewBuilder(grpcCode codes.Code, httpCode int, message string, tag string)
//     )
// }

func (c *ConfigurationError) build(err *errors.ErrorBuilder, errCode string, metadata map[string]string) error {
	return err.WithErrorInfo(errors.CodePrefixConfigurationStore+errCode, metadata).Build()
}
