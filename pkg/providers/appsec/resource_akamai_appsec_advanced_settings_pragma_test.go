package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
)

func TestAccAkamaiAdvancedSettingsPragma_res_basic(t *testing.T) {
	t.Run("match by AdvancedSettingsPragma ID", func(t *testing.T) {
		client := &mockappsec{}

		cu := appsec.UpdateAdvancedSettingsPragmaResponse{}
		expectJSU := compactJSON(loadFixtureBytes("testdata/TestResAdvancedSettingsPragma/AdvancedSettingsPragma.json"))
		json.Unmarshal([]byte(expectJSU), &cu)

		cd := appsec.UpdateAdvancedSettingsPragmaResponse{}
		expectJSD := compactJSON(loadFixtureBytes("testdata/TestResAdvancedSettingsPragma/AdvancedSettingsPragma.json"))
		json.Unmarshal([]byte(expectJSD), &cd)

		cr := appsec.GetAdvancedSettingsPragmaResponse{}
		expectJS := compactJSON(loadFixtureBytes("testdata/TestResAdvancedSettingsPragma/AdvancedSettingsPragma.json"))
		json.Unmarshal([]byte(expectJS), &cr)

		config := appsec.GetConfigurationResponse{}
		expectConfigs := compactJSON(loadFixtureBytes("testdata/TestResConfiguration/LatestConfiguration.json"))
		json.Unmarshal([]byte(expectConfigs), &config)

		client.On("GetConfiguration",
			mock.Anything,
			appsec.GetConfigurationRequest{ConfigID: 43253},
		).Return(&config, nil)

		client.On("GetAdvancedSettingsPragma",
			mock.Anything, // ctx is irrelevant for this test
			appsec.GetAdvancedSettingsPragmaRequest{ConfigID: 43253, Version: 7},
		).Return(&cr, nil)

		client.On("UpdateAdvancedSettingsPragma",
			mock.Anything, // ctx is irrelevant for this test
			appsec.UpdateAdvancedSettingsPragmaRequest{ConfigID: 43253, Version: 7, PolicyID: "", JsonPayloadRaw: json.RawMessage("{\"action\":\"REMOVE\"}\n")},
		).Return(&cu, nil)

		client.On("UpdateAdvancedSettingsPragma",
			mock.Anything, // ctx is irrelevant for this test
			appsec.UpdateAdvancedSettingsPragmaRequest{ConfigID: 43253, Version: 7, PolicyID: "", JsonPayloadRaw: json.RawMessage("{}")},
		).Return(&cu, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestResAdvancedSettingsPragma/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_appsec_advanced_settings_pragma_header.test", "id", "43253"),
						),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
