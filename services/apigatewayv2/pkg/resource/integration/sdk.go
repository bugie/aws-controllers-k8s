// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package integration

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.Integration{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetIntegrationWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.IntegrationId != nil {
		ko.Status.IntegrationID = resp.IntegrationId
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		ko.Status.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.APIID == nil || r.ko.Status.IntegrationID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetIntegrationInput, error) {
	res := &svcsdk.GetIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Status.IntegrationID)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.GetIntegrationsInput, error) {
	res := &svcsdk.GetIntegrationsInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateIntegrationWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.IntegrationId != nil {
		ko.Status.IntegrationID = resp.IntegrationId
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		ko.Status.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	}

	ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{OwnerAccountID: &rm.awsAccountID}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateIntegrationInput, error) {
	res := &svcsdk.CreateIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ConnectionID != nil {
		res.SetConnectionId(*r.ko.Spec.ConnectionID)
	}
	if r.ko.Spec.ConnectionType != nil {
		res.SetConnectionType(*r.ko.Spec.ConnectionType)
	}
	if r.ko.Spec.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*r.ko.Spec.ContentHandlingStrategy)
	}
	if r.ko.Spec.CredentialsARN != nil {
		res.SetCredentialsArn(*r.ko.Spec.CredentialsARN)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.IntegrationMethod != nil {
		res.SetIntegrationMethod(*r.ko.Spec.IntegrationMethod)
	}
	if r.ko.Spec.IntegrationType != nil {
		res.SetIntegrationType(*r.ko.Spec.IntegrationType)
	}
	if r.ko.Spec.IntegrationURI != nil {
		res.SetIntegrationUri(*r.ko.Spec.IntegrationURI)
	}
	if r.ko.Spec.PassthroughBehavior != nil {
		res.SetPassthroughBehavior(*r.ko.Spec.PassthroughBehavior)
	}
	if r.ko.Spec.PayloadFormatVersion != nil {
		res.SetPayloadFormatVersion(*r.ko.Spec.PayloadFormatVersion)
	}
	if r.ko.Spec.RequestParameters != nil {
		f11 := map[string]*string{}
		for f11key, f11valiter := range r.ko.Spec.RequestParameters {
			var f11val string
			f11val = *f11valiter
			f11[f11key] = &f11val
		}
		res.SetRequestParameters(f11)
	}
	if r.ko.Spec.RequestTemplates != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range r.ko.Spec.RequestTemplates {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		res.SetRequestTemplates(f12)
	}
	if r.ko.Spec.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*r.ko.Spec.TemplateSelectionExpression)
	}
	if r.ko.Spec.TimeoutInMillis != nil {
		res.SetTimeoutInMillis(*r.ko.Spec.TimeoutInMillis)
	}
	if r.ko.Spec.TLSConfig != nil {
		f15 := &svcsdk.TlsConfigInput{}
		if r.ko.Spec.TLSConfig.ServerNameToVerify != nil {
			f15.SetServerNameToVerify(*r.ko.Spec.TLSConfig.ServerNameToVerify)
		}
		res.SetTlsConfig(f15)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	r *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	input, err := rm.newUpdateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateIntegrationWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.IntegrationId != nil {
		ko.Status.IntegrationID = resp.IntegrationId
	}
	if resp.IntegrationResponseSelectionExpression != nil {
		ko.Status.IntegrationResponseSelectionExpression = resp.IntegrationResponseSelectionExpression
	}

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateIntegrationInput, error) {
	res := &svcsdk.UpdateIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ConnectionID != nil {
		res.SetConnectionId(*r.ko.Spec.ConnectionID)
	}
	if r.ko.Spec.ConnectionType != nil {
		res.SetConnectionType(*r.ko.Spec.ConnectionType)
	}
	if r.ko.Spec.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*r.ko.Spec.ContentHandlingStrategy)
	}
	if r.ko.Spec.CredentialsARN != nil {
		res.SetCredentialsArn(*r.ko.Spec.CredentialsARN)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Status.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Status.IntegrationID)
	}
	if r.ko.Spec.IntegrationMethod != nil {
		res.SetIntegrationMethod(*r.ko.Spec.IntegrationMethod)
	}
	if r.ko.Spec.IntegrationType != nil {
		res.SetIntegrationType(*r.ko.Spec.IntegrationType)
	}
	if r.ko.Spec.IntegrationURI != nil {
		res.SetIntegrationUri(*r.ko.Spec.IntegrationURI)
	}
	if r.ko.Spec.PassthroughBehavior != nil {
		res.SetPassthroughBehavior(*r.ko.Spec.PassthroughBehavior)
	}
	if r.ko.Spec.PayloadFormatVersion != nil {
		res.SetPayloadFormatVersion(*r.ko.Spec.PayloadFormatVersion)
	}
	if r.ko.Spec.RequestParameters != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range r.ko.Spec.RequestParameters {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		res.SetRequestParameters(f12)
	}
	if r.ko.Spec.RequestTemplates != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range r.ko.Spec.RequestTemplates {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		res.SetRequestTemplates(f13)
	}
	if r.ko.Spec.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*r.ko.Spec.TemplateSelectionExpression)
	}
	if r.ko.Spec.TimeoutInMillis != nil {
		res.SetTimeoutInMillis(*r.ko.Spec.TimeoutInMillis)
	}
	if r.ko.Spec.TLSConfig != nil {
		f16 := &svcsdk.TlsConfigInput{}
		if r.ko.Spec.TLSConfig.ServerNameToVerify != nil {
			f16.SetServerNameToVerify(*r.ko.Spec.TLSConfig.ServerNameToVerify)
		}
		res.SetTlsConfig(f16)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteIntegrationWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteIntegrationInput, error) {
	res := &svcsdk.DeleteIntegrationInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Status.IntegrationID)
	}

	return res, nil
}