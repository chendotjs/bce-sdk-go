/*
 * Copyright 2020 Baidu, Inc.
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

// client.go - define the client for DDC service

// Package ddc defines the DDC services of BCE. The supported APIs are all defined in sub-package
package ddc

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX               = bce.URI_PREFIX + "v1/ddc"
	DEFAULT_ENDPOINT         = "ddc.su.baidubce.com"
	REQUEST_DDC_INSTANCE_URL = "/instance"
	REQUEST_DDC_POOL_URL     = "/pool"
	REQUEST_DDC_HOST_URL     = "/host"
	REQUEST_DDC_DEPLOY_URL   = "/deploy"
	REQUEST_DDC_DATABASE_URL   = "/database"
	REQUEST_DDC_ACCOUNT_URL   = "/account"
	REQUEST_DDC_RoGroup_URL   = "/roGroup"
)

// Client of DDC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getDdcUri() string {
	return URI_PREFIX
}

func getDdcInstanceUri() string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL
}

// Pool URL
func getPoolUri() string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL
}

func getPoolUriWithId(poolId string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId
}

// Host URL
func getHostUri() string {
	return URI_PREFIX + REQUEST_DDC_HOST_URL
}

func getHostUriWithPnetIp(poolId, hostPnetIP string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId + REQUEST_DDC_HOST_URL + "/" + hostPnetIP
}

// DeploySet URL
func getDeploySetUri(poolId string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId + REQUEST_DDC_DEPLOY_URL
}

func getDeploySetUriWithId(poolId, id string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId + REQUEST_DDC_DEPLOY_URL + "/" + id
}

func getDdcUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId
}

// Database URL
func getDatabaseUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL
}

func getDatabaseUriWithDbName(instanceId string, dbName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL + "/" + dbName
}

// Account URL
func getAccountUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_ACCOUNT_URL
}

func getAccountUriWithAccountName(instanceId string, accountName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_ACCOUNT_URL + "/" + accountName
}

// RoGroup URL
func getRoGroupUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_RoGroup_URL
}
