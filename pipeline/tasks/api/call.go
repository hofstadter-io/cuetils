package api

import (
	"fmt"
	"io"
	"strings"
	"time"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"

	"github.com/hofstadter-io/cuetils/utils"
	"github.com/parnurzeal/gorequest"
)

type Call struct {
	Req cue.Value
	Ret cue.Value
}

func NewCall(val cue.Value) (flow.Runner, error) {
  return &Call{}, nil
}

func (T *Call) Run(t *flow.Task, err error) error {
  // fmt.Println("beg: API call")
	if err != nil {
		fmt.Println("Dep error", err)
	}

	val := t.Value()

	req := val.LookupPath(cue.ParsePath("req"))

	R, err := buildRequest(req)
	if err != nil {
		return err
	}

	actual, err := makeRequest(R)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(actual.Body)
	if err != nil {
		return err
	}

	// name better based on path in CUE code
	resp := val.Context().CompileBytes(body, cue.Filename("resp"))

	// Use fill to "return" a result to the workflow engine
	res := val.FillPath(cue.ParsePath("resp"), resp)

	attr := val.Attribute("print")
	err = utils.PrintAttr(attr, res)

	t.Fill(res)

  // fmt.Println("end: API call")
	return err
}

/********* old *********/

const HTTP2_GOAWAY_CHECK = "http2: server sent GOAWAY and closed the connection"

func buildRequest(val cue.Value) (R *gorequest.SuperAgent, err error) {
	req := val.Eval()
	R = gorequest.New()

	method := req.LookupPath(cue.ParsePath("method"))
	R.Method, err = method.String()
	if err != nil {
		return
	}

	host := req.LookupPath(cue.ParsePath("host"))
	hostStr, err := host.String()
	if err != nil {
		return
	}

	path := req.LookupPath(cue.ParsePath("path"))
	pathStr, err := path.String()
	if err != nil {
		return
	}
	R.Url = hostStr + pathStr

	headers := req.LookupPath(cue.ParsePath("headers"))
	if headers.Exists() {
		H, err := headers.Struct()
		if err != nil {
			return R, err
		}
		hIter := H.Fields()
		for hIter.Next() {
			label := hIter.Label()
			value, err := hIter.Value().String()
			if err != nil {
				return R, err
			}
			R.Header.Add(label, value)
		}
	}

	query := req.LookupPath(cue.ParsePath("query"))
	if query.Exists() {
		Q, err := query.Struct()
		if err != nil {
			return R, err
		}
		qIter := Q.Fields()
		for qIter.Next() {
			label := qIter.Label()
			value, err := qIter.Value().String()
			if err != nil {
				return R, err
			}
			R.QueryData.Add(label, value)
		}
	}

	data := req.LookupPath(cue.ParsePath("data"))
	if data.Exists() {
		err := data.Decode(&R.Data)
		if err != nil {
			return R, err
		}
	}

	timeout := req.LookupPath(cue.ParsePath("timeout"))
	if timeout.Exists() {
		to, err := timeout.String()
		if err != nil {
			return R, err
		}
		d, err := time.ParseDuration(to)
		if err != nil {
			return R, err
		}
		R.Timeout(d)
	}

	retry := req.LookupPath(cue.ParsePath("retry"))
	if retry.Exists() {
		C := 3
		count := retry.LookupPath(cue.ParsePath("count"))
		if count.Exists() {
			c, err := count.Int64()
			if err != nil {
				return R, err
			}
			C = int(c)
		}

		D := time.Second * 6
		timer := retry.LookupPath(cue.ParsePath("timer"))
		if timer.Exists() {
			t, err := timer.String()
			if err != nil {
				return R, err
			}
			d, err := time.ParseDuration(t)
			if err != nil {
				return R, err
			}
			D = d
		}

		CS := []int{}
		codes := retry.LookupPath(cue.ParsePath("codes"))
		if codes.Exists() {
			L, err := codes.List()
			if err != nil {
				return R, err
			}
			for L.Next() {
				v, err := L.Value().Int64()
				if err != nil {
					return R, err
				}
				CS = append(CS, int(v))
			}
		}

		R.Retry(C, D, CS...)
	}

	return
}

func makeRequest(R *gorequest.SuperAgent) (gorequest.Response, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in HTTP: %v %v\n", R, r)
		}
	}()

	resp, body, errs := R.End()

	if len(errs) != 0 && resp == nil {
		return resp, fmt.Errorf("%v", errs)
	}

	if len(errs) != 0 && !strings.Contains(errs[0].Error(), HTTP2_GOAWAY_CHECK) {
		return resp, fmt.Errorf("Internal Weirdr Error:\b%v\n%s\n", errs, body)
	}
	if len(errs) != 0 {
		return resp, fmt.Errorf("Internal Error:\n%v\n%s\n", errs, body)
	}

	return resp, nil
}
