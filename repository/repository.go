package httpclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// UpdateProcurersMaterialsInventory by itemDnaNbr
func UpdateProcurersMaterialsInventory(itemDnaNbr string, quantityavailable int) (success bool, err error) {
	manufacturermaterial, err := getmanufacturermaterialbyitemdna(itemDnaNbr)

	reqcommnad := updateProcurersMaterialsInventoryCommand{}
	reqcommnad.IsCompleteInventory = true
	reqcommnad.Inventory = []inventory{inventory{Handle: manufacturermaterial.ManufacturerMaterial.Handle, QuantityAvailable: quantityavailable, UnitOfMeasure: "EA"}}

	conf := getconfg()
	conf.GetManufacturerMaterialEndpoint = strings.Replace(conf.UpdateProcurersMaterialsInventoryEndpoint, "{ProcurerKey}", manufacturermaterial.ProcurerMaterialReference.ProcurerKey, -1)

	req, _ := http.NewRequest("PUT", conf.UpdateProcurersMaterialsInventoryEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	httpres, httperr := client.Do(req)
	if httperr != nil {
		err = httperr
	}
	success = httpres.StatusCode == 200
	return
}

func getmanufacturermaterialbyitemdna(itemDnaNbr string) (manufacturermaterial material, err error) {
	conf := getconfg()
	conf.GetManufacturerMaterialEndpoint = strings.Replace(conf.GetManufacturerMaterialEndpoint, "{Handle}", itemDnaNbr, -1)

	req, _ := http.NewRequest("GET", conf.GetManufacturerMaterialEndpoint, nil)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	httpres, httperr := client.Do(req)
	if httperr != nil {
		err = httperr
	} else if httpres.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(httpres.Body)
		json.Unmarshal(bodyBytes, &manufacturermaterial)
		return
	}

	return
}

func getconfg() *config {
	// Get config
	file, _ := os.Open("config.test.json")
	decoder := json.NewDecoder(file)
	conf := config{}
	decoder.Decode(&conf)
	return &conf
}

type config struct {
	GetManufacturerMaterialEndpoint           string `json:"IBK-endpoint-getprocurer"`
	UpdateProcurersMaterialsInventoryEndpoint string `json:"IBK-endpoint-updateprocurerinventory"`
}

type material struct {
	ManufacturerMaterial      manufacturermaterial
	ProcurerMaterialReference procurermaterialreference
}

type manufacturermaterial struct {
	ID                  string `json:"Id"`
	Handle              string
	Name                string
	UnitOfMeasure       string
	ProcurerMaterialRef string
	Modified            string
}

type procurermaterialreference struct {
	ProcurerKey            string
	ProcurerMaterialHandle string
}

type updateProcurersMaterialsInventoryCommand struct {
	IsCompleteInventory bool
	Inventory           []inventory
}

type inventory struct {
	Handle                string
	QuantityAvailable     int
	PlannedAvailabilities []plannedAvailabilities
	UnitOfMeasure         string
}

type plannedAvailabilities struct {
}
