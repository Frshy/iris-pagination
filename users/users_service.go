package users

import(
  "github.com/kataras/iris/v12"
  "iris-pagination/models"
  "iris-pagination/globals"

  "math"
)

type ServiceT struct {}
var Service ServiceT

func (c *ServiceT) Dashboard(ctx iris.Context) {
  page := ctx.URLParamIntDefault("page", 0)
  perPage := 2
  sortBy := ctx.URLParamDefault("sortBy", "ID")
  searchFor := ctx.URLParamDefault("searchFor", "")

  var usersCount int64
  var users []models.User

  globals.Db.Model(&models.User{}).Where("Username LIKE '%" + searchFor + "%'").Count(&usersCount)

  var pagesCount int64 = int64(math.Ceil(float64(usersCount) / float64(perPage)))

  if int64(page) > pagesCount - 1  {
    page = int(pagesCount - 1)
  }

  if page < 0 {
    page = 0
  }

  //may be sql injection vulnerable, also pagination view also may be vulnerable for xss
  globals.Db.Where("Username LIKE '%" + searchFor + "%'").Limit(perPage).Offset(page * perPage).Order(sortBy).Find(&users)
  canPreviousPage := page > 0
  canNextPage := int64(page) < pagesCount - 1

  var pageNumDisplay int64
  if pagesCount > 0 {
    pageNumDisplay = int64(page) + 1
  } else {
    pageNumDisplay = 0
  }

  ctx.ViewData("page", page)
  ctx.ViewData("pagesCount", pagesCount)
  ctx.ViewData("perPage", perPage)
  ctx.ViewData("users", users)
  ctx.ViewData("previousPage", page-1)
  ctx.ViewData("nextPage", page+1)
  ctx.ViewData("canPreviousPage", canPreviousPage)
  ctx.ViewData("canNextPage", canNextPage)
  ctx.ViewData("pageNumDisplay", pageNumDisplay)
  ctx.ViewData("sortBy", sortBy)
  ctx.ViewData("searchFor", searchFor)

  if err := ctx.View("dashboard.html"); err != nil {
    ctx.HTML("<h3>%s</h3>", err.Error())
    return
  }
}
