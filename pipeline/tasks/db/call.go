package db

import (
	"fmt"

	"cuelang.org/go/cue"

  "github.com/hofstadter-io/cuetils/pipeline/context"
)

func init() {
  context.Register("db.Call", NewCall)
}

type Call struct {}

func NewCall(val cue.Value) (context.Runner, error) {
  return &Call{}, nil
}

func (T *Call) Run(ctx *context.Context) (interface{}, error) {

	v := ctx.Value

  out, err := handleQuery(v)
  if err != nil {
    return nil, err
  }

	return v.FillPath(cue.ParsePath("results"), out), nil
}

func handleQuery(val cue.Value) (interface{}, error) {

  callType := ""
	query := val.LookupPath(cue.ParsePath("query"))
	if query.Exists() && query.Err() == nil {
    callType = "query"
	}
  if callType == "" {
    query = val.LookupPath(cue.ParsePath("exec"))
    if query.Exists() && query.Err() == nil {
      callType = "exec"
    }
  }
  if callType == "" {
    query = val.LookupPath(cue.ParsePath("stmts"))
    if query.Exists() && query.Err() == nil {
      callType = "stmts"
    }
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
      switch callType {
        case "query":
          qs, err := query.String()
          if err != nil {
            return nil, fmt.Errorf("in field 'query' at %v", err)
          }

          rows, err := handleSQLiteQuery(dbname, qs, iargs)
          if err != nil {
            return nil, fmt.Errorf("error during query %v", err)
          }

          jstr, err := scanRowToJson(rows)
          if err != nil {
            return nil, fmt.Errorf("error during scan %v", err)
          }
          return val.Context().CompileBytes(jstr), nil

        case "exec":
          qs, err := query.String()
          if err != nil {
            return nil, fmt.Errorf("in field 'exec' at %v", err)
          }

          out, err := handleSQLiteExec(dbname, qs, iargs)
          if err != nil {
            return nil, fmt.Errorf("error during exec %v", err)
          }
          return out, nil

        case "stmts":
          out, err := handleSQLiteStmts(dbname, query, iargs)
          if err != nil {
            return nil, fmt.Errorf("error during query %v", err)
          }
          return out, nil
      }

		}
	}

	return "", fmt.Errorf("no supported conn types found in db.Query %q", val.Path())
}
