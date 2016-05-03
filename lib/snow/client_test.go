package snow_test

import (
	"net/url"

	. "github.com/allspark-thd/git-cf-deploy/lib/snow"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Describe("NewSNOWClient", func() {
		It("requires a MarcosURL", func() {
			cfg := SNOWConfig{}
			_, err := NewSNOWClient(cfg)
			Ω(err).
				Should(HaveOccurred())

			cfg.URL, _ = url.Parse("https://snowfield.com")
			_, err = NewSNOWClient(cfg)
			Ω(err).
				ShouldNot(HaveOccurred())
		})
	})
	Describe("Routine Changes", func() {
		Describe("#Create", func() {
			It("requires `template`", func() {
				// client := NewSNOWClient()
				// rc := RoutineChange{}

			})
		})
	})
})

/*
echo '{"type":"Routine",
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
