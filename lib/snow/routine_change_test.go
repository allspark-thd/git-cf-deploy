package snow_test

import (
	"encoding/json"
	"time"

	. "github.com/allspark-thd/git-cf-deploy/lib/snow"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RoutineChange", func() {
	Describe("Routine Changes", func() {
		var (
			chg Change
			nrc RoutineChange
		)
		BeforeEach(func() {
			layout := "yyyy-mm-dd hh:MM:ss"
			startDate, _ := time.Parse(layout, "2020-11-07 05:00:00")
			endDate, _ := time.Parse(layout, "2020-11-08 05:00:00")
			chg = Change{
				Template: "STAC12324",
				Type:     "routine",
				CIs: []CI{
					"Yard Management System (YMS)",
					"xyzzydmg01_pilot_spec_serv_db",
				},
				Requestor: "LL65YX",
				Assignee:  "rxr3101",
				Validator: "Marcos Mendez",
				StartDate: startDate,
				EndDate:   endDate,
				DeploymentMethod: DeploymentMethod{
					Type: "Manual",
					ID:   "INT00C4E",
				},
				ProjectID: "IT-04421",
				Comments:  "testing comments",
				WorkNotes: "testing work notes",
			}
			nrc = NewRoutineChange(chg)
		})
		Describe("#Create", func() {
			It("returns a change`", func() {
				// client := NewSNOWClient()
				rc := NewRoutineChange(Change{})
				Ω(rc).ShouldNot(BeNil())
			})
			It("persists template information", func() {
				Ω(nrc).ShouldNot(BeNil())
				Ω(nrc.GetChange()).Should(Equal(chg))
			})
		})
		Describe("as JSON", func() {
			var chgmap map[string]interface{}
			JustBeforeEach(func() {
				byts, _ := json.Marshal(chg)
				json.Unmarshal(byts, &chgmap)
			})
			It("matches what Marcos expects", func() {
				Ω(chgmap).Should(HaveKeyWithValue("type", chg.Type))
				Ω(chgmap).Should(HaveKeyWithValue("", chg.Type))
			})
		})
	})
})

/*
NewRoutineChange( "templatexyz", Change{

})
echo '{
	"type":"Routine",
	"template":"STCA0001018",
	"affected_cis":"Yard Management System (YMS),xyzzydmg01_pilot_spec_serv_db",
	"requested_by":"LL65YX",
	"assigned_to": "rxr3101",
	"u_validated_by": "Marcos Mendez",
	"start_date": "2020-11-07 05:00:00",
	"end_date": "2020-11-08 05:00:00",
	"u_automation": "Manual",
	"u_automation_pkg_no": "INT00C4E",
	"u_fms_number": "IT-04421",
	"comments":"testing comments",
	"work_notes":"testing work notes"
	}' | curl -X -s --max-time 60 -XPOST -H"Content-Type: application/json" -d @- --url "http://localhost:8080/records?model=changerequest"
*/
