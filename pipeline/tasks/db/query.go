package db

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/utils"
)

type Query struct{}

func NewQuery(val cue.Value) (flow.Runner, error) {
	return &Query{}, nil
}

func (T *Query) Run(t *flow.Task, err error) error {
	if err != nil {
		fmt.Println("Dep error", err)
	}

	v := t.Value()

  out, err := handleQuery(v)
  if err != nil {
    return err
  }

	v = v.FillPath(cue.ParsePath("results"), out)

	attr := v.Attribute("print")
	err = utils.PrintAttr(attr, v)

	t.Fill(v)

	return err
}

func handleQuery(val cue.Value) (interface{}, error) {

  isQuery := true 
	query := val.LookupPath(cue.ParsePath("query"))
	if !query.Exists() {
    isQuery = false
    query = val.LookupPath(cue.ParsePath("exec"))
    if !query.Exists() {
      return nil, fmt.Errorf("field 'query' or 'exec' is required on db.Query at %q", val.Path())
    }
	}

	qs, err := query.String()
	if err != nil {
		return nil, fmt.Errorf("in field 'query' at %v", err)
	}

	conn := val.LookupPath(cue.ParsePath("conn"))
	if !conn.Exists() {
		return nil, fmt.Errorf("field 'conn' is required on db.Query at %q", val.Path())
	}

	iter, err := conn.Fields()
	if err != nil {
		return nil, fmt.Errorf("in field 'conn' at %v", err)
	}

	var args []string
	av := val.LookupPath(cue.ParsePath("args"))
	if av.Exists() {
		err = av.Decode(&args)
		if err != nil {
			return nil, fmt.Errorf("while decoding 'args' at %v", err)
		}
	}

	iargs := []interface{}{}
	for _, a := range args {
		iargs = append(iargs, a)
	}

	for iter.Next() {
		sel := iter.Selector().String()
		switch sel {
		case "sqlite":
			dbname, err := iter.Value().String()
			if err != nil {
				return nil, err
			}
      if isQuery {
        rows, err := handleSQLiteQuery(dbname, qs, iargs)
        if err != nil {
          return nil, fmt.Errorf("error during query %v", err)
        }

        jstr, err := scanRowToJson(rows)
        if err != nil {
          return nil, fmt.Errorf("error during scan %v", err)
        }

        return val.Context().CompileBytes(jstr), nil
      } else {
        out, err := handleSQLiteExec(dbname, qs, iargs)
        if err != nil {
          return nil, fmt.Errorf("error during query %v", err)
        }
        return out, nil
      }

		}
	}

	return "", fmt.Errorf("no supported conn types found in db.Query %q", val.Path())
}
