package utilization

import (
	"encoding/json"

	"github.com/HewlettPackard/oneview-golang/rest"
	"github.com/HewlettPackard/oneview-golang/utils"
	"github.com/docker/machine/libmachine/log"
)

//Enclosure and ServerHardware struct
type Utilization struct {
	RefreshTaskURI   utils.Nstring `json:"refreshTaskUri,omitempty"`   //"refreshTaskUri": null,
	URI              utils.Nstring `json:"uri,omitempty"`              //"uri": "/rest/enclosures/797740CZJ809076M"
	Resolution       int           `json:"resolution,omitempty"`       //"resolution": 3000000,
	SliceEndTime     string        `json:"sliceEndTime,omitempty"`     //"sliceEndTime": "2019-05-01T11:10:00.000Z",
	SliceStartTime   string        `json:"sliceStartTime,omitempty"`   //"sliceStartTime": "2019-05-01T11:10:00.000Z",
	NewestSampleTime string        `json:"newestSampleTime,omitempty"` //"newestSampleTime": "2019-05-01T11:10:00.000Z",
	OldestSampleTime string        `json:"oldestSampleTime,omitempty"` //"oldestSampleTime": "2019-05-01T11:10:00.000Z",
	IsFresh          bool          `json:"isFresh,omitempty"`          // "total": false,
	MetricList       []MetricList  `json:"metricList,omitempty"`       // "metricList":[]
}

//MetricList - EnclosureUtilization
type MetricList struct {
	MetricName     string          `json:"metricName,omitempty"`     //"metricName": "PeakPower",
	MetricSamples  [][]interface{} `json:"metricSamples,omitempty"`  //"metricSamples":[[1557345600000 911][1557345600000 935]],
	MetricCapacity int             `json:"metricCapacity,omitempty"` // "metricCapacity": 35
}

//InterconnectUtilization struct
type InterconnectUtilization struct {
	URI        utils.Nstring `json:"uri"` //"uri": "/rest/interconnects/797740CZJ809076M"
	MetricList []struct {    // "metricList":[]
		MetricName      string            `json:"metricName"`      //"metricName": "Cpu", / "Memory", / "Temperture"
		MetricSamples   [][][]interface{} `json:"metricSamples"`   //"metricSamples":[[1557345600000 911][1557345600000 935]],
		MetricCapacity  int               `json:"metricCapacity"`  // "metricCapacity": 100
		MetricThreshold [][][]string      `json:"metricThreshold"` // "metricThreshold": "100"
		MetricUnit      string            `json:"metricUnit"`      // "metricUnit": "Percentage", / "GB", / "Fahrenheit"
	} `json:"metricList"`
}

//GetUtilization for Enclosures and ServerHardware
func (c *OVClient) GetUtilization(fields string, filter string, refresh string, view string, uri string) (Utilization, error) {
	var (
		q           map[string]interface{}
		utilization Utilization
		uURI        string
	)

	q = make(map[string]interface{})
	if len(filter) > 0 {
		q["filter"] = filter
	}

	if fields != "" {
		q["fields"] = fields
	}

	if view != "" {
		q["view"] = view
	}

	if refresh != "" {
		q["refresh"] = refresh
	}

	//refresh login
	c.RefreshLogin()
	c.SetAuthHeaderOptions(c.GetAuthHeaderMap())
	//Setup query
	if len(q) > 0 {
		c.SetQueryString(q)
	}

	uURI = uri + "/utilization"
	data, err := c.RestAPICall(rest.GET, uURI, nil)
	if err != nil {
		return utilization, err
	}
	log.Debugf("GetEnclosuresUtilization ", data)
	if err := json.Unmarshal([]byte(data), &utilization); err != nil {
		return utilization, err
	}
	return utilization, err
}
