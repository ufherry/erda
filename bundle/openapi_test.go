// Copyright (c) 2021 Terminus, Inc.
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

package bundle

//import (
//	"fmt"
//	"net/http"
//	"os"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//
//	"github.com/erda-project/erda/apistructs"
//)
//
//func TestBundle_OpenapiOAuth2Token(t *testing.T) {
//	os.Setenv("OPENAPI_ADDR", "localhost:9529")
//	bdl := New(WithOpenapi())
//
//	// get token
//	ti, err := bdl.GetOpenapiOAuth2Token(apistructs.OpenapiOAuth2TokenGetRequest{
//		ClientID:     "pipeline",
//		ClientSecret: "devops/pipeline",
//		Payload: apistructs.OpenapiOAuth2TokenPayload{
//			AccessibleAPIs: []apistructs.AccessibleAPI{
//				{
//					Path:   "/api/pipelines/<pipelineID>/tasks/<taskID>/actions/get-bootstrap-info",
//					Method: http.MethodGet,
//					Schema: "http",
//				},
//			},
//			Metadata: map[string]string{
//				"pipelineID": "10000001",
//				"taskID":     "2",
//			},
//		},
//	})
//	assert.NoError(t, err)
//	fmt.Printf("%+v\n", ti)
//
//	// invalidate token
//	ti, err = bdl.InvalidateOpenapiOAuth2Token(apistructs.OpenapiOAuth2TokenInvalidateRequest{
//		AccessToken: ti.AccessToken,
//	})
//	assert.NoError(t, err)
//	fmt.Printf("%+v\n", ti)
//}
