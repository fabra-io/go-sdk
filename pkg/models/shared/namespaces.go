// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Namespaces struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

func (o *Namespaces) GetNamespaces() []string {
	if o == nil {
		return nil
	}
	return o.Namespaces
}
