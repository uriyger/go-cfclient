package cfclient

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient_ListV3ServicePlansForSpace(t *testing.T) {
	Convey("List V3 Service Plans", t, func() {
		mock := MockRoute{
			Method:   "GET",
			Endpoint: "/v3/service_plans",
			Output:   []string{listV3ServicePlansPayloadPage1, listV3ServicePlansPayloadPage2},
			Status:   200,
		}

		setup(mock, t)
		defer teardown()

		c := &Config{ApiAddress: server.URL, Token: "foobar"}
		client, err := NewClient(c)
		So(err, ShouldBeNil)

		servicePlans, serviceOfferings, err := client.ListV3ServicePlansForSpace("test-space", "broker-name")
		So(err, ShouldBeNil)

		// Service Plans assertions
		So(len(servicePlans), ShouldEqual, 2)
		So(len(serviceOfferings), ShouldEqual, 1)
		So(servicePlans[0].GUID, ShouldEqual, "6bf4a2dd-4d0e-4237-bec0-d06242c8802c")
		So(servicePlans[0].Name, ShouldEqual, "service-plan1")

		crDate, _ := time.Parse(time.RFC3339, "2019-06-17T13:48:43Z")
		So(servicePlans[0].CreatedAt, ShouldEqual, crDate)

		updDate, _ := time.Parse(time.RFC3339, "2022-10-26T09:37:19Z")
		So(servicePlans[0].UpdatedAt, ShouldEqual, updDate)

		So(servicePlans[0].VisibilityType, ShouldEqual, "public")
		So(servicePlans[0].Available, ShouldEqual, true)
		So(servicePlans[0].Free, ShouldEqual, true)
		So(servicePlans[0].Description, ShouldEqual, "ServicePlan 1 Description")

		So(servicePlans[0].BrokerCatalog.ID, ShouldEqual, "26c24a23-e6a2-4843-8bc6-f454edb1bd2e")
		So(servicePlans[0].BrokerCatalog.Features.PlanUpdateable, ShouldEqual, true)
		So(servicePlans[0].BrokerCatalog.Features.Bindable, ShouldEqual, true)

		So(servicePlans[0].Relationships["service_offering"].Data.GUID, ShouldEqual, "bcb04f74-4300-4885-a0be-38e3a53f8e37")

		So(servicePlans[1].BrokerCatalog.Features.PlanUpdateable, ShouldEqual, false)
		So(servicePlans[1].BrokerCatalog.Features.Bindable, ShouldEqual, true)

		So(servicePlans[1].GUID, ShouldEqual, "7bf4a2dd-4d0e-4237-bec0-d06242c8802c")
		So(servicePlans[1].Name, ShouldEqual, "service-plan2")

		// Service Offering assertions
		So(serviceOfferings[0].GUID, ShouldEqual, "bcb04f74-4300-4885-a0be-38e3a53f8e37")
		So(serviceOfferings[0].Name, ShouldEqual, "service_offering_1")

		oCrDate, _ := time.Parse(time.RFC3339, "2019-01-17T13:48:43Z")
		So(serviceOfferings[0].CreatedAt, ShouldEqual, oCrDate)
		oUpDate, _ := time.Parse(time.RFC3339, "2022-11-26T09:37:19Z")
		So(serviceOfferings[0].UpdatedAt, ShouldEqual, oUpDate)

		So(serviceOfferings[0].Description, ShouldEqual, "Service Offering 1 Description")
		So(serviceOfferings[0].Available, ShouldEqual, true)
		So(serviceOfferings[0].Tags[0], ShouldEqual, "tag1")
		So(serviceOfferings[0].Shareable, ShouldEqual, true)
		So(serviceOfferings[0].BrokerCatalog.ID, ShouldEqual, "7537ffbb-faf3-4efa-9806-af2fc14ca7cb")
		So(serviceOfferings[0].BrokerCatalog.Metadata["displayName"], ShouldEqual, "Display Name 1")
		So(serviceOfferings[0].BrokerCatalog.Features.PlanUpdateable, ShouldEqual, true)
		So(serviceOfferings[0].BrokerCatalog.Features.Bindable, ShouldEqual, true)
		So(serviceOfferings[0].BrokerCatalog.Features.InstancesRetrievable, ShouldEqual, true)
		So(serviceOfferings[0].BrokerCatalog.Features.BindingsRetrievable, ShouldEqual, true)
		So(serviceOfferings[0].BrokerCatalog.Features.AllowContextUpdates, ShouldEqual, true)

		So(serviceOfferings[0].Relationships["service_broker"].Data.GUID, ShouldEqual, "5314174b-6bbd-4fba-9f5b-a202d19deb21")

	})
}
