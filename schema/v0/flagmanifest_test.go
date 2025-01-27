package flagmanifest

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

var compiledFlagManifestSchema *jsonschema.Schema

func init() {
	sch, err := jsonschema.CompileString(SchemaPath, Schema)
	if err != nil {
		log.Fatal(fmt.Errorf("error compiling JSON schema: %v", err))
	}
	compiledFlagManifestSchema = sch
}

func TestPositiveFlagManifest(t *testing.T) {
	if err := walkPath(true, "./tests/positive"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestNegativeFlagManifest(t *testing.T) {
	if err := walkPath(false, "./tests/negative"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}


func walkPath(shouldPass bool, root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ps := strings.Split(path, ".")
		if ps[len(ps)-1] != "json" {
			return nil
		}

		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var v interface{}
		if err := json.Unmarshal([]byte(file), &v); err != nil {
			log.Fatal(err)
		}

		err = compiledFlagManifestSchema.Validate(v)

		if (err != nil && shouldPass == true) {
			return fmt.Errorf("file %s should not have failed validation, but did: %s", path, err)
		}

		if (err == nil && shouldPass == false) {
			return fmt.Errorf("file %s should have failed validation, but did not", path)
		}

		return nil
	})
}
