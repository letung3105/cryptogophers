package set02

import (
	"bytes"
	"log"
	"reflect"
	"testing"
)

func TestCreateParse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		key   []byte
		email []byte
		doc   map[string][]byte
	}{
		{
			[]byte("YELLOW SUBMARINE"),
			[]byte("foo@bar.com"),
			map[string][]byte{
				"email": []byte("foo@bar.com"),
				"uid":   []byte("10"),
				"role":  []byte("user"),
			},
		},
		{
			[]byte("YELLOW SUBMARINE"),
			[]byte("foo@bar.com&role=admin"),
			map[string][]byte{
				"email": []byte("foo@bar.comroleadmin"),
				"uid":   []byte("10"),
				"role":  []byte("user"),
			},
		},
	}

	for _, test := range tests {
		prof, err := ProfileForEncrypt(test.email, test.key)
		if err != nil {
			log.Fatalf("unexpected error: %+v", err)
		}

		doc, err := KVParseDecrypt(prof, test.key)
		if err != nil {
			log.Fatalf("unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(doc, test.doc) {
			log.Fatalf("unexpected output:\nhave %v\nwant %v", doc, test.doc)
		}
	}
}

func TestModProfileRole(t *testing.T) {
	t.Parallel()
	test := struct {
		key  []byte
		role []byte
	}{
		[]byte("YELLOW SUBMARINE"),
		[]byte("admin"),
	}

	out, err := ModProfileRole(test.key)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	doc, err := KVParseDecrypt(out, test.key)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if !bytes.Equal(doc["role"], test.role) {
		t.Errorf("unexpected output:\nhave %s\nwant %s", doc["role"], test.role)
	}
}
