/*
 * Copyright 2021 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package api

const (
	URI_PREFIX    = "/v1"
	URI_USER      = "/user"
	URI_GROUP     = "/group"
	URI_POLICY    = "/policy"
	URI_ACCESSKEY = "/accesskey"
	URI_ROLE      = "/role"

	POLICY_TYPE_SYSTEM = "System"
	POLICY_TYPE_CUSTOM = "Custom"

	ACCESSKEY_STATUS_ENABLE  = "enable"
	ACCESSKEY_STATUS_DISABLE = "disable"
)
