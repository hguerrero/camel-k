// +build integration

// To enable compilation of this file in Goland, go to "Settings -> Go -> Vendoring & Build Tags -> Custom Tags" and add "integration"

/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"testing"

	"github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	"github.com/apache/camel-k/pkg/client/clientset/versioned"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func TestClientFunctionalities(t *testing.T) {
	withNewTestNamespace(func(ns string) {
		RegisterTestingT(t)

		cfg, err := config.GetConfig()
		assert.Nil(t, err)
		camel, err := versioned.NewForConfig(cfg)
		assert.Nil(t, err)

		lst, err := camel.CamelV1alpha1().Integrations(ns).List(metav1.ListOptions{})
		assert.Nil(t, err)
		assert.Empty(t, lst.Items)

		integration, err := camel.CamelV1alpha1().Integrations(ns).Create(&v1alpha1.Integration{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dummy",
			},
		})
		assert.Nil(t, err)

		lst, err = camel.CamelV1alpha1().Integrations(ns).List(metav1.ListOptions{})
		assert.Nil(t, err)
		assert.NotEmpty(t, lst.Items)
		assert.Equal(t, lst.Items[0].Name, integration.Name)

		err = camel.CamelV1alpha1().Integrations(ns).Delete("dummy", nil)
		assert.Nil(t, err)
	})
}
