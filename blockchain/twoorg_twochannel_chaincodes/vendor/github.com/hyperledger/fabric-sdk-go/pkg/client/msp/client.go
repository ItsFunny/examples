/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package msp enables creation and update of users on a Fabric network.
// Msp client supports the following actions:
// Enroll, Reenroll, Register,  Revoke and GetSigningIdentity.
//
//  Basic Flow:
//  1) Prepare client context
//  2) Create msp client
//  3) Register user
//  4) Enroll user
package msp

import (
	"fmt"

	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	mspctx "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/msp"
	mspapi "github.com/hyperledger/fabric-sdk-go/pkg/msp/api"
	"github.com/pkg/errors"
)

// Client enables access to Client services
type Client struct {
	orgName string
	caName  string
	ctx     context.Client
}

// ClientOption describes a functional parameter for the New constructor
type ClientOption func(*Client) error

// WithOrg option
func WithOrg(orgName string) ClientOption {
	return func(msp *Client) error {
		msp.orgName = orgName
		return nil
	}
}

// opts allows the user to specify more advanced request options
type requestOptions struct {
	CA string
}

// RequestOption func for each Opts argument
type RequestOption func(ctx context.Client, opts *requestOptions) error

// WithCA allows for specifying optional CA name
func WithCA(caname string) RequestOption {
	return func(ctx context.Client, o *requestOptions) error {
		o.CA = caname
		return nil
	}
}

// New creates a new Client instance
func New(clientProvider context.ClientProvider, opts ...ClientOption) (*Client, error) {

	ctx, err := clientProvider()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create Client")
	}

	msp := Client{
		ctx: ctx,
	}

	for _, param := range opts {
		err := param(&msp)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create Client")
		}
	}

	if msp.orgName == "" {
		msp.orgName = ctx.IdentityConfig().Client().Organization
	}
	if msp.orgName == "" {
		return nil, errors.New("organization is not provided")
	}

	caConfig, ok := ctx.IdentityConfig().CAConfig(msp.orgName)
	if ok {
		msp.caName = caConfig.CAName
	}

	networkConfig := ctx.EndpointConfig().NetworkConfig()
	_, ok = networkConfig.Organizations[strings.ToLower(msp.orgName)]
	if !ok {
		return nil, fmt.Errorf("non-existent organization: '%s'", msp.orgName)
	}
	return &msp, nil
}

func newCAClient(ctx context.Client, orgName string) (mspapi.CAClient, error) {

	caClient, err := msp.NewCAClient(orgName, ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create CA Client")
	}

	return caClient, nil
}

// enrollmentOptions represent enrollment options
type enrollmentOptions struct {
	secret   string
	profile  string
	label    string
	typ      string
	attrReqs []*AttributeRequest
}

// EnrollmentOption describes a functional parameter for Enroll
type EnrollmentOption func(*enrollmentOptions) error

// WithSecret enrollment option
func WithSecret(secret string) EnrollmentOption {
	return func(o *enrollmentOptions) error {
		o.secret = secret
		return nil
	}
}

// WithProfile enrollment option
func WithProfile(profile string) EnrollmentOption {
	return func(o *enrollmentOptions) error {
		o.profile = profile
		return nil
	}
}

// WithType enrollment option
func WithType(typ string) EnrollmentOption {
	return func(o *enrollmentOptions) error {
		o.typ = typ
		return nil
	}
}

// WithLabel enrollment option
func WithLabel(label string) EnrollmentOption {
	return func(o *enrollmentOptions) error {
		o.label = label
		return nil
	}
}

// WithAttributeRequests enrollment option
func WithAttributeRequests(attrReqs []*AttributeRequest) EnrollmentOption {
	return func(o *enrollmentOptions) error {
		o.attrReqs = attrReqs
		return nil
	}
}

// CreateIdentity creates a new identity with the Fabric CA server. An enrollment secret is returned which can then be used,
// along with the enrollment ID, to enroll a new identity.
//  Parameters:
//  request holds info about identity
//
//  Returns:
//  Return identity info including the secret
func (c *Client) CreateIdentity(request *IdentityRequest) (*IdentityResponse, error) {

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	var attrs []mspapi.Attribute
	for i := range request.Attributes {
		attrs = append(attrs, mspapi.Attribute{Name: request.Attributes[i].Name, Value: request.Attributes[i].Value, ECert: request.Attributes[i].ECert})
	}

	req := &mspapi.IdentityRequest{
		ID:             request.ID,
		Type:           request.Type,
		MaxEnrollments: request.MaxEnrollments,
		Affiliation:    request.Affiliation,
		Attributes:     attrs,
		CAName:         request.CAName,
		Secret:         request.Secret,
	}

	response, err := ca.CreateIdentity(req)
	if err != nil {
		return nil, err
	}

	return getIdentityResponse(response), nil
}

// ModifyIdentity modifies identity with the Fabric CA server.
//  Parameters:
//  request holds info about identity
//
//  Returns:
//  Return updated identity info
func (c *Client) ModifyIdentity(request *IdentityRequest) (*IdentityResponse, error) {

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	var attrs []mspapi.Attribute
	for i := range request.Attributes {
		attrs = append(attrs, mspapi.Attribute{Name: request.Attributes[i].Name, Value: request.Attributes[i].Value, ECert: request.Attributes[i].ECert})
	}

	req := &mspapi.IdentityRequest{
		ID:             request.ID,
		Type:           request.Type,
		MaxEnrollments: request.MaxEnrollments,
		Affiliation:    request.Affiliation,
		Attributes:     attrs,
		CAName:         request.CAName,
		Secret:         request.Secret,
	}

	response, err := ca.ModifyIdentity(req)
	if err != nil {
		return nil, err
	}

	return getIdentityResponse(response), nil
}

// RemoveIdentity removes identity with the Fabric CA server.
//  Parameters:
//  request holds info about identity to be removed
//
//  Returns:
//  Return removed identity info
func (c *Client) RemoveIdentity(request *RemoveIdentityRequest) (*IdentityResponse, error) {

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	req := &mspapi.RemoveIdentityRequest{
		ID:     request.ID,
		Force:  request.Force,
		CAName: request.CAName,
	}

	response, err := ca.RemoveIdentity(req)
	if err != nil {
		return nil, err
	}

	return getIdentityResponse(response), nil
}

// GetAllIdentities returns all identities that the caller is authorized to see
//  Parameters:
//  options holds optional request options
//  Returns:
//  Response containing identities
func (c *Client) GetAllIdentities(options ...RequestOption) ([]*IdentityResponse, error) {

	// Read request options
	opts, err := c.prepareOptsFromOptions(c.ctx, options...)
	if err != nil {
		return nil, err
	}

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	responses, err := ca.GetAllIdentities(opts.CA)
	if err != nil {
		return nil, err
	}

	return getIdentityResponses(responses), nil

}

// GetIdentity retrieves identity information.
//  Parameters:
//  ID is required identity ID
//  options holds optional request options
//
//  Returns:
//  Response containing identity information
func (c *Client) GetIdentity(ID string, options ...RequestOption) (*IdentityResponse, error) {

	// Read request options
	opts, err := c.prepareOptsFromOptions(c.ctx, options...)
	if err != nil {
		return nil, err
	}

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	response, err := ca.GetIdentity(ID, opts.CA)
	if err != nil {
		return nil, err
	}

	return getIdentityResponse(response), nil

}

func getIdentityResponse(response *mspapi.IdentityResponse) *IdentityResponse {

	var attributes []Attribute
	for i := range response.Attributes {
		attributes = append(attributes, Attribute{Name: response.Attributes[i].Name, Value: response.Attributes[i].Value, ECert: response.Attributes[i].ECert})
	}

	res := &IdentityResponse{ID: response.ID,
		Affiliation:    response.Affiliation,
		Type:           response.Type,
		Attributes:     attributes,
		MaxEnrollments: response.MaxEnrollments,
		Secret:         response.Secret,
		CAName:         response.CAName,
	}

	return res
}

func getIdentityResponses(responses []*mspapi.IdentityResponse) []*IdentityResponse {

	ret := make([]*IdentityResponse, len(responses))
	for i, r := range responses {
		ret[i] = getIdentityResponse(r)
	}

	return ret
}

// Enroll enrolls a registered user in order to receive a signed X509 certificate.
// A new key pair is generated for the user. The private key and the
// enrollment certificate issued by the CA are stored in SDK stores.
// They can be retrieved by calling IdentityManager.GetSigningIdentity().
//  Parameters:
//  enrollmentID enrollment ID of a registered user
//  opts are optional enrollment options
//
//  Returns:
//  an error if enrollment fails
func (c *Client) Enroll(enrollmentID string, opts ...EnrollmentOption) error {

	eo := enrollmentOptions{}
	for _, param := range opts {
		err := param(&eo)
		if err != nil {
			return errors.WithMessage(err, "failed to enroll")
		}
	}

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return err
	}

	req := &mspapi.EnrollmentRequest{
		Name:    enrollmentID,
		Secret:  eo.secret,
		Profile: eo.profile,
		Type:    eo.typ,
		Label:   eo.label,
	}

	if len(eo.attrReqs) > 0 {
		attrs := make([]*mspapi.AttributeRequest, 0)
		for _, attr := range eo.attrReqs {
			attrs = append(attrs, &mspapi.AttributeRequest{Name: attr.Name, Optional: attr.Optional})
		}
		req.AttrReqs = attrs
	}

	return ca.Enroll(req)
}

// Reenroll reenrolls an enrolled user in order to obtain a new signed X509 certificate
//  Parameters:
//  enrollmentID enrollment ID of a registered user
//
//  Returns:
//  an error if re-enrollment fails
func (c *Client) Reenroll(enrollmentID string, opts ...EnrollmentOption) error {
	eo := enrollmentOptions{}
	for _, param := range opts {
		err := param(&eo)
		if err != nil {
			return errors.WithMessage(err, "failed to enroll")
		}
	}

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return err
	}

	req := &mspapi.ReenrollmentRequest{
		Name:    enrollmentID,
		Profile: eo.profile,
		Label:   eo.label,
	}
	if len(eo.attrReqs) > 0 {
		attrs := make([]*mspapi.AttributeRequest, 0)
		for _, attr := range eo.attrReqs {
			attrs = append(attrs, &mspapi.AttributeRequest{Name: attr.Name, Optional: attr.Optional})
		}
		req.AttrReqs = attrs
	}
	return ca.Reenroll(req)
}

// Register registers a User with the Fabric CA
//  Parameters:
//  request is registration request
//
//  Returns:
//  enrolment secret
func (c *Client) Register(request *RegistrationRequest) (string, error) {
	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return "", err
	}

	var a []mspapi.Attribute
	for i := range request.Attributes {
		a = append(a, mspapi.Attribute{Name: request.Attributes[i].Name, Value: request.Attributes[i].Value, ECert: request.Attributes[i].ECert})
	}

	r := mspapi.RegistrationRequest{
		Name:           request.Name,
		Type:           request.Type,
		MaxEnrollments: request.MaxEnrollments,
		Affiliation:    request.Affiliation,
		Attributes:     a,
		CAName:         request.CAName,
		Secret:         request.Secret,
	}
	return ca.Register(&r)
}

// Revoke revokes a User with the Fabric CA
//  Parameters:
//  request is revocation request
//
//  Returns:
//  revocation response
func (c *Client) Revoke(request *RevocationRequest) (*RevocationResponse, error) {
	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}
	req := mspapi.RevocationRequest(*request)
	resp, err := ca.Revoke(&req)
	if err != nil {
		return nil, err
	}
	var revokedCerts []RevokedCert
	for i := range resp.RevokedCerts {
		revokedCerts = append(
			revokedCerts,
			RevokedCert{
				Serial: resp.RevokedCerts[i].Serial,
				AKI:    resp.RevokedCerts[i].AKI,
			})
	}

	return &RevocationResponse{
		RevokedCerts: revokedCerts,
		CRL:          resp.CRL,
	}, nil
}

// GetCAInfo returns generic CA information
func (c *Client) GetCAInfo() (*GetCAInfoResponse, error) {
	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	resp, err := ca.GetCAInfo()
	if err != nil {
		return nil, err
	}

	return &GetCAInfoResponse{CAName: resp.CAName, CAChain: resp.CAChain[:], IssuerPublicKey: resp.IssuerPublicKey[:], IssuerRevocationPublicKey: resp.IssuerRevocationPublicKey[:], Version: resp.Version}, nil
}

// GetSigningIdentity returns signing identity for id
//  Parameters:
//  id is user id
//
//  Returns:
//  signing identity
func (c *Client) GetSigningIdentity(id string) (mspctx.SigningIdentity, error) {
	im, _ := c.ctx.IdentityManager(c.orgName)
	si, err := im.GetSigningIdentity(id)
	if err != nil {
		if err == mspctx.ErrUserNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return si, nil
}

// CreateSigningIdentity creates a signing identity with the given options
func (c *Client) CreateSigningIdentity(opts ...mspctx.SigningIdentityOption) (mspctx.SigningIdentity, error) {
	im, _ := c.ctx.IdentityManager(c.orgName)
	return im.CreateSigningIdentity(opts...)
}

//prepareOptsFromOptions reads request options from Option array
func (c *Client) prepareOptsFromOptions(ctx context.Client, options ...RequestOption) (requestOptions, error) {
	opts := requestOptions{}
	for _, option := range options {
		err := option(ctx, &opts)
		if err != nil {
			return opts, errors.WithMessage(err, "Failed to read opts")
		}
	}
	return opts, nil
}

// GetAffiliation returns information about the requested affiliation
func (c *Client) GetAffiliation(affiliation string, options ...RequestOption) (*AffiliationResponse, error) {
	// Read request options
	opts, err := c.prepareOptsFromOptions(c.ctx, options...)
	if err != nil {
		return nil, err
	}

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	r, err := ca.GetAffiliation(affiliation, opts.CA)
	if err != nil {
		return nil, err
	}

	resp := &AffiliationResponse{CAName: r.CAName, AffiliationInfo: AffiliationInfo{}}
	err = fillAffiliationInfo(&resp.AffiliationInfo, r.Name, r.Affiliations, r.Identities)

	return resp, err
}

// GetAllAffiliations returns all affiliations that the caller is authorized to see
func (c *Client) GetAllAffiliations(options ...RequestOption) (*AffiliationResponse, error) {
	// Read request options
	opts, err := c.prepareOptsFromOptions(c.ctx, options...)
	if err != nil {
		return nil, err
	}

	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	r, err := ca.GetAllAffiliations(opts.CA)
	if err != nil {
		return nil, err
	}

	resp := &AffiliationResponse{CAName: r.CAName, AffiliationInfo: AffiliationInfo{}}
	err = fillAffiliationInfo(&resp.AffiliationInfo, r.Name, r.Affiliations, r.Identities)

	return resp, err
}

// AddAffiliation adds a new affiliation to the server
func (c *Client) AddAffiliation(request *AffiliationRequest) (*AffiliationResponse, error) {
	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	req := &mspapi.AffiliationRequest{
		Name:   request.Name,
		Force:  request.Force,
		CAName: request.CAName,
	}

	r, err := ca.AddAffiliation(req)
	if err != nil {
		return nil, err
	}

	resp := &AffiliationResponse{CAName: r.CAName, AffiliationInfo: AffiliationInfo{}}
	err = fillAffiliationInfo(&resp.AffiliationInfo, r.Name, r.Affiliations, r.Identities)

	return resp, err
}

// ModifyAffiliation renames an existing affiliation on the server
func (c *Client) ModifyAffiliation(request *ModifyAffiliationRequest) (*AffiliationResponse, error) {
	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	req := &mspapi.ModifyAffiliationRequest{
		NewName: request.NewName,
		AffiliationRequest: mspapi.AffiliationRequest{
			Name:   request.Name,
			Force:  request.Force,
			CAName: request.CAName,
		},
	}

	r, err := ca.ModifyAffiliation(req)
	if err != nil {
		return nil, err
	}

	resp := &AffiliationResponse{CAName: r.CAName, AffiliationInfo: AffiliationInfo{}}
	err = fillAffiliationInfo(&resp.AffiliationInfo, r.Name, r.Affiliations, r.Identities)

	return resp, err
}

// RemoveAffiliation removes an existing affiliation from the server
func (c *Client) RemoveAffiliation(request *AffiliationRequest) (*AffiliationResponse, error) {
	ca, err := newCAClient(c.ctx, c.orgName)
	if err != nil {
		return nil, err
	}

	req := &mspapi.AffiliationRequest{
		Name:   request.Name,
		Force:  request.Force,
		CAName: request.CAName,
	}

	r, err := ca.RemoveAffiliation(req)
	if err != nil {
		return nil, err
	}

	resp := &AffiliationResponse{CAName: r.CAName, AffiliationInfo: AffiliationInfo{}}
	err = fillAffiliationInfo(&resp.AffiliationInfo, r.Name, r.Affiliations, r.Identities)

	return resp, err
}

func fillAffiliationInfo(info *AffiliationInfo, name string, affiliations []mspapi.AffiliationInfo, identities []mspapi.IdentityInfo) error {
	info.Name = name

	// Add identities which have this affiliation
	idents := []IdentityInfo{}
	for _, identity := range identities {
		idents = append(idents, IdentityInfo{ID: identity.ID, Type: identity.Type, Affiliation: identity.Affiliation, Attributes: getAllAttributes(identity.Attributes), MaxEnrollments: identity.MaxEnrollments})
	}
	if len(idents) > 0 {
		info.Identities = idents
	}

	// Create child affiliations (if any)
	children := []AffiliationInfo{}
	for _, aff := range affiliations {
		childAff := AffiliationInfo{Name: aff.Name}
		err := fillAffiliationInfo(&childAff, aff.Name, aff.Affiliations, aff.Identities)
		if err != nil {
			return err
		}
		children = append(children, childAff)
	}
	if len(children) > 0 {
		info.Affiliations = children
	}
	return nil
}

func getAllAttributes(attrs []mspapi.Attribute) []Attribute {
	attriburtes := []Attribute{}
	for _, attr := range attrs {
		attriburtes = append(attriburtes, Attribute{Name: attr.Name, Value: attr.Value, ECert: attr.ECert})
	}

	return attriburtes
}
