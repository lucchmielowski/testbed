package runtime

const (
	serviceID = "testbed.services.runtime.v1"
)

type service struct {
	containerAddr string
	namespace     string
}
