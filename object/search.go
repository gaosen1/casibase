// Copyright 2023 The Casibase Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"github.com/casibase/casibase/embedding"
)

type SearchProvider interface {
	Search(storeName string, embeddingProviderName string, embeddingProviderObj embedding.EmbeddingProvider, modelProviderName string, text string, knowledgeCount int) ([]Vector, *embedding.EmbeddingResult, error)
}

func GetSearchProvider(typ string, owner string) (SearchProvider, error) {
	var p SearchProvider
	var err error
	if typ == "Default" {
		p, err = NewDefaultSearchProvider(owner)
	} else if typ == "Hierarchy" {
		p, err = NewHierarchySearchProvider(owner)
	} else {
		p, err = NewDefaultSearchProvider(owner)
	}

	if err != nil {
		return nil, err
	}
	return p, nil
}
