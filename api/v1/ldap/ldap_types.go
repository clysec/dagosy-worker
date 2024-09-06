package ldap

type LdapConnection struct {
	Url                string `json:"url"`
	InsecureSkipVerify bool   `json:"insecureSkipVerify"`

	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`

	BaseDN string `json:"baseDN"`
}

type LdapQuery struct {
	Connection LdapConnection `json:"connection"`
	Query      string         `json:"query"`
	Attributes []string       `json:"attributes"`
}
