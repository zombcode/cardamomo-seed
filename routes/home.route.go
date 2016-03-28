package routes

import (
  "github.com/zombcode/cardamomo"
)

func HomeRoute(router *cardamomo.Router) {
  router.Get("/index", func(req cardamomo.Request, res cardamomo.Response) {
    res.Send("Hello route home/index!");
  })
}
