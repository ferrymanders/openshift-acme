package extended

import (
	"testing"

	_ "github.com/ferrymanders/openshift-acme/test/e2e/openshift/routes"
	exutil "github.com/ferrymanders/openshift-acme/test/e2e/openshift/util"
)

func init() {
	exutil.InitTest()
}

func TestExtended(t *testing.T) {
	exutil.ExecuteTest(t, "Extended")
}
