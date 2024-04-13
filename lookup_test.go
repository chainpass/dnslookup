package dnslookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupCNAME(t *testing.T) {
	record, err := Lookup("CNAME", "proxy.cymbal.co")
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, record, 1)
	assert.Equal(t, "proxy.cymbal.co", record[0].Name)
	assert.Equal(t, "cname.vercel-dns.com.", record[0].Value)
}

func TestLookupMX(t *testing.T) {
	records, err := Lookup("MX", "mail.cymbal.co")
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, records, 1)
	assert.Equal(t, "mail.cymbal.co", records[0].Name)
	assert.Equal(t, 10, records[0].Priority)
	assert.Equal(t, "feedback-smtp.us-west-2.amazonses.com.", records[0].Value)
}

func TestLookupTXT(t *testing.T) {
	records, err := Lookup("TXT", "mail.cymbal.co")
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, records, 1)
	assert.Equal(t, "mail.cymbal.co", records[0].Name)
	assert.Equal(t, "v=spf1 include:amazonses.com ~all", records[0].Value)
}
