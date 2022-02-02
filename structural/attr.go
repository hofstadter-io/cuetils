package structural

import (
	"fmt"

	"cuelang.org/go/cue"

  // "github.com/hofstadter-io/cuetils/utils"
)

func InjectAttrsValue(val cue.Value, attr string, kv map[string]string) (cue.Value, error) {
  return val, nil
//  fmt.Println(attr, kv)
	r, _ := injectAttrsValue(val, attr, kv)
 // s, _ := utils.PrintCue(r)
  //fmt.Println("s:", s)
	return r, nil
}

func injectAttrsValue(val cue.Value, attr string, kv map[string]string) (cue.Value, error) {
	switch val.IncompleteKind() {
	case cue.StructKind:
		return injectAttrsStruct(val, attr, kv)

	case cue.ListKind:
		return injectAttrsList(val, attr, kv)

	default:

    //  Attrs injected here 
		return injectAttrs(val, attr, kv)
	}
}

func injectAttrsStruct(val cue.Value, attr string, kv map[string]string) (cue.Value, error) {

	result := val
	iter, _ := val.Fields(defaultWalkOptions...)

	for iter.Next() {
		s := iter.Selector()
		p := cue.MakePath(s)

    r, err := injectAttrsValue(iter.Value(), attr, kv)
    if err != nil {
      return result, err
    }
    result = result.FillPath(p, r)
	}

	return result, nil 
}

func injectAttrsList(val cue.Value, attr string, kv map[string]string) (cue.Value, error) {
	ctx := val.Context()

	vi, _ := val.List()

	result := []cue.Value{}
	for vi.Next() {
		r, err := injectAttrsValue(vi.Value(), attr, kv)
    if err != nil {
      return val, err
    }
    result = append(result, r)
	}
	return ctx.NewList(result...), nil
}

func injectAttrs(val cue.Value, attr string, kv map[string]string) (cue.Value, error) {

  attrs := val.Attributes(cue.ValueAttr)
  for _, a := range attrs {
    // found our attribute, like @tag()
    if a.Name() == attr {
      if a.NumArgs() != 1 {
        return val, fmt.Errorf("@%s() attribute requires a single argument", attr)
      }
      
      key, err := a.String(0)
      if err != nil {
        return val, err
      }

      data, ok := kv[key]
      if !ok {
        continue
      }


      if val.IncompleteKind() == cue.StringKind {
        data = "\"" + data + "\""
      }
      // fmt.Println("injecting:", val.Path(), a, data)
      x := val.Context().CompileString(fmt.Sprint(data))
      if x.Err() != nil {
        return val, x.Err()
      }

      // fmt.Println("val.prefill:", val, x)
      val = val.Unify(x)
      if val.Err() != nil {
        return val, val.Err()
      }
      // fmt.Println("val.postfill:", val)
      break
    }
  }

  return val, nil
}
