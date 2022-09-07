package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-account-and-individual-customer-master-data-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetAccountAndIndividualCustomerMaster(contactID, accountID, employeeID, customerID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ContactCollection":
			func() {
				c.ContactCollection(contactID)
				wg.Done()
			}()
		case "CorporateAccountCollection":
			func() {
				c.CorporateAccount(accountID)
				wg.Done()
			}()
		case "EmployeeBasicData":
			func() {
				c.EmployeeBasicData(employeeID)
				wg.Done()
			}()
		case "IndividualCustomerCollection":
			func() {
				c.IndividualCustomerCollection(customerID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ContactCollection(contactID string) {
	contactCollectionData, err := c.callAccountAndIndividualCustomerMasterDataSrvAPIRequirementContactCollection("ContactCollection", contactID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactCollectionData)

	contactIsContactPersonForData, err := c.callContactIsContactPersonFor(contactCollectionData[0].ToContactIsContactPersonFor)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(contactIsContactPersonForData)

	corporateAccountData, err := c.callCorporateAccount(contactCollectionData[0].ToCorporateAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(corporateAccountData)
}

func (c *SAPAPICaller) callAccountAndIndividualCustomerMasterDataSrvAPIRequirementContactCollection(api, contactID string) ([]sap_api_output_formatter.ContactCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithContactCollection(req, contactID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callContactIsContactPersonFor(url string) ([]sap_api_output_formatter.ContactIsContactPersonFor, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToContactIsContactPersonFor(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCorporateAccount(url string) (*sap_api_output_formatter.CorporateAccount, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCorporateAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CorporateAccount(accountID string) {
	corporateAccountData, err := c.callAccountAndIndividualCustomerMasterSrvAPIRequirementCorporateAccount("CorporateAccount", accountID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(corporateAccountData)
}

func (c *SAPAPICaller) callAccountAndIndividualCustomerMasterSrvAPIRequirementCorporateAccount(api, accountID string) (*sap_api_output_formatter.CorporateAccount, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithCorporateAccount(req, accountID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCorporateAccount(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) EmployeeBasicData(employeeID string) {
	employeeBasicData, err := c.callAccountAndIndividualCustomerMasterSrvAPIRequirementEmployeeBasicData("EmployeeBasicData", employeeID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeBasicData)
}

func (c *SAPAPICaller) callAccountAndIndividualCustomerMasterSrvAPIRequirementEmployeeBasicData(api, employeeID string) ([]sap_api_output_formatter.EmployeeBasicData, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEmployeeBasicData(req, employeeID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeBasicData(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) IndividualCustomerCollection(customerID string) {
	individualCustomerCollection, err := c.callAccountAndIndividualCustomerMasterSrvAPIRequirementIndividualCustomerCollection("IndividualCustomerCollection", customerID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(individualCustomerCollection)
}

func (c *SAPAPICaller) callAccountAndIndividualCustomerMasterSrvAPIRequirementIndividualCustomerCollection(api, customerID string) ([]sap_api_output_formatter.IndividualCustomerCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithIndividualCustomerCollection(req, customerID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToIndividualCustomerCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithContactCollection(req *http.Request, contactID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ContactID eq '%s'", contactID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithCorporateAccount(req *http.Request, accountID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("AccountID eq '%s'", accountID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEmployeeBasicData(req *http.Request, employeeID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("EmployeeID eq '%s'", employeeID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithIndividualCustomerCollection(req *http.Request, customerID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("CustomerID eq '%s'", customerID))
	req.URL.RawQuery = params.Encode()
}
