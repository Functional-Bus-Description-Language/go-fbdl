package ast

import (
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func TestBuildNestedInstanceDocs(t *testing.T) {
	for _, parent := range []string{"main bus", "type Container block"} {
		t.Run(parent, func(t *testing.T) {
			src := []byte("# Parent documentation\n" + parent + `
  # First child
  first config
    width = 1
  sibling status
  # Nested group
  nested group
    # Leaf documentation
    leaf status
  after status
  # Detached comment

  detached status
  # Discarded type documentation
  type Inner group
    # Type member documentation
    member status
  trailing status`)
			toks, err := token.Parse(src)
			if err != nil {
				t.Fatal(err)
			}
			file, err := Build(toks)
			if err != nil {
				t.Fatal(err)
			}
			var body Body
			if len(file.Insts) == 1 {
				body = file.Insts[0].Body
				if got := file.Insts[0].Doc.Text(src); got != "Parent documentation" {
					t.Errorf("parent: got doc %q, want %q", got, "Parent documentation")
				}
			} else if len(file.Types) == 1 {
				body = file.Types[0].Body
				if !file.Types[0].Doc.isEmpty() {
					t.Error("type documentation should be discarded")
				}
			} else {
				t.Fatal("expected one parent instance or type")
			}

			wantDocs := []string{"First child", "", "Nested group", "", "", ""}
			if len(body.Insts) != len(wantDocs) {
				t.Fatalf("got %d instances, want %d", len(body.Insts), len(wantDocs))
			}
			for i, inst := range body.Insts {
				if got := inst.Doc.Text(src); got != wantDocs[i] {
					t.Errorf("instance %s: got doc %q, want %q", inst.Name.Text(src), got, wantDocs[i])
				}
			}
			leaves := body.Insts[2].Body.Insts
			if len(leaves) != 1 {
				t.Fatalf("got %d leaves, want 1", len(leaves))
			}
			if got := leaves[0].Doc.Text(src); got != "Leaf documentation" {
				t.Errorf("leaf: got doc %q, want %q", got, "Leaf documentation")
			}
			if len(body.Types) != 1 {
				t.Fatalf("got %d nested types, want 1", len(body.Types))
			}
			if !body.Types[0].Doc.isEmpty() {
				t.Error("nested type documentation should be discarded")
			}
			members := body.Types[0].Body.Insts
			if len(members) != 1 {
				t.Fatalf("got %d type members, want 1", len(members))
			}
			if got := members[0].Doc.Text(src); got != "Type member documentation" {
				t.Errorf("type member: got doc %q, want %q", got, "Type member documentation")
			}
		})
	}
}
