package controllers

import (
        "github.com/revel/revel"
        "github.com/dr013/carousel/app/models"
)

type App struct {
        *revel.Controller
}

func (c App) Index() revel.Result {
        var (
                schemas []models.Schema
                err     error
        )
        schemas, err = models.GetSchemas()
        if err != nil {
                errResp := buildErrResponse(err, "500")
                c.Response.Status = 500
                return c.RenderJson(errResp)
        }
        c.Response.Status = 200
        return c.Render(schemas)
}
