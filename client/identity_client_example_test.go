// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package client_test

import (
	"context"

	client "github.com/googleapis/gapic-showcase/client"
	genprotopb "github.com/googleapis/gapic-showcase/server/genproto"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewIdentityClient() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewIdentityRESTClient() {
	ctx := context.Background()
	c, err := client.NewIdentityRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleIdentityClient_CreateUser() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.CreateUserRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/gapic-showcase/server/genproto#CreateUserRequest.
	}
	resp, err := c.CreateUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_GetUser() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.GetUserRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/gapic-showcase/server/genproto#GetUserRequest.
	}
	resp, err := c.GetUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_UpdateUser() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.UpdateUserRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/gapic-showcase/server/genproto#UpdateUserRequest.
	}
	resp, err := c.UpdateUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_DeleteUser() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.DeleteUserRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/gapic-showcase/server/genproto#DeleteUserRequest.
	}
	err = c.DeleteUser(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleIdentityClient_ListUsers() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &genprotopb.ListUsersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/googleapis/gapic-showcase/server/genproto#ListUsersRequest.
	}
	it := c.ListUsers(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleIdentityClient_ListLocations() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	it := c.ListLocations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleIdentityClient_GetLocation() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_SetIamPolicy() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_GetIamPolicy() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_TestIamPermissions() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_ListOperations() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleIdentityClient_GetOperation() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleIdentityClient_DeleteOperation() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.DeleteOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#DeleteOperationRequest.
	}
	err = c.DeleteOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleIdentityClient_CancelOperation() {
	ctx := context.Background()
	c, err := client.NewIdentityClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.CancelOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#CancelOperationRequest.
	}
	err = c.CancelOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
