package api

import (
  "encoding/json"
	"fmt"
	// "io"
  "net/http"
  "strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"
  "github.com/labstack/echo/v4"

	// "github.com/hofstadter-io/cuetils/utils"
)

type Serve struct {}

func NewServe(val cue.Value) (flow.Runner, error) {
  return &Serve{}, nil
}

func (T *Serve) Run(t *flow.Task, err error) error {
  fmt.Println("api.Serve - starting")
	if err != nil {
		fmt.Println("Dep error", err)
	}

	val := t.Value()

  // get the port
  p := val.LookupPath(cue.ParsePath("port"))
  if p.Err() != nil {
    return p.Err()
  }
  port, err := p.String()
  if err != nil {
    return err
  }

  // create server
  e := echo.New()
  e.HideBanner = true

  //
  // Setup routes
  //
  routes := val.LookupPath(cue.ParsePath("routes"))
  iter, err := routes.Fields()
  if err != nil {
    return err
  }

  for iter.Next() {
    label := iter.Selector().String()
    route := iter.Value()

    fmt.Println("route:", label)

    err := routeFromValue(label, route, e)
    if err != nil {
      return err
    }
  }


  // run the server
	e.GET("/alive", func(c echo.Context) error {
		return c.String(http.StatusNoContent, "Hello, World!")
	})

  data, err := json.MarshalIndent(e.Routes(), "", "  ")
  if err != nil {
    return err
  }

  fmt.Println(string(data))

	e.Logger.Fatal(e.Start(":" + port))


  fmt.Println("SERVER EXITED")

  // - pull apart server value

  /*
  - loop over routes...
    - should be a pipeline, we need to load this
    - construct the routes, with pipeline & echo framework
  - go run the server
  - wait for exit signal


  */


	//req := val.LookupPath(cue.ParsePath("req"))

  // =======================


	//R, err := buildRequest(req)
	//if err != nil {
		//return err
	//}

	//actual, err := makeRequest(R)
	//if err != nil {
		//return err
	//}

	//body, err := io.ReadAll(actual.Body)
	//if err != nil {
		//return err
	//}

	//// name better based on path in CUE code
	//resp := val.Context().CompileBytes(body, cue.Filename("resp"))

	//// Use fill to "return" a result to the workflow engine
	//res := val.FillPath(cue.ParsePath("resp"), resp)

	//attr := val.Attribute("print")
	//err = utils.PrintAttr(attr, res)

	//t.Fill(res)

  // fmt.Println("end: API call")
	return err
}

func routeFromValue(path string, route cue.Value, e *echo.Echo) (error) {
  path = strings.Replace(path, "\"", "", -1)
  fmt.Println(path + ":", route)

  // is this a pipeline handler?
  attrs := route.Attributes(cue.ValueAttr)
  isPipe := false
  for _, a := range attrs {
    if a.Name() == "pipeline" {
      isPipe = true
    } 
  }

  if isPipe {
    // TODO
    return nil
  }

  local := route

  fmt.Println("setting up route:", path)

  // e.Match([]string{"GET"}, path, func (c echo.Context) error {
  e.GET(path, func (c echo.Context) error {
    // pull apart c.request

    tmp := local.FillPath(cue.ParsePath("req"), "foobar")
    if tmp.Err() != nil {
      return tmp.Err()
    }

    resp := tmp.LookupPath(cue.ParsePath("resp"))
    if resp.Err() != nil {
      return resp.Err()
    }

    var ret interface{}
    err := resp.Decode(&ret)
    if err != nil {
      return err
    }

    c.JSON(http.StatusOK, ret)

    return nil
  })

  return nil
}
