package service

import (
	"fmt"
	"strings"
	"time"

	WarpNetClient "github.com/thealonemusk/WarpNet/api/client"
	"github.com/thealonemusk/WarpNet/pkg/protocol"
)

// Client is a wrapper of an WarpNet client
// with additional metadata and syntax sugar
type Client struct {
	serviceID string
	*WarpNetClient.Client
}

// NewClient returns a new client with an associated service ID
func NewClient(serviceID string, c *WarpNetClient.Client) *Client {
	return &Client{serviceID: serviceID, Client: c}
}

// ListItems returns list of items associated with the serviceID and the given suffix
func (c Client) ListItems(serviceID, suffix string) (strs []string, err error) {
	buckets, err := c.Client.GetBucketKeys(serviceID)
	if err != nil {
		return
	}
	for _, b := range buckets {
		if strings.HasSuffix(b, suffix) {
			b = strings.ReplaceAll(b, "-"+suffix, "")
			strs = append(strs, b)
		}
	}
	return
}

type advertizeMessage struct {
	Time time.Time
}

// Advertize advertize the given uuid to the ledger
func (c Client) Advertize(uuid string) error {
	return c.Client.Put(c.serviceID, fmt.Sprintf("%s-uuid", uuid), advertizeMessage{Time: time.Now().UTC()})
}

// AdvertizingNodes returns a list of advertizing nodes
func (c Client) AdvertizingNodes() (active []string, err error) {
	uuids, err := c.ListItems(c.serviceID, "uuid")
	if err != nil {
		return
	}
	for _, u := range uuids {
		var d advertizeMessage
		res, err := c.Client.GetBucketKey(c.serviceID, fmt.Sprintf("%s-uuid", u))
		if err != nil {
			continue
		}
		res.Unmarshal(&d)

		if d.Time.Add(15 * time.Minute).After(time.Now().UTC()) {
			active = append(active, u)
		}
	}
	return
}

// ActiveNodes returns a list of active nodes
func (c Client) ActiveNodes() (active []string, err error) {
	res, err := c.Client.GetBucket(protocol.HealthCheckKey)
	if err != nil {
		return []string{}, err
	}

	for u, r := range res {
		var s string
		r.Unmarshal(&s)
		parsed, _ := time.Parse(time.RFC3339, s)
		if parsed.Add(15 * time.Minute).After(time.Now().UTC()) {
			active = append(active, u)
		}
	}
	return
}

// Clean cleans up the serviceID associated data
func (c Client) Clean() error {
	return c.Client.DeleteBucket(c.serviceID)
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

// Get returns generic data from the API
// e.g. get("ip", uuid)
func (c Client) Get(args ...string) (string, error) {
	reverse(args)
	key := strings.Join(args, "-")
	var role string
	d, err := c.Client.GetBucketKey(c.serviceID, key)
	if err == nil {
		d.Unmarshal(&role)
	}
	return role, err
}

// Set generic data to the API
// e.g. set("ip", uuid, "value")
func (c Client) Set(thing, uuid, value string) error {
	return c.Client.Put(c.serviceID, fmt.Sprintf("%s-%s", uuid, thing), value)
}
