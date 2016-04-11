// contivModelClient.go
// This file is auto generated by modelgen tool
// Do not edit this file manually

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// Link is a one way relattion between two objects
type Link struct {
	ObjType string `json:"type,omitempty"`
	ObjKey  string `json:"key,omitempty"`
}

func httpGet(url string, jdata interface{}) error {

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	switch {
	case r.StatusCode == int(404):
		return errors.New("Page not found!")
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode == int(500):
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return errors.New(string(response))

	case r.StatusCode != int(200):
		log.Debugf("GET Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(response, jdata); err != nil {
		return err
	}

	return nil
}

func httpDelete(url string) error {

	req, err := http.NewRequest("DELETE", url, nil)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// body, _ := ioutil.ReadAll(r.Body)

	switch {
	case r.StatusCode == int(404):
		// return errors.New("Page not found!")
		return nil
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode == int(500):
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return errors.New(string(response))

	case r.StatusCode != int(200):
		log.Debugf("DELETE Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	return nil
}

func httpPost(url string, jdata interface{}) error {
	buf, err := json.Marshal(jdata)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(buf)
	r, err := http.Post(url, "application/json", body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	switch {
	case r.StatusCode == int(404):
		return errors.New("Page not found!")
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode == int(500):
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		return errors.New(string(response))

	case r.StatusCode != int(200):
		log.Debugf("POST Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Debugf(string(response))

	return nil
}

// ContivClient has the contiv model client instance
type ContivClient struct {
	baseURL string
}

// NewContivClient returns a new client instance
func NewContivClient(baseURL string) (*ContivClient, error) {
	client := ContivClient{
		baseURL: baseURL,
	}

	return &client, nil
}

type AppProfile struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppProfileName string   `json:"appProfileName,omitempty"` // Application Profile Name
	EndpointGroups []string `json:"endpointGroups,omitempty"`
	NetworkName    string   `json:"networkName,omitempty"` // Network of App Prof
	TenantName     string   `json:"tenantName,omitempty"`  // Tenant Name

	// add link-sets and links
	LinkSets AppProfileLinkSets `json:"link-sets,omitempty"`
	Links    AppProfileLinks    `json:"links,omitempty"`
}

type AppProfileLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
}

type AppProfileLinks struct {
	Network Link `json:"Network,omitempty"`
	Tenant  Link `json:"Tenant,omitempty"`
}

type EndpointGroup struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	EndpointGroupID int      `json:"endpointGroupId,omitempty"` // Group Identifier
	GroupName       string   `json:"groupName,omitempty"`       // Group name
	NetworkName     string   `json:"networkName,omitempty"`     // Network
	Policies        []string `json:"policies,omitempty"`
	TenantName      string   `json:"tenantName,omitempty"` // Tenant

	// add link-sets and links
	LinkSets EndpointGroupLinkSets `json:"link-sets,omitempty"`
	Links    EndpointGroupLinks    `json:"links,omitempty"`
}

type EndpointGroupLinkSets struct {
	Policies map[string]Link `json:"Policies,omitempty"`
	Services map[string]Link `json:"Services,omitempty"`
}

type EndpointGroupLinks struct {
	AppProfile Link `json:"AppProfile,omitempty"`
	Network    Link `json:"Network,omitempty"`
	Tenant     Link `json:"Tenant,omitempty"`
}

type Global struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Name             string `json:"name,omitempty"`               // name of this block
	NetworkInfraType string `json:"network-infra-type,omitempty"` // Network infrastructure type
	Vlans            string `json:"vlans,omitempty"`              // Allowed vlan range
	Vxlans           string `json:"vxlans,omitempty"`             // Allwed vxlan range

}

type Bgp struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	As         string `json:"as,omitempty"`          // AS id
	Hostname   string `json:"hostname,omitempty"`    // host name
	Neighbor   string `json:"neighbor,omitempty"`    // Bgp  neighbor
	NeighborAs string `json:"neighbor-as,omitempty"` // AS id
	Routerip   string `json:"routerip,omitempty"`    // Bgp router intf ip

}

type Network struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Encap       string `json:"encap,omitempty"`       // Encapsulation
	Gateway     string `json:"gateway,omitempty"`     // Gateway
	NetworkName string `json:"networkName,omitempty"` // Network name
	PktTag      int    `json:"pktTag,omitempty"`      // Vlan/Vxlan Tag
	Subnet      string `json:"subnet,omitempty"`      // Subnet
	TenantName  string `json:"tenantName,omitempty"`  // Tenant Name

	// add link-sets and links
	LinkSets NetworkLinkSets `json:"link-sets,omitempty"`
	Links    NetworkLinks    `json:"links,omitempty"`
}

type NetworkLinkSets struct {
	AppProfiles    map[string]Link `json:"AppProfiles,omitempty"`
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Services       map[string]Link `json:"Services,omitempty"`
}

type NetworkLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type Policy struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	PolicyName string `json:"policyName,omitempty"` // Policy Name
	TenantName string `json:"tenantName,omitempty"` // Tenant Name

	// add link-sets and links
	LinkSets PolicyLinkSets `json:"link-sets,omitempty"`
	Links    PolicyLinks    `json:"links,omitempty"`
}

type PolicyLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Rules          map[string]Link `json:"Rules,omitempty"`
}

type PolicyLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type Rule struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Action            string `json:"action,omitempty"`            // Action
	Direction         string `json:"direction,omitempty"`         // Direction
	FromEndpointGroup string `json:"fromEndpointGroup,omitempty"` // From Endpoint Group
	FromIpAddress     string `json:"fromIpAddress,omitempty"`     // IP Address
	FromNetwork       string `json:"fromNetwork,omitempty"`       // From Network
	PolicyName        string `json:"policyName,omitempty"`        // Policy Name
	Port              int    `json:"port,omitempty"`              // Port No
	Priority          int    `json:"priority,omitempty"`          // Priority
	Protocol          string `json:"protocol,omitempty"`          // Protocol
	RuleID            string `json:"ruleId,omitempty"`            // Rule Id
	TenantName        string `json:"tenantName,omitempty"`        // Tenant Name
	ToEndpointGroup   string `json:"toEndpointGroup,omitempty"`   // To Endpoint Group
	ToIpAddress       string `json:"toIpAddress,omitempty"`       // IP Address
	ToNetwork         string `json:"toNetwork,omitempty"`         // To Network

	// add link-sets and links
	LinkSets RuleLinkSets `json:"link-sets,omitempty"`
}

type RuleLinkSets struct {
	Policies map[string]Link `json:"Policies,omitempty"`
}

type Service struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppName        string   `json:"appName,omitempty"` // Application Name
	Command        string   `json:"command,omitempty"` //
	Cpu            string   `json:"cpu,omitempty"`     //
	EndpointGroups []string `json:"endpointGroups,omitempty"`
	Environment    []string `json:"environment,omitempty"`
	ImageName      string   `json:"imageName,omitempty"` //
	Memory         string   `json:"memory,omitempty"`    //
	Networks       []string `json:"networks,omitempty"`
	Scale          int      `json:"scale,omitempty"`         //
	ServiceName    string   `json:"serviceName,omitempty"`   // Service Name
	TenantName     string   `json:"tenantName,omitempty"`    // Tenant Name
	VolumeProfile  string   `json:"volumeProfile,omitempty"` //

	// add link-sets and links
	LinkSets ServiceLinkSets `json:"link-sets,omitempty"`
	Links    ServiceLinks    `json:"links,omitempty"`
}

type ServiceLinkSets struct {
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Instances      map[string]Link `json:"Instances,omitempty"`
	Networks       map[string]Link `json:"Networks,omitempty"`
}

type ServiceLinks struct {
	App           Link `json:"App,omitempty"`
	VolumeProfile Link `json:"VolumeProfile,omitempty"`
}

type ServiceInstance struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppName     string   `json:"appName,omitempty"`     //
	InstanceID  string   `json:"instanceId,omitempty"`  // Service instance id
	ServiceName string   `json:"serviceName,omitempty"` //
	TenantName  string   `json:"tenantName,omitempty"`  // Tenant Name
	Volumes     []string `json:"volumes,omitempty"`

	// add link-sets and links
	LinkSets ServiceInstanceLinkSets `json:"link-sets,omitempty"`
	Links    ServiceInstanceLinks    `json:"links,omitempty"`
}

type ServiceInstanceLinkSets struct {
	Volumes map[string]Link `json:"Volumes,omitempty"`
}

type ServiceInstanceLinks struct {
	Service Link `json:"Service,omitempty"`
}

type Tenant struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DefaultNetwork string `json:"defaultNetwork,omitempty"` // Network name
	TenantName     string `json:"tenantName,omitempty"`     // Tenant Name

	// add link-sets and links
	LinkSets TenantLinkSets `json:"link-sets,omitempty"`
}

type TenantLinkSets struct {
	AppProfiles    map[string]Link `json:"AppProfiles,omitempty"`
	EndpointGroups map[string]Link `json:"EndpointGroups,omitempty"`
	Networks       map[string]Link `json:"Networks,omitempty"`
	Policies       map[string]Link `json:"Policies,omitempty"`
	VolumeProfiles map[string]Link `json:"VolumeProfiles,omitempty"`
	Volumes        map[string]Link `json:"Volumes,omitempty"`
}

type Volume struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DatastoreType string `json:"datastoreType,omitempty"` //
	MountPoint    string `json:"mountPoint,omitempty"`    //
	PoolName      string `json:"poolName,omitempty"`      //
	Size          string `json:"size,omitempty"`          //
	TenantName    string `json:"tenantName,omitempty"`    // Tenant Name
	VolumeName    string `json:"volumeName,omitempty"`    // Volume Name

	// add link-sets and links
	LinkSets VolumeLinkSets `json:"link-sets,omitempty"`
	Links    VolumeLinks    `json:"links,omitempty"`
}

type VolumeLinkSets struct {
	ServiceInstances map[string]Link `json:"ServiceInstances,omitempty"`
}

type VolumeLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

type VolumeProfile struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DatastoreType     string `json:"datastoreType,omitempty"`     //
	MountPoint        string `json:"mountPoint,omitempty"`        //
	PoolName          string `json:"poolName,omitempty"`          //
	Size              string `json:"size,omitempty"`              //
	TenantName        string `json:"tenantName,omitempty"`        // Tenant Name
	VolumeProfileName string `json:"volumeProfileName,omitempty"` // Volume profile Name

	// add link-sets and links
	LinkSets VolumeProfileLinkSets `json:"link-sets,omitempty"`
	Links    VolumeProfileLinks    `json:"links,omitempty"`
}

type VolumeProfileLinkSets struct {
	Services map[string]Link `json:"Services,omitempty"`
}

type VolumeProfileLinks struct {
	Tenant Link `json:"Tenant,omitempty"`
}

// AppProfilePost posts the appProfile object
func (c *ContivClient) AppProfilePost(obj *AppProfile) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.NetworkName + ":" + obj.AppProfileName
	url := c.baseURL + "/api/appProfiles/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating appProfile %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// AppProfileList lists all appProfile objects
func (c *ContivClient) AppProfileList() (*[]*AppProfile, error) {
	// build key and URL
	url := c.baseURL + "/api/appProfiles/"

	// http get the object
	var objList []*AppProfile
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting appProfiles. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// AppProfileGet gets the appProfile object
func (c *ContivClient) AppProfileGet(tenantName string, networkName string, appProfileName string) (*AppProfile, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName + ":" + appProfileName
	url := c.baseURL + "/api/appProfiles/" + keyStr + "/"

	// http get the object
	var obj AppProfile
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting appProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// AppProfileDelete deletes the appProfile object
func (c *ContivClient) AppProfileDelete(tenantName string, networkName string, appProfileName string) error {
	// build key and URL
	keyStr := tenantName + ":" + networkName + ":" + appProfileName
	url := c.baseURL + "/api/appProfiles/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting appProfile %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// EndpointGroupPost posts the endpointGroup object
func (c *ContivClient) EndpointGroupPost(obj *EndpointGroup) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.NetworkName + ":" + obj.GroupName
	url := c.baseURL + "/api/endpointGroups/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating endpointGroup %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// EndpointGroupList lists all endpointGroup objects
func (c *ContivClient) EndpointGroupList() (*[]*EndpointGroup, error) {
	// build key and URL
	url := c.baseURL + "/api/endpointGroups/"

	// http get the object
	var objList []*EndpointGroup
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting endpointGroups. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// EndpointGroupGet gets the endpointGroup object
func (c *ContivClient) EndpointGroupGet(tenantName string, networkName string, groupName string) (*EndpointGroup, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName + ":" + groupName
	url := c.baseURL + "/api/endpointGroups/" + keyStr + "/"

	// http get the object
	var obj EndpointGroup
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting endpointGroup %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// EndpointGroupDelete deletes the endpointGroup object
func (c *ContivClient) EndpointGroupDelete(tenantName string, networkName string, groupName string) error {
	// build key and URL
	keyStr := tenantName + ":" + networkName + ":" + groupName
	url := c.baseURL + "/api/endpointGroups/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting endpointGroup %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// GlobalPost posts the global object
func (c *ContivClient) GlobalPost(obj *Global) error {
	// build key and URL
	keyStr := obj.Name
	url := c.baseURL + "/api/globals/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating global %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// GlobalList lists all global objects
func (c *ContivClient) GlobalList() (*[]*Global, error) {
	// build key and URL
	url := c.baseURL + "/api/globals/"

	// http get the object
	var objList []*Global
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting globals. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// GlobalGet gets the global object
func (c *ContivClient) GlobalGet(name string) (*Global, error) {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/globals/" + keyStr + "/"

	// http get the object
	var obj Global
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting global %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// GlobalDelete deletes the global object
func (c *ContivClient) GlobalDelete(name string) error {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/globals/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting global %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// BgpPost posts the Bgp object
func (c *ContivClient) BgpPost(obj *Bgp) error {
	// build key and URL
	keyStr := obj.Hostname
	url := c.baseURL + "/api/Bgps/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating Bgp %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// BgpList lists all Bgp objects
func (c *ContivClient) BgpList() (*[]*Bgp, error) {
	// build key and URL
	url := c.baseURL + "/api/Bgps/"

	// http get the object
	var objList []*Bgp
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting Bgps. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// BgpGet gets the Bgp object
func (c *ContivClient) BgpGet(hostname string) (*Bgp, error) {
	// build key and URL
	keyStr := hostname
	url := c.baseURL + "/api/Bgps/" + keyStr + "/"

	// http get the object
	var obj Bgp
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting Bgp %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// BgpDelete deletes the Bgp object
func (c *ContivClient) BgpDelete(hostname string) error {
	// build key and URL
	keyStr := hostname
	url := c.baseURL + "/api/Bgps/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting Bgp %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// NetworkPost posts the network object
func (c *ContivClient) NetworkPost(obj *Network) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.NetworkName
	url := c.baseURL + "/api/networks/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating network %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// NetworkList lists all network objects
func (c *ContivClient) NetworkList() (*[]*Network, error) {
	// build key and URL
	url := c.baseURL + "/api/networks/"

	// http get the object
	var objList []*Network
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting networks. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// NetworkGet gets the network object
func (c *ContivClient) NetworkGet(tenantName string, networkName string) (*Network, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/networks/" + keyStr + "/"

	// http get the object
	var obj Network
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting network %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// NetworkDelete deletes the network object
func (c *ContivClient) NetworkDelete(tenantName string, networkName string) error {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/networks/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting network %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// PolicyPost posts the policy object
func (c *ContivClient) PolicyPost(obj *Policy) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.PolicyName
	url := c.baseURL + "/api/policys/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating policy %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// PolicyList lists all policy objects
func (c *ContivClient) PolicyList() (*[]*Policy, error) {
	// build key and URL
	url := c.baseURL + "/api/policys/"

	// http get the object
	var objList []*Policy
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting policys. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// PolicyGet gets the policy object
func (c *ContivClient) PolicyGet(tenantName string, policyName string) (*Policy, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/policys/" + keyStr + "/"

	// http get the object
	var obj Policy
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting policy %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// PolicyDelete deletes the policy object
func (c *ContivClient) PolicyDelete(tenantName string, policyName string) error {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/policys/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting policy %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// RulePost posts the rule object
func (c *ContivClient) RulePost(obj *Rule) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.PolicyName + ":" + obj.RuleID
	url := c.baseURL + "/api/rules/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating rule %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// RuleList lists all rule objects
func (c *ContivClient) RuleList() (*[]*Rule, error) {
	// build key and URL
	url := c.baseURL + "/api/rules/"

	// http get the object
	var objList []*Rule
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting rules. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// RuleGet gets the rule object
func (c *ContivClient) RuleGet(tenantName string, policyName string, ruleId string) (*Rule, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/rules/" + keyStr + "/"

	// http get the object
	var obj Rule
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting rule %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// RuleDelete deletes the rule object
func (c *ContivClient) RuleDelete(tenantName string, policyName string, ruleId string) error {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/rules/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting rule %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// ServicePost posts the service object
func (c *ContivClient) ServicePost(obj *Service) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.AppName + ":" + obj.ServiceName
	url := c.baseURL + "/api/services/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating service %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// ServiceList lists all service objects
func (c *ContivClient) ServiceList() (*[]*Service, error) {
	// build key and URL
	url := c.baseURL + "/api/services/"

	// http get the object
	var objList []*Service
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting services. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// ServiceGet gets the service object
func (c *ContivClient) ServiceGet(tenantName string, appName string, serviceName string) (*Service, error) {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName
	url := c.baseURL + "/api/services/" + keyStr + "/"

	// http get the object
	var obj Service
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting service %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ServiceDelete deletes the service object
func (c *ContivClient) ServiceDelete(tenantName string, appName string, serviceName string) error {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName
	url := c.baseURL + "/api/services/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting service %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// ServiceInstancePost posts the serviceInstance object
func (c *ContivClient) ServiceInstancePost(obj *ServiceInstance) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.AppName + ":" + obj.ServiceName + ":" + obj.InstanceID
	url := c.baseURL + "/api/serviceInstances/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating serviceInstance %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// ServiceInstanceList lists all serviceInstance objects
func (c *ContivClient) ServiceInstanceList() (*[]*ServiceInstance, error) {
	// build key and URL
	url := c.baseURL + "/api/serviceInstances/"

	// http get the object
	var objList []*ServiceInstance
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting serviceInstances. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// ServiceInstanceGet gets the serviceInstance object
func (c *ContivClient) ServiceInstanceGet(tenantName string, appName string, serviceName string, instanceId string) (*ServiceInstance, error) {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName + ":" + instanceId
	url := c.baseURL + "/api/serviceInstances/" + keyStr + "/"

	// http get the object
	var obj ServiceInstance
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting serviceInstance %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ServiceInstanceDelete deletes the serviceInstance object
func (c *ContivClient) ServiceInstanceDelete(tenantName string, appName string, serviceName string, instanceId string) error {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName + ":" + instanceId
	url := c.baseURL + "/api/serviceInstances/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting serviceInstance %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// TenantPost posts the tenant object
func (c *ContivClient) TenantPost(obj *Tenant) error {
	// build key and URL
	keyStr := obj.TenantName
	url := c.baseURL + "/api/tenants/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating tenant %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// TenantList lists all tenant objects
func (c *ContivClient) TenantList() (*[]*Tenant, error) {
	// build key and URL
	url := c.baseURL + "/api/tenants/"

	// http get the object
	var objList []*Tenant
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting tenants. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// TenantGet gets the tenant object
func (c *ContivClient) TenantGet(tenantName string) (*Tenant, error) {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/tenants/" + keyStr + "/"

	// http get the object
	var obj Tenant
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting tenant %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// TenantDelete deletes the tenant object
func (c *ContivClient) TenantDelete(tenantName string) error {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/tenants/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting tenant %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// VolumePost posts the volume object
func (c *ContivClient) VolumePost(obj *Volume) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.VolumeName
	url := c.baseURL + "/api/volumes/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating volume %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// VolumeList lists all volume objects
func (c *ContivClient) VolumeList() (*[]*Volume, error) {
	// build key and URL
	url := c.baseURL + "/api/volumes/"

	// http get the object
	var objList []*Volume
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting volumes. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// VolumeGet gets the volume object
func (c *ContivClient) VolumeGet(tenantName string, volumeName string) (*Volume, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/volumes/" + keyStr + "/"

	// http get the object
	var obj Volume
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting volume %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeDelete deletes the volume object
func (c *ContivClient) VolumeDelete(tenantName string, volumeName string) error {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/volumes/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting volume %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// VolumeProfilePost posts the volumeProfile object
func (c *ContivClient) VolumeProfilePost(obj *VolumeProfile) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.VolumeProfileName
	url := c.baseURL + "/api/volumeProfiles/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Debugf("Error creating volumeProfile %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// VolumeProfileList lists all volumeProfile objects
func (c *ContivClient) VolumeProfileList() (*[]*VolumeProfile, error) {
	// build key and URL
	url := c.baseURL + "/api/volumeProfiles/"

	// http get the object
	var objList []*VolumeProfile
	err := httpGet(url, &objList)
	if err != nil {
		log.Debugf("Error getting volumeProfiles. Err: %v", err)
		return nil, err
	}

	return &objList, nil
}

// VolumeProfileGet gets the volumeProfile object
func (c *ContivClient) VolumeProfileGet(tenantName string, volumeProfileName string) (*VolumeProfile, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/volumeProfiles/" + keyStr + "/"

	// http get the object
	var obj VolumeProfile
	err := httpGet(url, &obj)
	if err != nil {
		log.Debugf("Error getting volumeProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeProfileDelete deletes the volumeProfile object
func (c *ContivClient) VolumeProfileDelete(tenantName string, volumeProfileName string) error {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/volumeProfiles/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Debugf("Error deleting volumeProfile %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}
