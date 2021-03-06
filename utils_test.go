package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_copy(t *testing.T) {
	a := &AA{A: 999}
	b := &AA{}
	err := Clone(b, a)
	t.Log(a, b, err)
}

func TestPWD(t *testing.T) {
	t.Log(AalidataPwd("xxxx"))
}

func TestPhone(t *testing.T) {
	t.Log(PhoneRegexp("8601599999999"))
}

func TestVal(t *testing.T) {
	t.Log(PhoneValidate("17799892631"))
}

type AA struct {
	CC
	A int `json:"a"`
}

type BB interface {
	Decode(b *[]byte) error
	Encode() (*[]byte, error)
}

type CC struct{}

func (this *CC) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CC) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

/*
{"code":0,"data":{
	"country":"\u4e2d\u56fd",//国家
	"country_id":"CN",
	"area":"\u534e\u5357",//区域
	"area_id":"800000",
	"region":"\u5e7f\u4e1c\u7701",//所在省
	"region_id":"440000",
	"city":"\u6df1\u5733\u5e02",//所在市
	"city_id":"440300",
	"county":"\u5b9d\u5b89\u533a",//所在县
	"county_id":"440306",
	"isp":"\u7535\u4fe1",
	"isp_id":"100017",
	"ip":"119.137.54.77"}}
*/

func TestIPAddr(t *testing.T) {
	//ip, err := GetInternalIP()
	//t.Log("ip -> ", ip, err)
	//ip, err = GetInternalIP2()
	//t.Log("ip -> ", ip, err)
	//ip, err = GetExternalIP()
	//t.Log("ip -> ", ip, err)
	res, err := GetIPAddrByTaoBao("117.136.80.124")
	t.Logf("res -> %#v, err -> %v", res, err)
}

func TestIPMap(t *testing.T) {
	ip := "117.136.80.124"
	url := "http://api.map.baidu.com/location/ip?ak=F454f8a5efe5e577997931cc01de3974&ip=" + ip + "&coor=bd09ll"
	//url := "http://api.map.baidu.com/location/ip?ak=32f38c9491f2da9eb61106aaab1e9739&ip=" + ip + "&coor=bd09ll"
	//url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	t.Logf("b -> %s, err -> %v", string(b), err)
}

/*
func TestIPList(t *testing.T) {
	i := 0
	for k, v := range ip_list {
		i++
		if i > 10 {
			break
		}
		res, err := GetIPAddrByTaoBao(v)
		if err != nil {
			t.Logf("res -> %#v, err -> %v\n", res, err)
			continue
		}
		t.Logf("id %s, ip %s, Country %s, Region %s, City %s, County %s\n",
			k, v, res.Data.Country, res.Data.Region, res.Data.City, res.Data.County)
	}
}

var ip_list = map[string]string{
	"28352": "119.39.95.224",
	"28351": "223.104.63.56",
	"28350": "183.14.25.43",
	"28349": "183.14.25.43",
	"28348": "119.39.248.81",
	"28347": "113.73.57.249",
	"28346": "119.139.196.49",
	"28345": "171.113.115.87",
	"28344": "223.90.168.20",
	"28343": "183.18.204.210",
	"28342": "175.175.180.103",
	"28341": "106.19.16.174",
	"28340": "61.144.175.36",
	"28339": "113.220.81.13",
	"28338": "183.3.218.5",
	"28337": "106.19.130.66",
	"28336": "117.136.89.71",
	"28335": "118.249.189.130",
	"28334": "118.249.189.130",
	"28333": "118.249.189.130",
	"28372": "223.74.101.48",
	"28371": "119.44.13.92",
	"28370": "223.104.131.213",
	"28369": "117.136.88.217",
	"28368": "113.200.204.212",
	"28367": "14.222.195.161",
	"28366": "14.16.32.50",
	"28365": "112.17.241.133",
	"28364": "117.136.79.50",
	"28363": "17.200.11.44",
	"28362": "218.17.160.168",
	"28361": "183.15.177.231",
	"28360": "119.139.137.104",
	"28359": "183.39.197.188",
	"28358": "117.41.237.233",
	"28357": "106.19.150.41",
	"28356": "59.109.193.79",
	"28355": "223.104.21.69",
	"28354": "110.52.142.134",
	"28353": "121.224.10.213",
	"28392": "27.46.22.46",
	"28391": "219.134.6.187",
	"28390": "27.38.56.31",
	"28389": "120.135.23.121",
	"28388": "222.185.252.146",
	"28387": "117.148.98.227",
	"28386": "121.239.191.206",
	"28385": "14.121.247.215",
	"28384": "113.220.87.56",
	"28383": "117.136.80.113",
	"28382": "117.136.88.73",
	"28381": "113.83.124.213",
	"28380": "223.154.133.245",
	"28379": "106.16.81.137",
	"28378": "220.170.219.211",
	"28377": "220.202.152.97",
	"28376": "106.19.70.167",
	"28375": "183.37.27.180",
	"28374": "223.104.64.49",
	"28373": "211.143.15.174",
	"28412": "113.78.90.203",
	"28411": "60.6.219.172",
	"28410": "117.136.80.221",
	"28409": "183.156.67.189",
	"28408": "106.19.193.251",
	"28407": "123.15.162.155",
	"28406": "117.136.29.180",
	"28405": "43.250.201.77",
	"28404": "118.247.124.201",
	"28403": "183.198.195.45",
	"28402": "119.39.248.69",
	"28401": "219.157.237.65",
	"28400": "113.91.63.209",
	"28399": "113.91.63.209",
	"28398": "27.13.43.227",
	"28397": "112.234.97.51",
	"28396": "117.136.80.124",
	"28395": "117.136.80.188",
	"28394": "117.136.80.21",
	"28393": "183.204.74.242",
}
*/
